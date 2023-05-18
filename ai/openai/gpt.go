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

const SystemContent = `ASSISTANT's response must be in Markdown syntax format.`

// Client OpenAI 客户端
type Client struct {
    orgId  string
    apiKey string

    client *openai.Client
}

func (c *Client) login() {
    config := openai.DefaultConfig(c.apiKey)
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
    c.client = openai.NewClientWithConfig(config)
}

func (c *Client) Chat(req *openai.ChatCompletionRequest) (string, error) {
    ctx := context.Background()
    response, err := c.client.CreateChatCompletion(ctx, *req)
    if err != nil {
        return "", errors.New(fmt.Sprintf("CompletionStream error: %v\n", err))
    }
    return response.Choices[0].Message.Content, nil
}

func (c *Client) ChatStream(req *openai.ChatCompletionRequest) (*openai.ChatCompletionStream, error) {
    ctx := context.Background()
    stream, err := c.client.CreateChatCompletionStream(ctx, *req)
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

func NewClient(orgId string, apiKey string) *Client {
    client := &Client{
        orgId:  orgId,
        apiKey: apiKey,
    }
    client.login()
    return client
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
