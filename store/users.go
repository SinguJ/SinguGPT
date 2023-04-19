package store

import (
    "bytes"
    "fmt"
    "log"
    "path/filepath"
    "strings"

    "github.com/fsnotify/fsnotify"

    "SinguGPT/models"
)

var users []*models.User

var userMapping map[string]*models.User

func loadUsers(path string) error {
    // 加载新用户数据
    userConfig := models.UserConfig{}
    err := loadYaml(path, &userConfig)
    if err != nil {
        return err
    }

    // 构建用户映射及用户数组
    var _users = make([]*models.User, len(userConfig.Users))
    var _userMapping = make(map[string]*models.User)
    index := 0
    for userId := range userConfig.Users {
        user := userConfig.Users[userId]
        user.ID = userId
        _users[index] = user
        index++
        for _, email := range user.Emails {
            email = strings.ToLower(strings.TrimSpace(email))
            _userMapping[email] = user
        }
    }

    // 替换新用户数据
    users = _users
    userMapping = _userMapping

    // 遍历用户集
    var user *models.User
    var buff bytes.Buffer
    ForeachUsers(func(_user *models.User) {
        if user == nil {
            user = _user
        }
        buff.WriteString(fmt.Sprintf("  %s:\n", _user.Name))
        for _, email := range _user.Emails {
            buff.WriteString(fmt.Sprintf("      - %s\n", email))
        }
    })
    log.Printf("[INFO - USERS] 用户集：{{%%%%\n%s\n%%%%}}\n", buff.String())
    return nil
}

func LoadAndWatchUsers() {
    // 创建文件系统监视器
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        panic(err)
    }
    defer func(watcher *fsnotify.Watcher) {
        err := recover()
        if err != nil {
            _err := watcher.Close()
            if _err != nil {
                panic(_err)
            }
            panic(err)
        }
    }(watcher)
    // 监听文件改动事件
    go func() {
        defer func(watcher *fsnotify.Watcher) {
            err := watcher.Close()
            if err != nil {
                panic(err)
            }
        }(watcher)
        for {
            select {
            case event, ok := <-watcher.Events:
                if !ok {
                    return
                }
                if event.Has(fsnotify.Write) {
                    log.Println("[INFO - USERS] 检测到用户配置文件变更")
                    // 文件发生写入操作，调用 Load 函数
                    err := loadUsers(event.Name)
                    if err != nil {
                        log.Printf("[ERROR - USERS] %v\n", err)
                    }
                }
            case err, ok := <-watcher.Errors:
                if !ok {
                    return
                }
                log.Printf("[ERROR - USERS] %v\n", err)
            }
        }
    }()

    // 添加要监视的文件
    absPath, err := filepath.Abs(Config.App.UserDataFile)
    if err != nil {
        panic(err)
    }
    err = watcher.Add(absPath)
    if err != nil {
        panic(err)
    }
    // 手动触发一次
    err = loadUsers(absPath)
    if err != nil {
        panic(err)
    }
}

func ForeachUsers(action func(*models.User)) {
    for _, user := range users {
        action(user)
    }
}

func FindUser(email string) *models.User {
    return userMapping[strings.ToLower(strings.TrimSpace(email))]
}
