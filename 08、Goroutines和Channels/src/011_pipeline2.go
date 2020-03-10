package main

import (
	"fmt"
	"math"
	"time"
)

// 011、Channels - 串联的 Channels 2
func main() {
	// （上接 pipeline1）
	// 在下面的改进中，我们的计数器 goroutine 只生成 100 个含数字的序列，然后关闭 naturals 对应的 channel
	// 这将导致计算平方数的 squarer 对应的 goroutine 可以正常终止循环并关闭 squares 对应的 channel
	// （在一个更复杂的程序中，可以通过 defer 语句关闭对应的 channel）
	// 最后，主 goroutine 也可以正常终止循环并退出程序

	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
			fmt.Print("C")
		}
		close(naturals)
	}()

	// Squarer
	go func() {
		for x := range naturals {
			squares <- x * x
			fmt.Print("S")
		}
		close(squares)
	}()

	// Printer (in main goroutine)
	for x := range squares {
		z := int64(math.Sqrt(float64(x)))
		fmt.Printf("@%dx%d=%d\n", z, z, x)
		time.Sleep(1 * time.Second)
	}

	// 其实你并不需要关闭每一个 channel
	// 只有当需要告诉接收者 goroutine，所有的数据已经全部发送时才需要关闭 channel
	// 不管一个 channel 是否被关闭，当它没有被引用时，将会被 Go 语言的垃圾自动回收器回收
	// （不要将一个关闭打开文件的操作和关闭一个 channel 的操作混淆，
	// 对于每个打开的文件，都要在不使用的时候调用对应的 Close 方法来关闭文件）
	// 试图重复关闭一个 channel 将导致 panic 异常，试图关闭一个 nil 值的 channel 也将导致 panic 异常
	// 关闭一个 channel 时还会触发一个广播机制，我们将在 8.9 节讨论
}
