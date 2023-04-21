package smtp

import (
    "github.com/go-gomail/gomail"

    "SinguGPT/models"
)

// 添加文本内容
func addText (message *gomail.Message, content *models.TextContent) {
    message.AddAlternative("text/plain", content.ToString())
}
