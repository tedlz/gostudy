package main

import (
	"fmt"
	"os"
	"time"
)

// 021、基于 select 的多路复用 - countdown4，对 countdown3 的改造，防止 goroutine 泄露
// go run 021_countdown4.go
func main() {
	// 创建用来终止的 channel
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // 读取一个字节
		abort <- struct{}{}
	}()

	fmt.Println("开始倒计时……按 return 来终止")
	// tick := time.Tick(1 * time.Second) // -
	ticker := time.NewTicker(1 * time.Second) // +
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		// case <-tick:  // -
		case <-ticker.C: // +
			// 什么也不做
		case <-abort:
			fmt.Println("发射终止！")
			ticker.Stop() // +
			return
		}
	}
	launch4()

	// 下面的 select 语句会在 abort channel 中有值时，从其中接收值；无值时什么都不做
	// 这是一个非阻塞的接收操作；反复做这样的操作叫做轮询 channel
	// select {
	// case <-abort:
	// 	   fmt.Println("发射终止！")
	// 	   return
	// default:
	//     什么都做不做
	// }

	// channel 的零值是 nil，对一个 nil 的 channel 发送和接收操作会永远阻塞，
	// 在 select 语句中操作 nil 的 channel 永远都不会被 select 到
	// 这使得我们可以用 nil 来激活或禁用 case，来达成处理其它输入或输出事件时，超时和取消的逻辑
	// 我们会在下一节中看到一个例子
}

func launch4() {
	fmt.Println("发射！")
}
