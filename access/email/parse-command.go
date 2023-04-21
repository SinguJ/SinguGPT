package email

import (
    "SinguGPT/models"
    "strings"
)

// 清理命令字符串
func clearCommandStr(str string) string {
    return strings.ToLower(strings.TrimSpace(str))
}

// 解析命令
func parseCommand(subject string) models.Contents {
    var contents models.Contents
    // 为了便于处理，将 `：` 替换为 `:`，
    // 将 `，` 替换为 `,`，
    subject = strings.ReplaceAll(strings.ReplaceAll(subject, "：", ":"), "，", ",")
    // 寻找 `:` 字符的位置
    if colonIndex := strings.IndexRune(subject, ':'); colonIndex == -1 {
        // 若找不到 `:`，则说明 subject 就是命令，且为无参调用
        contents = append(contents, models.NewTextContent(models.TagCommand, clearCommandStr(subject)))
    } else {
        // 若找到了 `:`，则说明本次为有参调用
        command := models.NewTextContent(models.TagCommand, clearCommandStr(subject[:colonIndex]))
        // 使用 `,` 拆分参数
        if args := strings.Split(subject[colonIndex+1:], ","); len(args) == 0 {
            // 无参数
            contents = append(contents, command)
        } else {
            // 根据参数数量构造 Contents
            contents := make(models.Contents, 1+len(args))
            // 记录 Command
            contents[0] = command
            // 循环处理每一个参数
            for argIndex, arg := range args {
                // 寻找 '=' 字符的位置
                if equalSignIndex := strings.IndexRune(arg, '='); equalSignIndex == -1 {
                    // 若找不到，说明该参数是位置参数
                    contents[1+argIndex] = models.NewPositionalArg(strings.TrimSpace(arg))
                } else {
                    // 若能找到，说明该参数是 Key-Value 参数
                    contents[1+argIndex] = models.NewKeyValueArg(
                        clearCommandStr(arg[:equalSignIndex]),
                        strings.TrimSpace(arg[equalSignIndex+1:]),
                    )
                }
            }
        }
    }
    return contents
}
