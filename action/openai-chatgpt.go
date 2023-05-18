package action

import (
    "regexp"
    "strings"

    "SinguGPT/ai/openai"
    "SinguGPT/models"
    "SinguGPT/utils"
)

// 正则规则
var regexRules = []struct {
    pattern *regexp.Regexp
    replace string
}{
    // 在英文字符与汉字之间插入空格
    {regexp.MustCompile(`([a-zA-Z])(\p{Han})`), "$1 $2"},
    {regexp.MustCompile(`(\p{Han})([a-zA-Z])`), "$1 $2"},
    // 在英文符号与汉字之间插入空格
    {regexp.MustCompile(`([[:punct:]])(\p{Han})`), "$1 $2"},
    {regexp.MustCompile(`(\p{Han})([[:punct:]])`), "$1 $2"},
    // 在数字与汉字之间插入空格
    {regexp.MustCompile(`([0-9])(\p{Han})`), "$1 $2"},
    {regexp.MustCompile(`(\p{Han})([0-9])`), "$1 $2"},
}

// 内容优化
func responseOptimization(input string) string {
    // 通过正则表达式规则集优化输入的 Markdown 原文
    output := input
    for _, rule := range regexRules {
        output = rule.pattern.ReplaceAllString(output, rule.replace)
    }
    return output
}

func init() {
    RegisterActionFunc(func(sessionId string, _ string, user *models.User, _ models.Contents, contents models.Contents) (models.Contents, error) {
        content := strings.Join(utils.Map(contents.Find(models.TagBody), func(content models.Content) string {
            return content.ToString()
        }), "\n")
        req := openai.NewChatRequest(openai.GPT3Dot5, sessionId, user, content)
        resp, err := openai.Chat(req)
        if err != nil {
            return nil, err
        }
        // 优化响应内容
        resp = responseOptimization(resp)
        markdown := models.NewMarkdownContent(models.TagBody, resp)
        return models.Contents{
            markdown,
            models.NewFileContent("原文.md", markdown),
        }, nil
    }, "ChatGPT", "GPT3_5", "GPT-3", "GPT3_5.0", "GPT-3.0")
}
