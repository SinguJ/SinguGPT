package main

import (
	"fmt"
	"log"
	"time"

	"SinguGPT/imap"
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
	// 对接 IMAP
	imapClient := imap.NewClient(&imap.EmailConfig{
		Host:     store.Config.Email.IMAP.Host,
		Port:     store.Config.Email.IMAP.Port,
		Username: store.Config.Email.IMAP.UserName,
		Password: store.Config.Email.IMAP.Password,
		Debug:    false,
	})
	mails := make(chan *imap.Mail, 20)
	errorChannel := make(chan error, 1)
	err = imapClient.Read(mails, errorChannel, 5*time.Second)
	if err != nil {
		panic(err)
	}
	go func() {
		err := <-errorChannel
		if err != nil {
			panic(err)
		}
	}()
	for {
		mail := <-mails
		log.Printf("Email => %v", mail)
		for index, content := range mail.Contents {
			log.Printf("EmailContents %d => %v", index, content)
		}
	}
}
