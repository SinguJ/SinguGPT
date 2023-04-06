package main

import (
	"fmt"

	"SinguGPT/models"
	"SinguGPT/smtp"
	"SinguGPT/store"
)

func main() {
	println(fmt.Sprintf("应用程序名称：%s\n", store.Config.App.Name))
	// 加载并监听用户配置文件
	store.LoadAndWatchUsers()
	// 遍历用户集
	var user *models.User
	fmt.Println("用户：")
	store.ForeachUsers(func(_user *models.User) {
		if user == nil {
			user = _user
		}
		fmt.Printf("    %s\n", _user.Name)
	})
	// 对接 SMTP
	smtpClient := smtp.NewSmtpClient(
		store.Config.Email.SMTP.Host,
		store.Config.Email.SMTP.Port,
		store.Config.Email.SMTP.UserName,
		store.Config.Email.SMTP.Password,
		store.Config.App.Name,
		fmt.Sprintf("[%s] 响应", store.Config.App.Name),
	)
	err := smtpClient.Push(user, user.Emails[0], &models.Message{
		Msg: "测试",
	})
	if err != nil {
		panic(err)
	}
}
