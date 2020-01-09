package main

import "fmt"

// 018、panic 异常
// go run 019_panic.go
// 输出：
// f(3)
// f(2)
// f(1)
// defer 1
// defer 2
// defer 3
// panic: runtime error: integer divide by zero

// goroutine 1 [running]:
// main.f(0x0)
//         /data/go/src/gostudy/05、函数/src/018_panic.go:47 +0x1be
// main.f(0x1)
//         /data/go/src/gostudy/05、函数/src/018_panic.go:49 +0x18d
// main.f(0x2)
//         /data/go/src/gostudy/05、函数/src/018_panic.go:49 +0x18d
// main.f(0x3)
//         /data/go/src/gostudy/05、函数/src/018_panic.go:49 +0x18d
// main.main()
//         /data/go/src/gostudy/05、函数/src/018_panic.go:34 +0x2a
// exit status 2
func main() {
	// Go 的类型系统会在编译时捕获很多错误，但有些错误只能在运行时检查，如数组访问越界、空指针引用等
	// 这些运行时错误会引起 panic 异常

	// 一般而言，当 panic 异常发生时，程序会中断运行，并立即执行在该 goroutine 中被延迟的函数（defer 机制）
	// 随后，程序崩溃并输出日志信息，日志信息包括 panic value 和函数调用的堆栈跟踪信息，panic value 通常是某种错误信息
	// 对于每个 goroutine，日志信息中都会有与之相对的，发生 panic 时的函数调用堆栈跟踪信息
	// 通常，我们不需要再次运行程序去定位问题，日志信息已经提供了足够的诊断依据
	// 因此，在我们填写问题报告时，一般会将 panic 异常和日志信息一并记录

	// 不是所有的 panic 异常都来自于运行时，直接调用内置的 panic 函数也会引发 panic 异常
	// panic 函数接受任何值作为参数
	// switch suit := "unknown"; suit {
	// case "A":
	// case "B":
	// case "C":
	// case "D":
	// 	   fmt.Println("bingo")
	// default:
	// 	   panic(fmt.Sprintf("invalid suit %q", suit)) // panic: invalid suit "unknown"
	// }

	// 虽然 Go 的 panic 机制类似于其它语言的异常，但 panic 的适用场景有些不同
	// 由于 panic 会引起程序崩溃，因此 panic 一般都用于严重错误，如程序内部的逻辑不一致
	// 对于大部分漏洞，我们应该使用 Go 提供的错误机制，而不是 panic，尽量避免程序的崩溃
	// 在健壮的程序中，任何可预见的错误，如不正确的输入、错误的配置或是失败的 I/O 操作都应该被优雅的处理
	// 最好的方式，就是使用 Go 的错误机制

	f(3)
}

// 断言函数必须满足前置条件是明智的做法，但这很容易被滥用，除非你能提供更多的错误信息，或能更快速的发现错误
// 否则不需要使用断言，编译器在运行时会帮你检查代码
// func reset(x *Buffer) {
// 	   if x == nil {
// 		   panic("x is nil") // 没必要
// 	   }
// 	   x.elements = nil
// }

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}
