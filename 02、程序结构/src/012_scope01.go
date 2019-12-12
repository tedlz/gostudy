package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// 012、作用域 01
// go run 012_scope01.go
// 输出：
// f
// g
// HELLO
// HELLO
// 2 1
// scope
// scope
func main() {
	f := "f"
	fmt.Println(f) // f，func main 内部变量
	fmt.Println(g) // g，包级变量
	// fmt.Println(h) // 编译错误，undefined: h

	// for1
	x1 := "hello!" // func main 内部变量 x1，显式创建
	for i := 0; i < len(x1); i++ {
		x1 := x1[i] // for 循环内部变量 x1，显式创建
		if x1 != '!' {
			x1 := x1 + 'A' - 'a' // if 内部变量 x1，显式创建
			fmt.Printf("%c", x1) // 迭代输出 HELLO
		}
	}
	fmt.Println()
	// for2
	x2 := "hello!"          // func main 内部变量 x2，显式创建
	for _, x2 := range x2 { // for 循环初始化词法域 x2，隐式创建
		x2 := x2 + 'A' - 'a' // for 循环内部变量 x2，显式创建
		fmt.Printf("%c", x2)
	}
	fmt.Println()
	// if else
	if x := ff(); x == 0 { // if 内部变量 x，隐式创建
		fmt.Println(x) // 如果 ff() 返回 0，则输出 0
	} else if y := gg(x); x == y { // if 内部变量 y，隐式创建
		fmt.Println(x, y) // 如果 ff() 返回 1，则输出 1 1
	} else {
		fmt.Println(x, y) // 如果 ff() 返回 2，则输出 2 1
	}
	// fmt.Println(x, y) // 变量 x, y 的作用域仅限于 if 内部，go run 会报编译错误，undefined x, undefined y

	// if 错误示例
	ifErrDemo()
	// if 正确示例
	ifOkDemo()
}

var g = "g"

func f()             {}
func ff() int8       { return 2 }
func gg(x int8) int8 { return x + 1 - x }
func ifErrDemo() {
	// 	fname := "../files/012_scope.txt"
	// 	if file, err := os.Open(fname); err != nil { // 二、file 没有使用：file declared and not used
	// 		fmt.Println(err)
	// 	}
	// 	file.Close() // 一、编译错误：undefined file，变量 file 不能出 if 使用
}
func ifOkDemo() {
	// 把 file 写到 if 前面
	fname := "../files/012_scope.txt"
	file, err := os.Open(fname)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(contents))

	// 或在 else 里处理 file（不是 Go 语言推荐的做法，推荐 if 中处理错误并返回，这样之后的代码不需要缩进）
	if file, err := os.Open(fname); err != nil {
		fmt.Println(err)
	} else {
		contents, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print(string(contents))
		file.Close()
	}
}
