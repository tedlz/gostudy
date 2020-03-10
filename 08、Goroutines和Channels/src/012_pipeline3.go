package main

import (
	"fmt"
	"math"
	"time"
)

// 012、Channels - 单方向的 Channel
func main() {
	// 随着程序的增长，人们习惯于将大的函数拆分为小的函数
	// 我们前面的例子中使用了三个 goroutine，然后用两个 channel 来连接它们，它们都是 main 函数的局部变量
	// 将三个 goroutine 拆分为以下三个函数是自然的想法：
	//
	// func counter(out chan int)
	// func squarer(out, in chan int)
	// func printer(in chan int)
	//
	// 其中计算平方的 squarer 函数在两个串联的 channel 的中间，因此拥有两个 channel 类型的参数，一个用于输入，一个用于输出
	// 两个 channel 都拥有相同的类型，但是它们的使用方式相反：一个只用于接收，另一个只用于发送
	// 参数的名字 in 和 out 已经明确表示了这个意图，
	// 但是并无法保证 squarer 函数向一个 in 参数对应的 channel 发送数据或者从一个 out 参数对应的 channel 接收数据

	// 这种场景是典型的，当一个 channel 作为一个函数参数时，它一般总是被专门用于只发送或者只接收

	// 为了表明这种意图并防止被滥用，Go 语言的类型系统提供了单方向的 channel 类型，分别用于只发送或只接收的 channel
	// 类型 chan<- int 表示一个只发送 int 的 channel，只能发送不能接收
	// 相反，类型 <-chan int 表示一个只接收 int 的 channel，只能接收不能发送
	// （箭头 <- 和关键字 chan 的相对位置表明了 channel 的方向）
	// 这种限制将在编译期检测

	// 因为关闭操作只用于断言不再向 channel 发送新的数据，所以只有在发送者所在的 goroutine 才会调用 close 函数
	// 因此对一个只接收的 channel 调用 close 将是一个编译错误

	// 这是改进的版本，这一次参数使用了单方向 channel 类型：
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)

	// 调用 counter (naturals) 时，naturals 的类型将隐式的从 chan int 转换成 chan<- int
	// 调用 printer 也会导致相似的隐式转换，这一次是转换为 <-chan int 类型只接收型的 channel
	// 任何双向 channel 向单向 channel 变量的赋值操作都将导致该隐式转换

	// 这里并没有反向转换的语法，
	// 也就是不能将一个类似 chan<- int 类型的单向型的 channel 转换为 chan int 类型的双向型的 channel
}

func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
		fmt.Print("C")
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
		fmt.Print("S")
	}
	close(out)
}

func printer(in <-chan int) {
	for x := range in {
		z := int64(math.Sqrt(float64(x)))
		fmt.Printf("@%dx%d=%d\n", z, z, x)
		time.Sleep(1 * time.Second)
	}
}
