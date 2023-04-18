package action

import (
    "SinguGPT/models"
    "SinguGPT/store"
)

func init() {
    // 3.  邮件主题为 "NewSession" 或 "新会话"，邮件内容为 "ChatGPT"
    // 当您需要创建一个新的 ChatGPT 会话时，请将邮件主题设置为 "NewSession" 或 "新会话"，并将邮件内容设置为 "ChatGPT"。程序将为您创建一个全新的会话。
    RegisterActionFunc(func(sessionKey string, user *models.User, content string) (resp string, err error) {
        return stringFormat(`
亲爱的用户，

欢迎使用我们的 {AppName} 程序！为了帮助您更好地使用本程序，我们为您准备了以下简单的操作指南。请仔细阅读以下内容，以便了解如何与我们的 GPT 程序进行互动。
	
1.  向 {AppEmail} 发送邮件
    要与我们的 GPT 程序互动，请向以下电子邮件地址发送邮件：{AppEmail}。我们的程序将会自动处理您发送的邮件。
	
2.  邮件主题为 "ChatGPT"
    如果您希望获得 ChatGPT 的回复，请将邮件标题设置为 "ChatGPT"。程序将识别您的请求，并给您回复 ChatGPT 响应的消息。
	
3.  其他邮件标题
    如果您发送的邮件标题不属于上述任何一种，那么程序将自动回复本帮助信息，以便您随时了解如何使用我们的 GPT 程序。

4.  彩蛋
    我们在 GPT 程序中保留一些彩蛋，这样您才知道您用的是 {AppName}。:)
    您可以尝试不同的邮件标题，或许可以触发某个彩蛋呦~~ 友情提示：程序员、世界
	
希望这些信息能够帮助您更好地使用我们的邮件版 GPT 程序。如果您有任何疑问，请随时联系我们，我们将竭诚为您服务。祝您使用愉快！

{AppName} 团队
`, "{AppName}", store.Config.App.Name, "{AppEmail}", store.Config.Email.IMAP.UserName), nil
    }, "help", "帮助", DefaultAction)
}
