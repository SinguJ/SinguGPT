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

var roles []*models.Role

var userMapping map[string]*models.User

func loadRolesAndUsers(path string) error {
    // 加载新用户与角色数据
    userConfig := models.UserConfig{}
    err := loadYaml(path, &userConfig)
    if err != nil {
        return err
    }

    // 构建角色映射
    var _roleMapping = make(map[string]*models.Role)
    for _, role := range userConfig.Roles {
        _roleMapping[role.Name] = role
    }
    // 构建用户映射及用户数组
    var _users = make([]*models.User, len(userConfig.Users))
    var _userMapping = make(map[string]*models.User)
    index := 0
    for userId := range userConfig.Users {
        user := userConfig.Users[userId]
        user.ID = userId
        user.Role = _roleMapping[user.RoleName]
        _users[index] = user
        index++
        for _, email := range user.Emails {
            email = strings.ToLower(strings.TrimSpace(email))
            _userMapping[email] = user
        }
    }

    // 替换新用户及角色数据
    users = _users
    userMapping = _userMapping
    roles = userConfig.Roles

    var buff bytes.Buffer
    // 遍历角色集
    buff.WriteString("角色集：{{%%\n")
    ForeachRoles(func(role *models.Role) {
        buff.WriteString(fmt.Sprintf("  %s:\n", role.Name))
        for _, perm := range role.Perms {
            buff.WriteString(fmt.Sprintf("      - %s\n", perm))
        }
    })
    buff.WriteString("%%}}\n")
    // 遍历用户集
    buff.WriteString("用户集：{{%%\n")
    ForeachUsers(func(user *models.User) {
        buff.WriteString(fmt.Sprintf("  %s:\n", user.Name))
        for _, email := range user.Emails {
            buff.WriteString(fmt.Sprintf("      - %s\n", email))
        }
    })
    buff.WriteString("%%}}\n")
    log.Printf("[INFO - USERS] \n%s", buff.String())
    return nil
}

func LoadAndWatchRolesAndUsers() {
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
                    err := loadRolesAndUsers(event.Name)
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
    err = loadRolesAndUsers(absPath)
    if err != nil {
        panic(err)
    }
}

func ForeachRoles(action func(role *models.Role)) {
    for _, role := range roles {
        action(role)
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
