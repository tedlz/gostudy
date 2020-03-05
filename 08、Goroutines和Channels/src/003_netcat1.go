package main

import (
	"io"
	"log"
	"net"
	"os"
)

// 003、并发的 Clock 服务 - netcat1
// 需要先 go run 002_clock1.go 或 004_clock2.go
// 再 go run 003_netcat1.go
// 002 为单个 goroutine，004 为多个 goroutine
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
