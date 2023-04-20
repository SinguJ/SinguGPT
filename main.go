package main

import (
    "fmt"
    "log"

    "SinguGPT/access/email"
    "SinguGPT/action"
    "SinguGPT/models"
    "SinguGPT/store"
)

// 接入方式映射表
var accessMappingTable = make(map[models.AccessMethod]func())

// 启用电子邮件接入
func accessEmail () {
    // 创建邮件调度器
    emailDispatcher := email.NewDispatcher()
    err := emailDispatcher.OnMessageReceive(action.DoAction)
    if err != nil {
        log.Fatalf("[ERROR] %v", err)
    }
    // 启动监听
    err = emailDispatcher.Listen()
    if err != nil {
        log.Fatalf("[ERROR] %v", err)
    }
}

func init () {
    // 添加邮箱接入方式
    accessMappingTable[models.AccessMethodEmail] = accessEmail
}

func main() {
    println(fmt.Sprintf("应用程序名称：%s\n", store.Config.App.Name))
    defer func() {
        err := recover()
        if err != nil {
            log.Fatalf("[ERROR] %v", err)
        }
    }()
    // 加载并监听用户配置文件
    store.LoadAndWatchUsers()
    // 初始化接入方式
    for _, accessMethod := range store.Config.App.AccessMethods {
        if bootFunc := accessMappingTable[accessMethod]; bootFunc != nil {
            bootFunc()
        } else {
            log.Fatalf("[ERROR] 无效的接入方式：%s", accessMethod)
        }
    }
    // 挂起主线程
    <-make(chan struct{})
}
