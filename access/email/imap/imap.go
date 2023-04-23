package imap

import (
    "fmt"
    "log"
    "os"
    "time"

    "github.com/emersion/go-imap/v2"
    "github.com/emersion/go-imap/v2/imapclient"

    "SinguGPT/errors"
)

// EmailConfig IMAP 服务配置
type EmailConfig struct {
    // 主机名
    Host string
    // 端口
    Port int
    // 用户名
    Username string
    // 密码
    Password string
    // 调试模式
    Debug bool
}

// Client IMAP 客户端
type Client struct {
    // 配置
    config *EmailConfig
    // IMAP 客户端
    dialer *imapclient.Client
    // 定时器
    ticker *time.Ticker
}

// 获取 IMAP Dialer
func (c *Client) createDialer() (*imapclient.Client, error) {
    var dialer *imapclient.Client
    var err error

    port := c.config.Port
    serverAddr := fmt.Sprintf("%s:%d", c.config.Host, port)
    options := &imapclient.Options{
        DebugWriter: nil,
    }
    if c.config.Debug {
        options.DebugWriter = os.Stdout
    }
    dialer, err = imapclient.DialTLS(serverAddr, options)
    if err != nil {
        _ = c.dialer.Close()
        if c.config.Debug {
            panic(errors.Wrap(err))
        }
        return nil, err
    }
    return dialer, nil
}

// Login 登录
func (c *Client) login() error {
    if c.dialer != nil && c.dialer.State() != imap.ConnStateLogout {
        err := c.close()
        if err != nil {
            if c.config.Debug {
                panic(errors.Wrap(err))
            }
            return err
        }
    }
    dialer, err := c.createDialer()
    if err != nil {
        if c.config.Debug {
            panic(errors.Wrap(err))
        }
        return err
    }
    c.dialer = dialer
    return c.dialer.Login(c.config.Username, c.config.Password).Wait()
}

// Close 关闭
func (c *Client) close() error {
    if c.ticker != nil {
        c.ticker.Stop()
    }
    if c.dialer == nil {
        return nil
    }
    defer func(dialer *imapclient.Client) {
        err := dialer.Close()
        if err != nil {
            if c.config.Debug {
                panic(errors.Wrap(err))
            }
            panic(errors.Wrap(err))
        }
    }(c.dialer)
    err := c.dialer.Logout().Wait()
    if err != nil {
        return err
    }
    c.dialer = nil
    return nil
}

// 尝试修复错误
func (c *Client) attemptToFixErrors(err error) {
    // 检查客户端状态
    if c.dialer == nil || c.dialer.State() == imap.ConnStateLogout {
        log.Println("[WARNING - IMAP] IMAP 连接被意外关闭，程序将自动重连...")
        err := c.login()
        if err != nil {
            log.Printf("[ERROR - IMAP] IMAP 重连失败，原因是：%v\n", err)
            panic(errors.Wrap(err))
        }
        return
    }
    log.Printf("[ERROR - IMAP] IMAP 重连失败，原因是：%v\n", err)
    panic(errors.Wrap(err))
}

func (c *Client) Listen(channel chan *Mail, errorChannel chan error, duration time.Duration) error {
    go func() {
        defer func() {
            if !c.config.Debug {
                err := recover()
                if err != nil {
                    errorChannel <- err.(error)
                    return
                }
            }
        }()
        defer func() {
            err := c.close()
            if err != nil {
                panic(errors.Wrap(err))
            }
        }()
        err := c.login()
        if err != nil {
            panic(errors.Wrap(err))
        }
        // 定期检查新邮件
        c.ticker = time.NewTicker(duration)
        for {
            select {
            case <-c.ticker.C:
                // 选择收件箱
                _, err = c.dialer.Select("INBOX", nil).Wait()
                if err != nil {
                    // 尝试修复错误
                    c.attemptToFixErrors(err)
                    continue
                }
                // 使用 UNSEEN 搜索条件获取未读邮件
                searchCriteria := &imap.SearchCriteria{
                    NotFlag: []imap.Flag{imap.FlagSeen},
                }
                //goland:noinspection SpellCheckingInspection
                searchMessages, err := c.dialer.Search(searchCriteria, nil).Wait()
                if err != nil {
                    // 尝试修复错误
                    c.attemptToFixErrors(err)
                    continue
                }
                if len(searchMessages.All) == 0 {
                    log.Println("[INFO - IMAP] 无未读邮件")
                    continue
                }
                log.Printf("[INFO - IMAP] 收到 %d 封邮件\n", len(searchMessages.All))
                // 获取未读邮件的详细信息
                //goland:noinspection SpellCheckingInspection
                seqSet := searchMessages.All
                fetchItems := []imap.FetchItem{
                    imap.FetchItemEnvelope,
                    imap.FetchItemBodyStructure,
                    &imap.FetchItemBodySection{
                        Specifier: imap.PartSpecifierText,
                    },
                }
                var mailSeqSet imap.SeqSet
                if messageBuffers, err := c.dialer.Fetch(seqSet, fetchItems, nil).Collect(); err != nil {
                    // 尝试修复错误
                    c.attemptToFixErrors(err)
                    continue
                } else {
                    // 遍历未读邮件
                    for _, msg := range messageBuffers {
                        mailSeqSet.AddNum(msg.UID)
                        channel <- readMail(msg)
                    }
                }
                // 将未读邮件标记为已读
                flags := &imap.StoreFlags{
                    Flags: []imap.Flag{imap.FlagSeen},
                }
                err = c.dialer.UIDStore(mailSeqSet, flags, nil).Wait()
                if err != nil {
                    // 尝试修复错误
                    c.attemptToFixErrors(err)
                    continue
                }
            }
        }
    }()
    return nil
}

func NewClient(config *EmailConfig) *Client {
    return &Client{
        config: config,
        dialer: nil,
        ticker: nil,
    }
}
