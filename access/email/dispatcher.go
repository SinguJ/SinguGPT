package email

import (
    "fmt"
    "log"
    "time"

    "github.com/google/uuid"

    "SinguGPT/access"
    "SinguGPT/access/email/imap"
    "SinguGPT/access/email/smtp"
    "SinguGPT/models"
    "SinguGPT/store"
)

type Dispatcher struct {
    smtpClient *smtp.Client
    imapClient *imap.Client

    mails        chan *imap.Mail
    errorChannel chan error

    messageHandler access.MessageHandler
}

func (d *Dispatcher) OnMessageReceive(handler access.MessageHandler) error {
    d.messageHandler = handler
    return nil
}

func (d *Dispatcher) OnBatchUserChange(_ access.BatchUserChangeHandler) error {
    // operation not supported
    return nil
}

func (d *Dispatcher) Listen() error {
    d.smtpClient = smtp.NewSmtpClient(
        store.Config.Email.SMTP.Host,
        store.Config.Email.SMTP.Port,
        store.Config.Email.SMTP.UserName,
        store.Config.Email.SMTP.Password,
        store.Config.App.Name,
        fmt.Sprintf("[%s] 响应", store.Config.App.Name),
    )
    d.imapClient = imap.NewClient(&imap.EmailConfig{
        Host:     store.Config.Email.IMAP.Host,
        Port:     store.Config.Email.IMAP.Port,
        Username: store.Config.Email.IMAP.UserName,
        Password: store.Config.Email.IMAP.Password,
        Debug:    false,
    })
    err := d.imapClient.Listen(d.mails, d.errorChannel, 5*time.Second)
    if err != nil {
        return err
    }
    go func() {
        for {
            select {
            case err := <-d.errorChannel:
                if err != nil {
                    log.Printf("[ERROR] %v", err)
                }
            case mail := <-d.mails:
                go func() {
                    email := mail.From[0][0]
                    user := store.FindUser(email)
                    if user == nil {
                        log.Printf("[WARNING] 邮箱用户 %s<%s> 不是有效用户，跳过\n", mail.From[0][1], mail.From[0][0])
                        return
                    }
                    requestId := uuid.NewString()
                    log.Printf("[INFO]>>> %s 处理用户 %s<%s> 的请求\n", requestId, user.Name, email)
                    req := models.Contents{
                        models.NewTextContent(models.TagCommand, mail.Subject),
                        models.NewTextContent(models.TagBody, mail.Contents[0].Text),
                    }
                    resp, err := d.messageHandler(user.ID, requestId, user, req)
                    if err != nil {
                        resp = models.Contents{
                            models.NewTextContent(models.TagError, fmt.Sprintf("%v", err)),
                        }
                        log.Printf("[ERROR]--- %s %v", requestId, err)
                    }
                    err = d.smtpClient.Push(user, email, resp[0])
                    if err != nil {
                        log.Printf("[ERROR]--- %s %v", requestId, err)
                    }
                    log.Printf("[INFO]<<< %s 用户 %s<%s> 的请求处理完成\n", requestId, user.Name, email)
                }()
            }
        }
    }()
    return nil
}

func NewDispatcher() access.Dispatcher {
    return &Dispatcher{
        mails:        make(chan *imap.Mail, 20),
        errorChannel: make(chan error, 1),
    }
}
