package models

import (
    "bytes"
    "io"
)

// CommandArgContent 命令参数类型的内容
type CommandArgContent struct {
    // 参数名
    // 为空串时表示位置参数
    name string
    // 参数值
    value string
}

// Type 内容类型
func (c *CommandArgContent) Type() ContentType {
    return ContentTypeText
}

// Tag 内容标记
func (c *CommandArgContent) Tag() Tag {
    return TagCommandArg
}

// Name 参数名
func (c *CommandArgContent) Name() string {
    return c.name
}

// Value 参数名
func (c *CommandArgContent) Value() string {
    return c.value
}

// IsPositionalArg 是否为位置参数
func (c *CommandArgContent) IsPositionalArg() bool {
    return c.name == ""
}

// IsKeyValueArg 是否为 Key-Value 参数
func (c *CommandArgContent) IsKeyValueArg() bool {
    return c.name != ""
}

// Len 内容长度
func (c *CommandArgContent) Len() int64 {
    return int64(len(c.ToString()))
}

// ToBytes 转为字节数组
func (c *CommandArgContent) ToBytes() []byte {
    return []byte(c.ToString())
}

// ToString 转为字符串
func (c *CommandArgContent) ToString() string {
    return c.name + ": " + c.value
}

// ToReader 转为字节读取流
func (c *CommandArgContent) ToReader() io.Reader {
    return bytes.NewBufferString(c.ToString())
}

// NewPositionalArg 创建位置参数的命令参数对象
func NewPositionalArg(value string) *CommandArgContent {
    return &CommandArgContent{
        name:  "",
        value: value,
    }
}

// NewKeyValueArg 创建 Key-Value 参数的命令参数对象
func NewKeyValueArg(name string, value string) *CommandArgContent {
    return &CommandArgContent{
        name:  name,
        value: value,
    }
}
