package main

import (
	"fmt"
	"time"
)

// 017、基于 select 的多路复用 - countdown1
// go run 017_countdown1.go
func main() {
	fmt.Println("开始倒计时……")
	tick := time.Tick(1 * time.Second) // time.Tick 函数返回一个 channel，程序会周期性的向 channel 发送事件
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
	launch()
}

func launch() {
	fmt.Println("发射！")
}
