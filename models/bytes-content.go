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
func (b *BytesContent) Type() ContentType {
    return ContentTypeBytes
}

// Tag 内容标记
func (b *BytesContent) Tag() Tag {
    return b.tag
}

// ExtName 文件扩展名
// 包含 '.' 前缀
func (b *BytesContent) ExtName() string {
    panic("不受支持的方法")
}

// Len 内容长度
func (b *BytesContent) Len() int64 {
    return int64(len(b.bytes))
}

// ToBytes 转为字节数组
func (b *BytesContent) ToBytes() []byte {
    return b.bytes
}

// ToString 转为字符串
func (b *BytesContent) ToString() string {
    return string(b.bytes)
}

// ToReader 转为字节读取流
func (b *BytesContent) ToReader() io.Reader {
    return bytes.NewReader(b.bytes)
}

// NewByteContent 创建字节数组类型的 Content 对象
func NewByteContent(tag Tag, bytes []byte) *BytesContent {
    return &BytesContent{
        tag:   tag,
        bytes: bytes,
    }
}
