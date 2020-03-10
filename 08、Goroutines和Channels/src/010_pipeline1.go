package main

import (
	"fmt"
	"math"
	"time"
)

// 010、Channels - 串联的 Channels
func main() {
	// channels 也可以用于将多个 goroutine 连接在一起，一个 channel 的输出作为下一个 channel 的输入
	// 这种串联的 channels 就是所谓的管道（pipeline）
	// 下面的程序用两个 channel 把三个 goroutine 串联起来

	naturals := make(chan int) // channel 1
	squares := make(chan int)  // channel 2

	// Counter (goroutine 1)
	go func() {
		for x := 0; ; x++ {
			naturals <- x
			fmt.Print("C")
		}
	}()

	// Squarer (goroutine 2)
	go func() {
		for {
			x := <-naturals
			squares <- x * x
			fmt.Print("S")
		}
	}()

	// Printer (in main goroutine, goroutine 3)
	for {
		result := <-squares
		x := int64(math.Sqrt(float64(result)))
		fmt.Printf("@%dx%d=%d\n", x, x, result)
		time.Sleep(1 * time.Second)
	}

	// 如果发送者知道，没有更多的值需要发送到 channel 的话，
	// 那么让接收者也能及时知道没有多余的值可接收将是有用的，因为接收者可以停止不必要的接收等待
	// 这可以通过内置的 close 函数来关闭 channel 实现
	// close(naturals)

	// 当一个 channel 被关闭后，再向该 channel 发送数据将导致 panic 异常
	// 当一个被关闭的 channel 中已经发送的数据都被成功接收后，后续的接收操作将不再阻塞，它们会立即返回一个零值
	// 关闭上面例子中的 naturals 变量对应的 channel 并不能终止循环，
	// 它依然会收到一个永无休止的零值序列，然后将它们发送给打印者（Printer）goroutine

	// 没有办法直接测试一个 channel 是否被关闭，但是接收操作有一个变体形式：
	// 它多接收一个结果，多接收的第二个结果是一个布尔值 ok，
	// true 表示成功从 channel 接收到值，false 表示 channel 已经被关闭并且里面没有值可接收

	// 使用这个特性，我们可以修改 Squarer 函数中的循环代码，
	// 当 naturals 对应的 channel 被关闭并且没有值可接收时跳出循环，同时也关闭 squares 对应的 channel
	//
	// Squarer
	// go func() {
	// 	for {
	// 		x, ok := <-naturals
	// 		if !ok {
	// 			break
	// 		}
	// 		squares <- x * x
	// 	}
	// 	close(squares)
	// }()
	//
	// 因为上面的语法是笨拙的，而且这种处理模式很常见，因此 go 语言的 range 循环可直接在 channels 上面迭代
	// 使用 range 循环是上面处理模式的简洁语法，它依次从 channel 接收数据，当 channel 被关闭并且没有值可接收时跳出循环
	// （下转 pipeline2）
}
