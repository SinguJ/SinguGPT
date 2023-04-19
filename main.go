package main

import (
    "fmt"
    "log"

    "SinguGPT/access/email"
    "SinguGPT/action"
    "SinguGPT/store"
)

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
    // 挂起主线程
    <-make(chan struct{})
}
