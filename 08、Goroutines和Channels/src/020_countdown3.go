package main

import (
	"fmt"
	"os"
	"time"
)

// 020、基于 select 的多路复用 - countdown3
// go run 020_countdown3.go
func main() {
	// 创建用来终止的 channel
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // 读取一个字节
		abort <- struct{}{}
	}()

	fmt.Println("开始倒计时……按 return 来终止")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			// 什么也不做
		case <-abort:
			fmt.Println("发射终止！")
			return
		}
	}
	launch3()

	// time.Tick 函数表现的好像它创建了一个在循环中调用 time.Sleep 的 goroutine，每次被唤醒时发送一个事件
	// 当 countdown 函数返回时，它会停止从 tick 中接收事件，但是 ticker 这个 goroutine 还依然存活，
	// 继续徒劳地尝试向 channel 中发送值，然而这时候已经没有其它的 goroutine 会从该 channel 中接收值了，
	// 这被称为 goroutine 泄露

	// Tick 函数挺方便，但是只有当程序整个生命周期都需要这个时间时，我们使用它才比较合适
	// 否则的话，我们应该使用下面的这种模式
	// ticker := time.NewTicker(1 * time.Second)
	// <-ticker.C    // 从 ticker's channel 接收
	// ticker.Stop() // 使 ticker's goroutine 终止
}

func launch3() {
	fmt.Println("发射！")
}
