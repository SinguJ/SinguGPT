package smtp

import (
    "github.com/go-gomail/gomail"

    "SinguGPT/models"
)

// 添加文本内容
func addText (message *gomail.Message, content *models.TextContent) {
    message.AddAlternative("text/plain", content.ToString())
}

// 添加 HTML 内容
func addHTML (message *gomail.Message, content *models.HTMLContent) {
    message.SetBody("text/html", content.ToString())
}

// 添加 Markdown 内容
func addMarkdown (message *gomail.Message, content *models.MarkdownContent) {
    html := models.MarkdownToHTML(content)
    addHTML(message, html)
}
