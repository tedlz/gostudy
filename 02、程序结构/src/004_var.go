package main

import "fmt"

// 004、变量
// go run 004_var.go
// 输出：
// 变量声明的一般语法如下：
// var 变量名 类型 = 表达式

// 零值初始化机制示例：
// var s string
// fmt.Println(s) // ""
// var i int
// fmt.Println(i) // 0

// 简短声明（可用于声明和初始化局部变量）：
// anim := gif.GIF{LoopCount: nframes}
// freq := rand.Float64() * 3.0
// t := 0.0

// 声明一组变量：
// var i, j, k int // int, int, int
// var b, f, s = true, 2.3, "four" // bool, float64, string
// var f, err = os.Open(name) // os.Open returns a file and an error
// i, j := 0, 1

// i := 100 // int 类型
// var boiling float64 = 100 // float64 类型
// var names []string // array 类型，内部元素为 string 类型
// var err error // error 类型
// var p Point // 自定义类型

// i, j = j, i // 交换 i 和 j 的值

// 简短变量声明语句中必须要包含一个新变量，以下代码将编译不通过（同作用域内）：
// f, err := os.Open(infile)
// f, err := os.Create(outfile) // compile error: no new variables
// 正确代码应为：
// in, err := os.Open(infile)
// out, err := os.Create(outfile)
func main() {
	var output = []string{
		"变量声明的一般语法如下：",
		"var 变量名 类型 = 表达式",
		"",
		"零值初始化机制示例：",
		"var s string",
		"fmt.Println(s) // \"\"",
		"var i int",
		"fmt.Println(i) // 0",
		"",
		"简短声明（可用于声明和初始化局部变量）：",
		"anim := gif.GIF{LoopCount: nframes}",
		"freq := rand.Float64() * 3.0",
		"t := 0.0",
		"",
		"声明一组变量：",
		"var i, j, k int // int, int, int",
		"var b, f, s = true, 2.3, \"four\" // bool, float64, string",
		"var f, err = os.Open(name) // os.Open returns a file and an error",
		"i, j := 0, 1",
		"",
		"i := 100 // int 类型",
		"var boiling float64 = 100 // float64 类型",
		"var names []string // array 类型，内部元素为 string 类型",
		"var err error // error 类型",
		"var p Point // 自定义类型",
		"",
		"i, j = j, i // 交换 i 和 j 的值",
		"",
		"简短变量声明语句中必须要包含一个新变量，以下代码将编译不通过（同作用域内）：",
		"f, err := os.Open(infile)",
		"f, err := os.Create(outfile) // compile error: no new variables",
		"正确代码应为：",
		"in, err := os.Open(infile)",
		"out, err := os.Create(outfile)",
	}
	for _, v := range output {
		fmt.Println(v)
	}
}
