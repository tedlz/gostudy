package main

import (
	"io"
	"log"
	"net"
	"os"
)

// 009、Channels - 不带缓存的 Channels
func main() {
	// 一个基于无缓存 Channels 的发送操作将导致发送者 goroutine 阻塞，
	// 直到另一个 goroutine 在在相同的 Channels 上执行接收操作
	// 当发送的值通过 Channels 成功传输之后，两个 goroutine 可以继续执行后面的语句
	// 反之，如果接收操作先发生，那么接收者 goroutine 也将阻塞，
	// 直到有另一个 goroutine 在相同的 Channels 上执行发送操作

	// 基于无缓存 Channels 的发送和接收操作将导致两个 goroutine 做一次同步操作
	// 因为这个原因，无缓存 Channels 有时候也被称为同步 Channels
	// 当通过一个无缓存的 Channels 发送数据时，接收者收到数据发生在唤醒发送者 goroutine 之前

	// 在讨论并发编程时，当我们说 x 事件在 y 事件之前发生，我们并不是说 x 事件在时间上比 y 时间更早；
	// 我们要表达的意思是要保证在此之前的事件都已经完成了，
	// 例如在此之前的更新某些变量的操作已经完成，你可以放心依赖这些已完成的事件了

	// 当我们说 x 事件既不是在 y 事件之前发生也不是在 y 事件之后发生，我们就说 x 事件和 y 事件是并发的
	// 这并不意味着 x 事件和 y 事件就一定是同时发生的，我们只是不能确定这两个事件发生的先后顺序

	// 在下一章中我们将看到，当两个 goroutine 并发访问了相同的变量时，
	// 我们有必要保证某些事件的执行顺序，以避免出现某些并发问题

	// 在 8.3 节的客户端程序，它在主 goroutine 中，将标准输入复制到 server，
	// 因此当客户端程序关闭标准输入时，后台的 goroutine 可能依然在工作
	// 我们需要让主 goroutine 等待后台 goroutine 完成工作后再退出，我们使用了一个 channel 来同步两个 goroutine：
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // 忽略错误
		log.Println(done)
		done <- struct{}{} // 发信号到主 goroutine
	}()
	mustCopy3(conn, os.Stdin)

	// 原始代码：
	// conn.Close()

	// 练习 8.3：
	cw := conn.(*net.TCPConn) // 类型断言
	cw.CloseWrite()

	<-done // 等待后台的 goroutine 完成

	// 当用户关闭了标准输入，主 goroutine 中的 mustCopy3 函数调用将返回，
	// 然后调用 conn.Close() 关闭读和写方向的网络连接
	// 关闭网络连接中的写方向的连接将导致 server 程序收到一个文件（end-of-file）结束的信号
	// 关闭网络连接中的读方向的连接将导致后台 goroutine 的 io.Copy 函数调用返回一个
	// read from closed connection（从已关闭的连接中读取）类似的错误，因此我们临时移除了错误日志语句
	// 在练习 8.3 将会提供一个更好的解决方案

	// 在后台 goroutine 返回之前，它先打印一个日志信息，然后向 done 对应的 channel 发送一个值
	// 主 goroutine 在退出前先等待从 done 对应的 channel 接收一个值
	// 因此，总是可以在程序退出前正确输出 done 消息

	// 基于 channels 发送消息有两个重要方面
	// 首先每个消息都有一个值，但是有时候通讯的事实和发生的时刻也同样重要
	// 当我们更希望强调通讯发生的时刻时，我们将它称为消息事件
	// 有些消息事件并不携带额外的信息，它仅仅是用作两个 goroutine 之间的同步，
	// 这时候我们可以用 struct{} 空结构体作为 channels 元素的类型，虽然也可以使用 bool 或 int 类型实现同样的功能
	// done <- -1 语句也比 done <- struct{}{} 更短

	// 练习 8.3：
	// 在 netcat3 例子中，conn 虽然是一个 interface 类型的值，但是其底层真实类型是 *net.TCPConn，代表一个 TCP 连接
	// 一个 TCP 连接有读和写两个部分，可以使用 CloseRead 和 CloseWrite 方法分别关闭它们
	// 修改 netcat3 的主 goroutine 代码，只关闭网络连接中写的部分，
	// 这样的话后台 goroutine 可以在标准输入被关闭后继续打印从 reverb1 服务器传回的数据
	// （要在 reverb2 服务器也完成同样的功能是比较困难的，参考练习 8.4）
}

func mustCopy3(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
