package main

import (
	"fmt"

	"SinguGPT/models"
	"SinguGPT/store"
)

func main() {
	println(fmt.Sprintf("应用程序名称：%s\n", store.Config.App.Name))
	// 加载并监听用户配置文件
	store.LoadAndWatchUsers()
	// 遍历用户集
	fmt.Println("用户：")
	store.ForeachUsers(func(user *models.User) {
		fmt.Printf("    %s\n", user.Name)
	})
}
