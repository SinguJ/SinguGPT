package models

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

// ToString 转为字符串
func (t *TextContent) ToString() string {
    return t.text
}

// NewTextContent 构造文本内容
func NewTextContent(tag Tag, text string) *TextContent {
    return &TextContent{
        tag:  tag,
        text: text,
    }
}
