package models

import (
    "bytes"
    "io"
)

// TextContent 文本类型内容
type TextContent struct {
    tag  Tag
    text string
}

// Type 内容类型
func (c *TextContent) Type() ContentType {
    return ContentTypeText
}

// Tag 内容标记
func (c *TextContent) Tag() Tag {
    return c.tag
}

// Len 内容长度
func (c *TextContent) Len() int64 {
    return int64(len(c.text))
}

// ToBytes 转为字节数组
func (c *TextContent) ToBytes() []byte {
    return []byte(c.text)
}

// ToString 转为字符串
func (c *TextContent) ToString() string {
    return c.text
}

// ToReader 转为字节读取流
func (c *TextContent) ToReader() io.Reader {
    return bytes.NewBufferString(c.text)
}

// NewTextContent 构造文本内容
func NewTextContent(tag Tag, text string) *TextContent {
    return &TextContent{
        tag:  tag,
        text: text,
    }
}
