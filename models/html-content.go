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
func (t *HTMLContent) Type() ContentType {
    return ContentTypeHTML
}

// Tag 内容标记
func (t *HTMLContent) Tag() Tag {
    return t.tag
}

// ExtName 适合的文件扩展名
func (t *HTMLContent) ExtName() string {
    return ".html"
}

// ToBytes 转为字节数组
func (t *HTMLContent) ToBytes() []byte {
    return []byte(t.html)
}

// ToString 转为字符串
func (t *HTMLContent) ToString() string {
    return t.html
}

// ToReader 转为字节读取流
func (t *HTMLContent) ToReader() io.Reader {
    return bytes.NewBufferString(t.html)
}

// NewHTMLContent 构造 HTML 内容
func NewHTMLContent(tag Tag, html string) *HTMLContent {
    return &HTMLContent{
        tag:  tag,
        html: html,
    }
}
