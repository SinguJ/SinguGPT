package action

import (
    "SinguGPT/errors"
    "log"
    "strings"

    "SinguGPT/models"
)

// 默认处理器命令
const defaultActionCommand = "*"

// ActionFunc 处理函数
//goland:noinspection GoNameStartsWithPackageName
type ActionFunc func(sessionId string, requestId string, user *models.User, args models.Contents, contents models.Contents) (resp models.Contents, err error)

// Action 处理器
type Action struct {
    Commands []string
    Action   ActionFunc
}

// Actions 处理器注册表
var Actions = make(map[string]*Action)

// RegisterAction 注册一个处理器
func RegisterAction(action *Action) {
    for _, command := range action.Commands {
        command = strings.ToLower(command)
        if Actions[command] != nil {
            log.Fatalln("存在重复的处理关键字")
        }
        Actions[command] = action
    }
}

// RegisterActionFunc 注册一个处理函数
func RegisterActionFunc(af ActionFunc, commands ...string) {
    RegisterAction(&Action{
        Commands: commands,
        Action:   af,
    })
}

// DoAction 执行处理流程
func DoAction(sessionId string, requestId string, user *models.User, req models.Contents) (resp models.Contents, err error) {
    // 根据 Tag 为 Command 类型的 Content，匹配对应的 Action
    command := req.MustFindOne(models.TagCommand).ToString()
    var action *Action
    if action = Actions[command]; action == nil {
        action = Actions[defaultActionCommand]
        if action == nil {
            return nil, errors.New("未知指令：%s", command)
        }
    }
    // 拆分 Command 参数集与普通 Content 数据集
    var commandArgs models.Contents
    var contents models.Contents
    for _, content := range req {
        if content.Tag() == models.TagCommandArg {
            commandArgs = append(commandArgs, content)
        } else {
            contents = append(contents, content)
        }
    }
    // 执行 Action
    if resp, err = action.Action(sessionId, requestId, user, commandArgs, contents); err != nil {
        if err == ErrorNoPermission {
            return nil, errors.New("未知指令：%s", command)
        }
    }
    return
}
