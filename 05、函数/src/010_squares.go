package main

import (
	"fmt"
	"strings"
)

// 010、匿名函数
// go run 010_squares.go
// 输出：
// IBM.:111
// 1
// 4
// 9
// 16
func main() {
	// 拥有函数名的函数只能在包级语法块中被声明，通过函数字面量（function literal）
	// 我们可以绕过这一限制，在任何表达式中表示一个函数值
	// 函数字面量的语法和函数声明相似，区别在于 func 关键字后没有函数名
	// 函数值字面量是一种表达式，它的值被称为匿名函数（anonymous function）
	// 函数字面量允许我们在使用时，再定义它，我们可以改写之前对 strings.Map 的调用
	fmt.Println(strings.Map(func(r rune) rune { return r + 1 }, "HAL-9000")) // IBM.:111
	// 更重要的是，通过这种方式定义的函数可以访问完整的词法环境（lexical environment）
	// 这意味着在函数中定义的内部函数可以引用该函数的变量
	f := squares()
	fmt.Println(f()) // 1
	fmt.Println(f()) // 4
	fmt.Println(f()) // 9
	fmt.Println(f()) // 16
	// squares 的例子证明，函数值不仅仅是一串代码，还记录了状态
	// 在 squares 中定义的匿名内部函数可以访问和更新 squares 中的局部变量
	// 这意味着匿名函数和 squares 中，存在变量引用
	// 这就是函数值属于引用类型和函数值不可比较的原因
	// Go 使用闭包（closures）技术实现函数值，Go 程序员也把函数叫做闭包
	// 通过这个例子，我们看到变量的生命周期不由它的作用域决定，squares 返回后，变量 x 仍然隐式的存在于 f 中
}

// squares 返回一个匿名函数
// 该匿名函数每次被调用时都会返回下一个数的平方
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
