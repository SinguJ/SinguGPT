package models

import (
    "SinguGPT/errors"
    "io"
)

// ContentType 内容类型
//goland:noinspection GoNameStartsWithPackageName
type ContentType string

const (
    // ContentTypeText 文本内容
    ContentTypeText ContentType = "text"
    // ContentTypeMarkdown 内容
    ContentTypeMarkdown = "markdown"
    // ContentTypeHTML 内容
    ContentTypeHTML = "html"
    // ContentTypeURL 链接
    ContentTypeURL = "url"
    // ContentTypeBytes 字节数组
    ContentTypeBytes = "bytes"
    // ContentTypeFile 文件
    ContentTypeFile = "file"
    // ContentTypeUnknown 未知内容类型
    ContentTypeUnknown = "unknown"
)

// Tag 内容标记
type Tag string

const (
    // TagTitle 标题
    TagTitle Tag = "title"
    // TagCommand 指令
    TagCommand = "command"
    // TagCommandArg 指令参数
    TagCommandArg = "command-arg"
    // TagBody 内容体
    TagBody = "body"
    // TagError 错误
    TagError = "error"
    // TagFile 文件
    TagFile = "file"
)

// Content 内容
// 它表示本程序的输入输出内容
type Content interface {
    // Type 内容的类型
    Type() ContentType
    // Tag 内容标记
    Tag() Tag
    // ExtName 适合的文件扩展名
    // 必须以 '.' 开头，
    // 无适合的扩展名，可以返回空字符串
    ExtName() string
    // Len 内容长度
    Len() int64
    // ToBytes 转为字节数组
    ToBytes() []byte
    // ToString 转为字符串
    ToString() string
    // ToReader 转为字节读取流
    ToReader() io.Reader
}

// Contents 一组内容
type Contents []Content

// Find 根据 Tag 寻找一组内容
func (contents *Contents) Find(tags ...Tag) []Content {
    var newContents []Content
    for _, content := range *contents {
        if content == nil {
            continue
        }
        tag := content.Tag()
        for _, _tag := range tags {
            if tag == _tag {
                newContents = append(newContents, content)
                break
            }
        }
    }
    return newContents
}

// FindOne 根据 Tag 寻找唯一的 Content 对象
func (contents *Contents) FindOne(tag Tag) (Content, error) {
    newContents := contents.Find(tag)
    if len(newContents) != 1 {
        return nil, errors.New("未找到或存在多个 Tag 为 %s 的 Content 对象", tag)
    }
    return newContents[0], nil
}

// MustFindOne 不返回 error 的 FindOne 实现
func (contents *Contents) MustFindOne(tag Tag) Content {
    if content, err := contents.FindOne(tag); err != nil {
        panic(err)
    } else {
        return content
    }
}
