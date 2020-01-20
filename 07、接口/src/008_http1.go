package main

import (
	"fmt"
	"log"
	"net/http"
)

// 008、http.Handler 接口
// 先启动服务器：go run 008_http1.go&
// 再抓取页面：go run ../../01、入门/src/008_fetch.go http://localhost:8000
// 输出：
// shoes: $50.00
// socks: $5.00
func main() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
