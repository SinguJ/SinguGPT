package imap

import (
    "time"
)

// ContentType 内容类型
type ContentType int

const (
    // Text 文本
    Text ContentType = iota
    // HTML 页面文本
    HTML
    // Other 其他
    Other
)

// Content 内容
type Content struct {
    // 内容类型
    Type ContentType
    // 内容长度
    Len int
    // 内容文本
    Text string
}

// Attach 附件
type Attach struct {
    // 内容类型
    Type ContentType
    // 文件名
    Filename string
    // 文件大小
    Size int
    // 文件数据
    Bytes []byte
}

// Mail 邮件
type Mail struct {
    // ID
    ID string
    // 序列号
    SeqNum uint32
    // 发件人
    From [][2]string
    // 收件人
    To [][2]string
    // 发送日期
    Date time.Time
    // 主题
    Subject string
    // 内容
    Contents []*Content
    // 附件
    Attaches []*Attach
}

func ToTypeName(t ContentType) string {
    switch t {
    case Text:
        return "Text"
    case HTML:
        return "HTML"
    case Other:
        return "Other"
    }
    return ""
}
