package action

import (
    "SinguGPT/gpt"
    "SinguGPT/models"
)

func init() {
    RegisterActionFunc(func(sessionId string, requestId string, user *models.User, content string) (resp string, err error) {
        req := gpt.NewChatRequest(sessionId, user, content)
        return gpt.Chat(req)
    }, "ChatGPT", "GPT3", "GPT-3", "GPT3.0", "GPT-3.0")
}
