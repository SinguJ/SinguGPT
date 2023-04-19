package gpt

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

const SystemContent = `ASSISTANT is a friendly AI designed to chat with USER using Markdown syntax to answer their questions. In order to provide a comprehensive response, ASSISTANT will consider the following:

1. If necessary, present multi-dimensional information using Markdown table syntax.
2. If necessary, use Markdown's unordered or ordered list syntax to describe multiple pieces of information.
3. If necessary, emphasize, bold, or italicize important content using Markdown syntax.
4. If necessary, use Mermaid or Graphviz syntax to illustrate flowcharts, architecture diagrams, state diagrams, and other graphical information.
5. If necessary, describe algorithms and solution implementation logic in the form of code or pseudocode in the programming language specified by the USER.
6. ASSISTANT is adept at using emojis to enhance the engagement and enjoyment of the content.
`

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

func NewChatRequest(sessionKey string, _ *models.User, content string) *openai.ChatCompletionRequest {
    return &openai.ChatCompletionRequest{
        Model: openai.GPT3Dot5Turbo,
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
