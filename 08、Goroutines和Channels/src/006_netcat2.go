package main

import (
	"io"
	"log"
	"net"
	"os"
)

// 006、并发的 Echo 服务 - netcat2
// 先启动服务
// go run 005_reverb1.go&
// 再启动客户端
// go run 006_netcat2.go
// 然后输入内容，得到回响
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go mustCopy2(os.Stdout, conn)
	mustCopy2(conn, os.Stdin)
}

func mustCopy2(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
