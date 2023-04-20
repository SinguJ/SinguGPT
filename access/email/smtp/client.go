package smtp

import (
    "strings"

    "github.com/go-gomail/gomail"

    "SinguGPT/models"
)

type EmailConfig struct {
    // SMTP 服务器主机名
    host string
    // SMTP 服务器端口
    port int
    // SMTP 用户名
    username string
    // SMTP 密码
    password string
    // 发件人名称
    senderName string
    // 主题
    subject string
}

type Client struct {
    // Email 配置
    emailConfig *EmailConfig
}

func (p *Client) Push(user *models.User, email string, content models.Content) error {
    message := p.buildMessage(content, p.emailConfig.subject, email, user.Name)
    dialer := p.createDialer()
    return dialer.DialAndSend(message)
}

func (p *Client) buildMessage(content models.Content, subject string, receiverAddr string, receiverName string) *gomail.Message {
    message := gomail.NewMessage()
    // 设置发件人
    message.SetAddressHeader("From", p.emailConfig.username, p.emailConfig.senderName)
    // 设置收件人
    message.SetAddressHeader("To", strings.ToLower(strings.TrimSpace(receiverAddr)), receiverName)
    // 设置抄送人
    // message.SetHeader("Cc", message.FormatAddress("", ""), ...)
    // 设置密送人
    // message.SetHeader("Bcc", message.FormatAddress("", ""), ...)
    message.SetHeader("Subject", subject)
    message.AddAlternative("text/plain", content.ToString())
    // 添加附件
    // message.Attach("")
    return message
}

func (p *Client) createDialer() *gomail.Dialer {
    config := p.emailConfig
    return gomail.NewDialer(config.host, config.port, config.username, config.password)
}

// NewSmtpClient 创建 SMTP 客户端
func NewSmtpClient(host string, port int, username string, password string, senderName string, subject string) *Client {
    config := EmailConfig{
        host,
        port,
        username,
        password,
        senderName,
        subject,
    }
    client := Client{&config}
    return &client
}
