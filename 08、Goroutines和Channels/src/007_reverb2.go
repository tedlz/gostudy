package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

// 007、并发的 Echo 服务 - reverb2
// 先启动服务
// go run 007_reverb2.go&
// 再启动客户端
// go run 006_netcat2.go
// 然后输入内容，得到回响
func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go hanldeConn2(conn)
	}
}

func echo2(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func hanldeConn2(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo2(c, input.Text(), 1*time.Second)
	}
	c.Close()
}
