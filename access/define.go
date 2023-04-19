package access

import (
    "SinguGPT/models"
)

// MessageHandler 消息处理器
//   Parameters:
//     sessionId => 会话 ID
//     requestId => 请求 ID
//     user      => 发起请求的用户
//     contents  => 用户提供的一组请求内容
//   Returns:
//     resp      => 响应的一组数据
//     err       => 错误
type MessageHandler func(sessionId, requestId string, user *models.User, contents models.Contents) (resp models.Contents, err error)

// Dispatcher 调度器
type Dispatcher interface {
    // OnMessageReceive 接收消息事件
    OnMessageReceive(MessageHandler) error
    // Listen 启动监听
    // 该监听函数不可以阻塞程序继续运行
    Listen() error
}
