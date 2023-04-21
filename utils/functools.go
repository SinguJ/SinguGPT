package utils

// Map 函数将一个函数应用于一个切片的每个元素，并返回一个新的切片，包含应用函数后的结果。
//
// 参数：
// arr: 输入切片。
// fn: 应用于每个元素的函数。
//
// 返回值：
// []R: 包含应用函数后的结果的新切片。
//
// 示例：
// package main
//
// import (
//     "fmt"
// )
//
// func main() {
//     // 定义一个整数切片
//     numbers := []int{1, 2, 3, 4, 5}
//
//     // 使用 Map 函数对切片中的每个元素进行平方操作
//     squares := Map(numbers, func(n int) int {
//         return n * n
//     })
//
//     // 输出结果
//     fmt.Println(squares) // 输出：[1 4 9 16 25]
// }
func Map[T, R any](arr []T, fn func(T) R) []R {
    result := make([]R, len(arr))
    for i, v := range arr {
        result[i] = fn(v)
    }
    return result
}

// Reduce 函数对一个切片中的所有元素进行累积操作，返回一个最终结果。
//
// 参数：
// arr: 输入切片。
// fn: 用于累积的函数，输入参数为上一次累积的结果和当前元素，输出参数为累积结果。
// initial: 累积结果的初始值。
//
// 返回值：
// R: 累积操作的最终结果。
//
// 示例：
// package main
//
// import "fmt"
//
// func main() {
//     // 定义一个整数切片
//     numbers := []int{1, 2, 3, 4, 5}
//
//     // 使用 Reduce 函数对切片中的每个元素进行累积操作，计算切片中所有元素的和
//     sum := Reduce(numbers, func(acc, n int) int {
//         return acc + n
//     }, 0)
//
//     // 输出结果
//     fmt.Println(sum) // 输出：15
// }
func Reduce[T any, R any](arr []T, fn func(R, T) R, initial R) R {
    result := initial
    for _, v := range arr {
        result = fn(result, v)
    }
    return result
}
