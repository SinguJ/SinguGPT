package action

import (
    "SinguGPT/gpt"
    "SinguGPT/models"
)

func init() {
    RegisterActionFunc(func(sessionKey string, user *models.User, content string) (resp string, err error) {
        return gpt.Chat(sessionKey, user, content)
    }, "ChatGPT", "GPT3", "GPT-3", "GPT3.0", "GPT-3.0")
}
