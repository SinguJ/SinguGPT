package errors

import (
    "fmt"
    "runtime"
)

// Error 程序错误接口
type Error interface {
    // Error 完整的错误信息
    Error() string
    // Message 错误信息
    Message() string
    // StackTrace 错误调用栈
    StackTrace() string
}

type _error struct {
    // 错误信息
    msg string
    // 错误调用栈
    stack string
}

func (e *_error) Error() string {
    return e.Message() + "\n" + e.StackTrace()
}

func (e *_error) Message() string {
    return e.msg
}

func (e *_error) StackTrace() string {
    return e.stack
}

// New 创建错误对象
func New(msg string, args ...any) Error {
    if len(args) != 0 {
        msg = fmt.Sprintf(msg, args...)
    }
    // 分配足够大的缓冲区以容纳调用栈信息
    buf := make([]byte, 1<<16)
    stackSize := runtime.Stack(buf, true)
    return &_error{
        msg:   msg,
        stack: fmt.Sprintf("%s", buf[:stackSize]),
    }
}

// Wrap 包装异常对象
func Wrap(err any) Error {
    if err == nil {
        return nil
    }
    if _err, ok := err.(Error); ok {
        return _err
    }
    if _err, ok := err.(error); ok {
        return New(_err.Error())
    }
    return New(fmt.Sprintf("%v", err))
}
