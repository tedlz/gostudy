package main

import (
	"fmt"
	"math"
)

// 001、函数声明
// go run 001_declare.go
// 输出:
// 5
// func(int, int) int
// func(int, int) int
// func(int, int) int
// func(int, int) int
func main() {
	fmt.Println(hypot(3, 4)) // 5

	fmt.Printf("%T\n", add)   // func(int, int) int
	fmt.Printf("%T\n", sub)   // func(int, int) int
	fmt.Printf("%T\n", first) // func(int, int) int
	fmt.Printf("%T\n", zero)  // func(int, int) int
}

// 函数声明包括函数名、形式参数列表、返回值列表（可省略），以及函数体
// func name(parameter-list) (result-list) {
// 	body
// }
// 形式参数列表描述了函数的参数名以及参数类型
// 返回值列表描述了函数返回值的变量名以及类型
func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

// 如果一组形参或返回值有相同的类型，我们不必为每个形参都写出参数类型
// 以下两个声明是等价的：
// func f(i, j, k int, s, t string)                {}
// func f(i int, j int, k int, s string, t string) {}

// 用四种方法声明拥有 2 个 int 型参数和 1 个 int 型返回值的函数（空格标识符 _ 可以强调某个参数未被使用）
func add(x int, y int) int   { return x + y }
func sub(x, y int) (z int)   { z = x - y; return }
func first(x int, _ int) int { return x }
func zero(int, int) int      { return 0 }

// 函数的类型被称为函数的标识符
// 你可能会偶尔遇到没有函数体的函数声明，这表示该函数不是以 Go 实现的。这样的声明定义了函数标识符
// package math
// func Sin(x float64) float // 用汇编语言实现
