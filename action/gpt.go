package action

import (
    "SinguGPT/gpt"
    "SinguGPT/models"
)

func init() {
    RegisterActionFunc(func(sessionKey string, user *models.User, content string) (resp string, err error) {
        req := gpt.NewChatRequest(sessionKey, user, content)
        return gpt.Chat(req)
    }, "ChatGPT", "GPT3", "GPT-3", "GPT3.0", "GPT-3.0")
}
