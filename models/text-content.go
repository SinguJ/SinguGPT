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
func (t *TextContent) Type() ContentType {
    return ContentTypeText
}

// Tag 内容标记
func (t *TextContent) Tag() Tag {
    return t.tag
}

// Len 内容长度
func (t *TextContent) Len() int64 {
    return int64(len(t.text))
}

// ToBytes 转为字节数组
func (t *TextContent) ToBytes() []byte {
    return []byte(t.text)
}

// ToString 转为字符串
func (t *TextContent) ToString() string {
    return t.text
}

// ToReader 转为字节读取流
func (t *TextContent) ToReader() io.Reader {
    return bytes.NewBufferString(t.text)
}

// NewTextContent 构造文本内容
func NewTextContent(tag Tag, text string) *TextContent {
    return &TextContent{
        tag:  tag,
        text: text,
    }
}
