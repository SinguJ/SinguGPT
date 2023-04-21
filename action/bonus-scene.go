package action

import (
    _ "embed"

    "SinguGPT/models"
)

//go:embed banner.txt
var banner string

func init() {
    RegisterActionFunc(func(_ string, _ string, _ *models.User, _ models.Contents, _ models.Contents) (models.Contents, error) {
        return models.Contents{
            models.NewTextContent(models.TagBody, banner+`

Github: https://github.com/singu-tech/SinguGPT
`),
        }, nil
    }, "Hello-World", "HelloWorld", "Hello World", "你好世界", "你好，世界")
}
