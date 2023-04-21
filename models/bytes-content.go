package models

import (
    "bytes"
    "io"
)

// BytesContent 字节数组类型的内容
type BytesContent struct {
    // 标记
    tag Tag
    // 字节数组
    bytes []byte
}

// Type 内容的类型
func (c *BytesContent) Type() ContentType {
    return ContentTypeBytes
}

// Tag 内容标记
func (c *BytesContent) Tag() Tag {
    return c.tag
}

// Len 内容长度
func (c *BytesContent) Len() int64 {
    return int64(len(c.bytes))
}

// ToBytes 转为字节数组
func (c *BytesContent) ToBytes() []byte {
    return c.bytes
}

// ToString 转为字符串
func (c *BytesContent) ToString() string {
    return string(c.bytes)
}

// ToReader 转为字节读取流
func (c *BytesContent) ToReader() io.Reader {
    return bytes.NewReader(c.bytes)
}

// NewByteContent 创建字节数组类型的 Content 对象
func NewByteContent(tag Tag, bytes []byte) *BytesContent {
    return &BytesContent{
        tag:   tag,
        bytes: bytes,
    }
}
