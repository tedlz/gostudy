package main

import (
	"fmt"
	"strings"
)

// 008、函数值
// go run 008_funcvalues.go
// 9
// -3
// func(int) int
// 6
// func(int, int) int
// IBM.:111
// WNT
// Benjy
func main() {
	// 在 Go 中，函数被看做第一类值（first-class values）
	// 函数像其它值一样，拥有类型，可以被赋值给其它变量，传递给函数，从函数返回
	// 对函数值的调用类似函数调用，如下：
	f := square
	fmt.Println(f(3)) // 9
	f = negative
	fmt.Println(f(3))     // -3
	fmt.Printf("%T\n", f) // func(int) int

	// 编译错误，不能把 func(int, int) int 类型分配给 func(int) int
	// f = product // cannot use product (type func(int, int) int) as type func(int) int in assignment

	g := product
	fmt.Println(g(2, 3))  // 6
	fmt.Printf("%T\n", g) // func(int int) int

	var f1 func(int) int
	// f1(3) // panic: runtime error: invalid memory address or nil pointer dereference
	// 函数值可以用来和 nil 比较
	if f1 != nil {
		f1(3)
	}
	// 但函数值之间是不可比较的，也不能用函数值作为 map 的 key

	// 函数值使得我们不仅仅可以通过数据来参数化函数，亦可通过行为
	// 标准库中包含许多这样的例子，下面的代码展示了如何使用
	fmt.Println(strings.Map(add1, "HAL-9000")) // IBM.:111
	fmt.Println(strings.Map(add1, "VMS"))      // WNT
	fmt.Println(strings.Map(add1, "Admix"))    // Benjy
}

func square(n int) int     { return n * n }
func negative(n int) int   { return -n }
func product(m, n int) int { return m * n }

func add1(r rune) rune { return r + 1 }
