package imap

import (
    "fmt"
    "os"
    "time"

    "github.com/emersion/go-imap/v2"
    "github.com/emersion/go-imap/v2/imapclient"
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
            panic(err)
        }
        return nil, err
    }
    return dialer, nil
}

// Login 登录
func (c *Client) login() error {
    if c.dialer != nil && c.dialer.State() == imap.ConnStateLogout {
        err := c.close()
        if err != nil {
            if c.config.Debug {
                panic(err)
            }
            return err
        }
    }
    dialer, err := c.createDialer()
    if err != nil {
        if c.config.Debug {
            panic(err)
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
                panic(err)
            }
            panic(err)
        }
    }(c.dialer)
    err := c.dialer.Logout().Wait()
    if err != nil {
        return err
    }
    c.dialer = nil
    return nil
}

func (c *Client) Read(channel chan *Mail, errorChannel chan error, duration time.Duration) error {
    err := c.login()
    if err != nil {
        if c.config.Debug {
            panic(err)
        }
        return err
    }
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
                panic(err)
            }
        }()

        // 定期检查新邮件
        c.ticker = time.NewTicker(duration)
        for {
            select {
            case <-c.ticker.C:
                // 选择收件箱
                _, err := c.dialer.Select("INBOX", nil).Wait()
                if err != nil {
                    panic(err)
                }
                // 使用 UNSEEN 搜索条件获取未读邮件
                searchCriteria := &imap.SearchCriteria{
                    NotFlag: []imap.Flag{imap.FlagSeen},
                }
                //goland:noinspection SpellCheckingInspection
                searchMessages, err := c.dialer.Search(searchCriteria, nil).Wait()
                if err != nil {
                    panic(err)
                }
                if len(searchMessages.All) == 0 {
                    continue
                }
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
                if messageBuffers, err := c.dialer.Fetch(seqSet, fetchItems, nil).Collect(); err != nil {
                    panic(err)
                } else {
                    // 遍历未读邮件
                    for _, msg := range messageBuffers {
                        channel <- readMail(msg)
                    }
                }
                // 将未读邮件标记为已读
                flags := &imap.StoreFlags{
                    Flags: []imap.Flag{imap.FlagSeen},
                }
                err = c.dialer.UIDStore(seqSet, flags, nil).Wait()
                if err != nil {
                    panic(err)
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
