package action

import (
    "SinguGPT/ai/openai"
    "SinguGPT/store"
)

func init() {
    openai.Login(store.Config.OpenAI.ApiKey)
}
