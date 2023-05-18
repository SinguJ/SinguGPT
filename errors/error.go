package errors

import (
    "fmt"
    "runtime"
)

// NormalError 一般错误接口
type NormalError interface {
    // Error 完整的错误信息
    Error() string
    // Message 错误信息
    Message() string
}

// ProgramError 程序错误接口
type ProgramError interface {
    NormalError
    // StackTrace 错误调用栈
    StackTrace() string
}

type _error struct {
    // 错误信息
    msg string
}

func (e *_error) Error() string {
    return e.Message()
}

func (e *_error) Message() string {
    return e.msg
}

type _programError struct {
    *_error
    // 错误调用栈
    stack string
}

func (e *_programError) Error() string {
    return e.Message() + "\n" + e.StackTrace()
}

func (e *_programError) Message() string {
    return e._error.Message()
}

func (e *_programError) StackTrace() string {
    return e.stack
}

func newNormalError(msg string, args ...any) *_error {
    if len(args) != 0 {
        msg = fmt.Sprintf(msg, args...)
    }
    return &_error{
        msg: msg,
    }
}

// NewNormalError 创建一般错误对象
func NewNormalError(msg string, args ...any) NormalError {
    return newNormalError(msg, args...)
}

// NewProgramError 创建程序错误对象
func NewProgramError(msg string, args ...any) ProgramError {
    // 分配足够大的缓冲区以容纳调用栈信息
    buf := make([]byte, 1<<16)
    stackSize := runtime.Stack(buf, true)
    return &_programError{
        newNormalError(msg, args...),
        fmt.Sprintf("%s", buf[:stackSize]),
    }
}

// Wrap 包装异常对象
func Wrap(err any) ProgramError {
    if err == nil {
        return nil
    }
    if _err, ok := err.(ProgramError); ok {
        return _err
    }
    if _err, ok := err.(error); ok {
        return NewProgramError(_err.Error())
    }
    return NewProgramError(fmt.Sprintf("%v", err))
}
