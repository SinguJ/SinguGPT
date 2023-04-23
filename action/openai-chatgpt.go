package action

import (
    "regexp"
    "strings"

    "SinguGPT/ai/openai"
    "SinguGPT/models"
    "SinguGPT/utils"
)

var __front__ = regexp.MustCompile(`(\p{Han}+)(\W+)`)
var __behind__ = regexp.MustCompile(`(\W+)(\p{Han}+)`)

// 在汉字与英语单词之间添加空格
func addSpacesBetweenChineseCharactersAndEnglishWords(str string) string {
    return __front__.ReplaceAllString(__behind__.ReplaceAllString(str, "$1 $2"), "$1 $2")
}

// 内容优化
func responseOptimization(resp string) string {
    return addSpacesBetweenChineseCharactersAndEnglishWords(resp)
}

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
        // 优化响应内容
        resp = responseOptimization(resp)
        markdown := models.NewMarkdownContent(models.TagBody, resp)
        return models.Contents{
            markdown,
            models.NewFileContent("Reply.md", markdown),
        }, nil
    }, "ChatGPT", "GPT3", "GPT-3", "GPT3.0", "GPT-3.0")
}
