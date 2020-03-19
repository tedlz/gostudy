package main

import (
	"fmt"
	"os"
	"time"
)

// 018、基于 select 的多路复用 - countdown2
// go run 018_countdown2.go
func main() {
	// 创建用来终止的 channel
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // 读取一个字节
		abort <- struct{}{}
	}()

	fmt.Println("开始倒计时……按 return 来终止")
	select {
	case <-time.After(10 * time.Second):
		// 什么也不做
	case <-abort:
		fmt.Println("发射终止！")
		return
	}
	launch2()
}

func launch2() {
	fmt.Println("发射！")
}
