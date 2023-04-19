package action

import (
    _ "embed"

    "SinguGPT/models"
)

//go:embed banner.txt
var banner string

func init() {
    RegisterActionFunc(func(sessionId string, requestId string, user *models.User, content string) (resp string, err error) {
        return banner + `

Github: https://github.com/singu-tech/SinguGPT
`, nil
    }, "Hello-World", "HelloWorld", "Hello World", "你好世界", "你好，世界")
}
