package action

import (
    "SinguGPT/store"
    "fmt"
    "strings"

    "SinguGPT/ai/openai"
    "SinguGPT/models"
    "SinguGPT/utils"
)

func verifyPermission(perm string, user *models.User) bool {
    // 获取用户具有的角色
    role := user.Role
    // 判断该角色是否具有该权限
    for _, _perm := range role.Perms {
        if _perm == perm {
            return true
        }
    }
    return false
}

func toOpenAiModel(model models.OpenAiModel) openai.Model {
    switch model {
    case models.GPT3Dot5:
        return openai.GPT3Dot5
    case models.GPT4:
        return openai.GPT4
    default:
        panic(fmt.Sprintf("不支持的模型名称：%s", model))
    }
}

func newOpenAiAction(client *openai.Client, perm string, model openai.Model, commands ...string) {
    RegisterActionFunc(func(sessionId string, _ string, user *models.User, _ models.Contents, contents models.Contents) (models.Contents, error) {
        // 校验该用户是否具有对该模型的使用权限
        if !verifyPermission(perm, user) {
            return nil, ErrorNoPermission
        }

        content := strings.Join(utils.Map(contents.Find(models.TagBody), func(content models.Content) string {
            return content.ToString()
        }), "\n")
        req := openai.NewChatRequest(model, sessionId, user, content)
        resp, err := client.Chat(req)
        if err != nil {
            return nil, err
        }
        // 优化响应内容
        resp = utils.BeautifulMarkdown(resp)
        markdown := models.NewMarkdownContent(models.TagBody, resp)
        return models.Contents{
            markdown,
            models.NewFileContent("原文.md", markdown),
        }, nil
    }, commands...)
}

func init() {
    for _, config := range store.Config.OpenAI {
        client := openai.NewClient(config.OrgId, config.ApiKey)
        for _, model := range config.Models {
            newOpenAiAction(client, config.ID, toOpenAiModel(model), config.Commands...)
        }
    }
}
