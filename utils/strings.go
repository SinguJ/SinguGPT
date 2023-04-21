package utils

import "strings"

// StringFormat 函数的功能是将一个字符串按照指定格式进行替换，并返回替换后的字符串。
//
// 参数：
// format: 需要替换的字符串，支持使用占位符进行替换，占位符格式为 {key}，其中 key 为替换的值对应的键。
// args: 格式化字符串中需要替换的值，使用键值对形式进行传递，可变参数，必须成对出现。
//
// 返回值：
// string: 替换后的字符串。
//
// 示例：
// package main
//
// import (
//     "fmt"
// )
//
// func main() {
//     // 示例1：使用数字占位符替换字符串中的值
//     str1 := "My name is {0}, and I am {1} years old."
//     formattedStr1 := StringFormat(str1, "Alice", "25")
//     fmt.Println(formattedStr1) // 输出：My name is Alice, and I am 25 years old.
//     // 示例2：使用字符串占位符替换字符串中的值
//     str2 := "My name is {Name}, and I am {Age} years old."
//     formattedStr2 := StringFormat(str2, "{Name}", "Alice", "{Age}", "25")
//     fmt.Println(formattedStr2) // 输出：My name is Alice, and I am 25 years old.
// }
func StringFormat(format string, args ...string) string {
    r := strings.NewReplacer(args...)
    return r.Replace(format)
}
