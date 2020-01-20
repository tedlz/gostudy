package main

import (
	"fmt"
	"log"
	"net/http"
)

// 009、http.Handler 接口 2
// 先启动服务器：go run 009_http2.go&
// 再抓取页面：go run ../../01、入门/src/008_fetch.go http://localhost:8000/list
// 输出：
// shoes: $50.00
// socks: $5.00
//
// go run ../../01、入门/src/008_fetch.go http://localhost:8000/price?item=shoes
// 输出：
// $50.00
//
// go run ../../01、入门/src/008_fetch.go http://localhost:8000/help
// 输出：
// no such page: /help
func main() {
	db := database2{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}

type dollars2 float32

func (d dollars2) String() string { return fmt.Sprintf("$%.2f", d) }

type database2 map[string]dollars2

func (db database2) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		// w.WriteHeader(http.StatusNotFound)
		// fmt.Fprintf(w, "no such page: %s\n", req.URL)
		// 以上等同于
		msg := fmt.Sprintf("no such page: %s\n", req.URL)
		http.Error(w, msg, http.StatusNotFound)
	}
}
