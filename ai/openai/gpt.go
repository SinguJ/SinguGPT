package openai

import (
    "context"
    "errors"
    "fmt"
    "log"
    "net/http"
    "net/url"
    "os"

    "github.com/sashabaranov/go-openai"

    "SinguGPT/models"
)

var client *openai.Client

const SystemContent = `ASSISTANT's response must be in Markdown syntax format.`

func Login(apiKey string) {
    config := openai.DefaultConfig(apiKey)
    if proxyServ := os.Getenv("SINGU_GPT_PROXY"); proxyServ != "" {
        log.Printf("使用代理服务器：%s\n", proxyServ)
        _proxyUrl, err := url.Parse(proxyServ)
        if err != nil {
            panic(err)
        }
        httpClient := &http.Client{
            Transport: &http.Transport{
                Proxy: http.ProxyURL(_proxyUrl),
            },
        }
        config.HTTPClient = httpClient
    }
    client = openai.NewClientWithConfig(config)
}

func NewChatRequest(model Model, sessionKey string, _ *models.User, content string) *openai.ChatCompletionRequest {
    return &openai.ChatCompletionRequest{
        Model: string(model),
        User:  sessionKey,
        Messages: []openai.ChatCompletionMessage{
            {
                Role:    openai.ChatMessageRoleSystem,
                Content: SystemContent,
            },
            {
                Role:    openai.ChatMessageRoleUser,
                Content: content,
            },
        },
    }
}

func Chat(req *openai.ChatCompletionRequest) (string, error) {
    ctx := context.Background()
    response, err := client.CreateChatCompletion(ctx, *req)
    if err != nil {
        return "", errors.New(fmt.Sprintf("CompletionStream error: %v\n", err))
    }
    return response.Choices[0].Message.Content, nil
}

func ChatStream(req *openai.ChatCompletionRequest) (*openai.ChatCompletionStream, error) {
    ctx := context.Background()
    stream, err := client.CreateChatCompletionStream(ctx, *req)
    if err != nil {
        return nil, errors.New(fmt.Sprintf("CompletionStream error: %v\n", err))
    }
    return stream, nil
    // defer stream.Close()
    // for {
    //	response, err := stream.Recv()
    //	if errors.Is(err, io.EOF) {
    //		return
    //	}
    //	if err != nil {
    //		panic(fmt.Sprintf("Stream error: %v\n", err))
    //	}
    //	fmt.Printf("%v", response.Choices[0].Delta.Content)
    // }
}
