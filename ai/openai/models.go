package openai

import "github.com/sashabaranov/go-openai"

type Model string

const (
    GPT3Dot5 Model = openai.GPT3Dot5Turbo
    GPT4           = openai.GPT4
)
