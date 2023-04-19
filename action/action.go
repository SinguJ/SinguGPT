package action

import (
    "log"
    "strings"

    "SinguGPT/models"
)

const DefaultAction = "*"

//goland:noinspection GoNameStartsWithPackageName
type ActionFunc func(sessionId string, requestId string, user *models.User, content string) (resp string, err error)

type Action struct {
    Keywords []string
    Action   ActionFunc
}

var Actions = make(map[string]*Action)

func RegisterAction(action *Action) {
    for _, keyword := range action.Keywords {
        keyword = strings.ToLower(keyword)
        if Actions[keyword] != nil {
            log.Fatalln("存在重复的处理关键字")
        }
        Actions[keyword] = action
    }
}

func RegisterActionFunc(af ActionFunc, keywords ...string) {
    RegisterAction(&Action{
        Keywords: keywords,
        Action:   af,
    })
}

func DoAction(subject string, sessionId string, requestId string, user *models.User, content string) (resp string, err error) {
    keyword := strings.ToLower(strings.TrimSpace(subject))
    var action *Action
    if action = Actions[keyword]; action == nil {
        action = Actions[DefaultAction]
    }
    return action.Action(sessionId, requestId, user, content)
}
