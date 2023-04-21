package action

import (
    "SinguGPT/ai/openai"
    "SinguGPT/models"
    "SinguGPT/utils"
    "strings"
)

func init() {
    RegisterActionFunc(func(sessionId string, _ string, user *models.User, _ models.Contents, contents models.Contents) (models.Contents, error) {
        content := strings.Join(utils.Map(contents.Find(models.TagBody), func(content models.Content) string {
            return content.ToString()
        }), "\n")
        req := openai.NewChatRequest(sessionId, user, content)
        resp, err := openai.Chat(req)
        if err != nil {
            return nil, err
        }
        return models.Contents{
            models.NewMarkdownContent(models.TagBody, resp),
        }, nil
    }, "ChatGPT", "GPT3", "GPT-3", "GPT3.0", "GPT-3.0")
}
