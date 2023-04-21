package smtp

import (
    "github.com/go-gomail/gomail"
    "io"

    "SinguGPT/models"
)

// 添加文本内容
func addText(message *gomail.Message, content *models.TextContent) {
    message.AddAlternative("text/plain", content.ToString())
}

// 添加 HTML 内容
func addHTML(message *gomail.Message, content *models.HTMLContent) {
    message.SetBody("text/html", content.ToString())
}

// 添加 Markdown 内容
func addMarkdown(message *gomail.Message, content *models.MarkdownContent) {
    html := models.MarkdownToHTML(content)
    addHTML(message, html)
}

// 添加消息内容
func addMessageContent(message *gomail.Message, content models.Content) {
    // 根据 Content 的类型，执行相应的处理方式
    switch content.Type() {
    case models.ContentTypeText:
        addText(message, content.(*models.TextContent))
    case models.ContentTypeHTML:
        addHTML(message, content.(*models.HTMLContent))
    case models.ContentTypeMarkdown:
        addMarkdown(message, content.(*models.MarkdownContent))
    default:
    }
}

// 添加附件
func addAttach(message *gomail.Message, filename string, content models.Content) {
    reader := content.ToReader()
    message.Attach(filename, gomail.SetCopyFunc(func(writer io.Writer) error {
        _, err := io.Copy(writer, reader)
        return err
    }))
}
