package store

import (
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

	return nil
}

func LoadAndWatchUsers() {
	// 创建文件系统监视器
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer func(watcher *fsnotify.Watcher) {
		err := watcher.Close()
		if err != nil {
			panic(err)
		}
	}(watcher)

	// 监听文件改动事件
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					// 文件发生写入操作，调用 Load 函数
					err := loadUsers(event.Name)
					if err != nil {
						log.Println(err)
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	// 添加要监视的文件
	absPath, err := filepath.Abs(Config.App.UserDataFile)
	if err != nil {
		log.Fatal(err)
	}
	err = watcher.Add(absPath)
	if err != nil {
		log.Fatal(err)
	}
	// 手动触发一次
	err = loadUsers(absPath)
	if err != nil {
		log.Fatal(err)
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
