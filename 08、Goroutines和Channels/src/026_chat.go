package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// 026、示例：聊天服务
// go run 026_chat.go&
// go run 009_netcat3.go
// 然后发个消息试试
func main() {
	// 我们用一个聊天服务器来终结本章的内容，这个程序可以让一些用户通过服务器向其它所有用户广播文本消息
	// 这个程序中有四种 goroutine
	// main 和 broadcaster 各自是一个 goroutine 实例，
	// 每一个客户端的连接都会有一个 handleConn 和 clientWriter 的 goroutine
	// broadcaster 是 select 用法的不错的样例，因为它需要处理三种不同类型的消息

	// 下面演示的 main goroutine 的工作，是 listen 和 accept 从客户端过来的连接
	// 对每一个连接，程序都会建立一个新的 handleConn 的 goroutine，就像我们在本章开头的并发的 echo 服务器里所做的那样
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

// 然后是 broadcaster 的 goroutine
// 它的内部变量 clients 会记录当前建立连接的客户端集合
// 其记录的内容是每一个客户端的消息发出 channel 的资格信息

type client chan<- string // 外发消息通道

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // 所有传入的客户端消息
)

func broadcaster() {
	clients := make(map[client]bool) // 所有连接的客户端
	for {
		select {
		case msg := <-messages:
			// 将传入的消息广播到所有客户端
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

// broadcaster 监听来自全局的 entering 和 leaving 的 channel 来获知客户端的到来和离开事件
// 当其接收到其中的一个事件时，会更新 clients 集合，当该事件是离开行为时，它会关闭客户端的消息发出 channel
// broadcaster 也会监听全局的消息 channel，所有的客户端都会向这个 channel 中发送消息
// 当 broadcaster 接收到什么消息时，就会将其广播至所有连接到服务端的客户端

// 现在让我们看看每一个客户端的 goroutine
// handleConn 函数会为它的客户端创建一个消息发出 channel 并通过 entering channel 来通知客户端的到来
// 然后它会读取客户端发出来的每一行文本，并通过全局的消息 channel 来将这些文本发送出去，
// 并为每条消息带上发送者的前缀来标明消息身份
// 当客户端发送完毕后，handleConn 会通过 leaving 这个 channel 来通知客户端的离开并关闭连接

func handleConn(conn net.Conn) {
	ch := make(chan string) // 外发客户端消息
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// 注意：忽略了 input.Err() 的潜在错误

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // 注意：忽略了网络错误
	}
}

// 另外，handleConn 为每一个客户端创建了一个 clientWriter 的 goroutine 来接收向客户端发出消息 channel 中发送的广播消息，
// 并将它们写入到客户端的网络连接。客户端的读取方循环会在 broadcaster 接收到 leaving 通知并关闭了 channel 后终止

// 当与 n 个客户端保持聊天 session 时，这个程序会有 2n+2 个并发的 goroutine，然而这个程序却并不需要显式的锁（见 9.2 节）
// clients 这个 map 被限制在了一个独立的 goroutine 中，broadcaster，所以它不能被并发的访问
// 多个 goroutine 共享的变量只有这些 channel 和 net.Conn 的实例，两个东西都是并发安全的
// 我们会在下一章中更多的解决约束，并发安全以及 goroutine 中共享变量的含义
