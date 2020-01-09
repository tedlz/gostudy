package main

import (
	"fmt"
	"os"
	"runtime"
)

// 019、panic 异常
// go run 019_defer2.go
// 输出：
// f(3)
// f(2)
// f(1)
// defer 1
// defer 2
// defer 3
// goroutine 1 [running]:
// main.printStack()
//         /data/go/src/gostudy/05、函数/src/019_defer2.go:20 +0x5b
// panic(0x4a8320, 0x5598a0)
//         /usr/local/go/src/runtime/panic.go:679 +0x1b2
// main.f(0x0)
//         /data/go/src/gostudy/05、函数/src/019_defer2.go:25 +0x1be
// main.f(0x1)
//         /data/go/src/gostudy/05、函数/src/019_defer2.go:27 +0x18d
// main.f(0x2)
//         /data/go/src/gostudy/05、函数/src/019_defer2.go:27 +0x18d
// main.f(0x3)
//         /data/go/src/gostudy/05、函数/src/019_defer2.go:27 +0x18d
// main.main()
//         /data/go/src/gostudy/05、函数/src/019_defer2.go:15 +0x50
// panic: runtime error: integer divide by zero

// goroutine 1 [running]:
// main.f(0x0)
//         /data/go/src/gostudy/05、函数/src/019_defer2.go:25 +0x1be
// main.f(0x1)
//         /data/go/src/gostudy/05、函数/src/019_defer2.go:27 +0x18d
// main.f(0x2)
//         /data/go/src/gostudy/05、函数/src/019_defer2.go:27 +0x18d
// main.f(0x3)
//         /data/go/src/gostudy/05、函数/src/019_defer2.go:27 +0x18d
// main.main()
//         /data/go/src/gostudy/05、函数/src/019_defer2.go:15 +0x50
// exit status 2
func main() {
	defer printStack()
	f(3)
}

// runtime.Stack 为什么能输出已经被释放函数的信息？
// 在 Go 的 panic 机制中，延迟函数的调用在释放堆栈信息之前
func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}
