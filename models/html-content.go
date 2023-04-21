package models

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

// ToString 转为字符串
func (t *HTMLContent) ToString() string {
    return t.html
}

// NewHTMLContent 构造 HTML 内容
func NewHTMLContent(tag Tag, html string) *HTMLContent {
    return &HTMLContent{
        tag:  tag,
        html: html,
    }
}
