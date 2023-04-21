package models

import (
    "bytes"
    "io"
)

// HTMLContent 文本类型内容
type HTMLContent struct {
    // 标签
    tag Tag
    // HTML 内容
    html string
}

// Type 内容类型
func (c *HTMLContent) Type() ContentType {
    return ContentTypeHTML
}

// Tag 内容标记
func (c *HTMLContent) Tag() Tag {
    return c.tag
}

// Len 内容长度
func (c *HTMLContent) Len() int64 {
    return int64(len(c.html))
}

// ToBytes 转为字节数组
func (c *HTMLContent) ToBytes() []byte {
    return []byte(c.html)
}

// ToString 转为字符串
func (c *HTMLContent) ToString() string {
    return c.html
}

// ToReader 转为字节读取流
func (c *HTMLContent) ToReader() io.Reader {
    return bytes.NewBufferString(c.html)
}

// NewHTMLContent 构造 HTML 内容
func NewHTMLContent(tag Tag, html string) *HTMLContent {
    return &HTMLContent{
        tag:  tag,
        html: html,
    }
}
