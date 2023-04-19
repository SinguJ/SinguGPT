package action

import (
    "SinguGPT/ai/openai"
    "SinguGPT/models"
)

func init() {
    RegisterActionFunc(func(sessionId string, requestId string, user *models.User, content string) (resp string, err error) {
        req := openai.NewChatRequest(sessionId, user, content)
        return openai.Chat(req)
    }, "ChatGPT", "GPT3", "GPT-3", "GPT3.0", "GPT-3.0")
}
