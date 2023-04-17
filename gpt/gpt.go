package gpt

import (
    "context"
    "errors"
    "fmt"

    "github.com/sashabaranov/go-openai"

    "SinguGPT/models"
)

var client *openai.Client

func Login(apiKey string) {
    config := openai.DefaultConfig(apiKey)
    client = openai.NewClientWithConfig(config)
}

func Chat(sessionKey string, user *models.User, content string) (string, error) {
    ctx := context.Background()

    req := openai.ChatCompletionRequest{
        Model: openai.GPT3Dot5Turbo,
        User:  sessionKey,
        Messages: []openai.ChatCompletionMessage{
            {
                Role:    openai.ChatMessageRoleUser,
                Content: content,
            },
        },
    }
    response, err := client.CreateChatCompletion(ctx, req)
    if err != nil {
        return "", errors.New(fmt.Sprintf("CompletionStream error: %v\n", err))
    }
    return response.Choices[0].Message.Content, nil
}

func ChatStream(sessionKey string, user *models.User, content string) (*openai.ChatCompletionStream, error) {
    ctx := context.Background()

    req := openai.ChatCompletionRequest{
        Model: openai.GPT3Dot5Turbo,
        User:  sessionKey,
        Messages: []openai.ChatCompletionMessage{
            {
                Role:    openai.ChatMessageRoleUser,
                Content: content,
            },
        },
    }
    stream, err := client.CreateChatCompletionStream(ctx, req)
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
