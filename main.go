package main

import (
    "fmt"
    "log"
    "time"

    "github.com/google/uuid"

    "SinguGPT/action"
    "SinguGPT/gpt"
    "SinguGPT/imap"
    "SinguGPT/models"
    "SinguGPT/smtp"
    "SinguGPT/store"
)

var smtpClient *smtp.Client

func init() {
    smtpClient = smtp.NewSmtpClient(
        store.Config.Email.SMTP.Host,
        store.Config.Email.SMTP.Port,
        store.Config.Email.SMTP.UserName,
        store.Config.Email.SMTP.Password,
        store.Config.App.Name,
        fmt.Sprintf("[%s] 响应", store.Config.App.Name),
    )
    gpt.Login(store.Config.OpenAI.ApiKey)
}

func findSession(user *models.User, _ string) string {
    return user.ID
}

func sendEmail(user *models.User, email string, content string) {
    err := smtpClient.Push(user, email, &models.Message{
        Msg: content,
    })
    if err != nil {
        log.Fatal(err)
    }
}

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
        for _, email := range _user.Emails {
            fmt.Printf("        %s\n", email)
        }
    })
    // 接收邮件
    imapClient := imap.NewClient(&imap.EmailConfig{
        Host:     store.Config.Email.IMAP.Host,
        Port:     store.Config.Email.IMAP.Port,
        Username: store.Config.Email.IMAP.UserName,
        Password: store.Config.Email.IMAP.Password,
        Debug:    false,
    })
    mails := make(chan *imap.Mail, 20)
    errorChannel := make(chan error, 1)
    err := imapClient.Listen(mails, errorChannel, 5*time.Second)
    if err != nil {
        panic(err)
    }
    for {
        select {
        case err := <-errorChannel:
            if err != nil {
                log.Printf("[ERROR] %v", err)
            }
        case mail := <-mails:
            go func() {
                email := mail.From[0][0]
                currUser := store.FindUser(email)
                if currUser == nil {
                    log.Printf("[WRANNING] 邮箱用户 %s<%s> 不是有效用户，跳过\n", mail.From[0][1], mail.From[0][0])
                    return
                }
                requestId := uuid.NewString()
                log.Printf("[INFO] %s 处理用户 %s<%s> 的请求\n", requestId, currUser.Name, email)
                sessionId := findSession(currUser, email)
                resp, err := action.DoAction(mail.Subject, sessionId, requestId, currUser, mail.Contents[0].Text)
                if err != nil {
                    resp = fmt.Sprintf("%v", err)
                    log.Printf("[ERROR] %s %v", requestId, err)
                }
                sendEmail(currUser, email, resp)
            }()
        }
    }
}
