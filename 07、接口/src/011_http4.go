package main

import (
	"fmt"
	"log"
	"net/http"
)

// 011、http.Handler 接口 4
func main() {
	// 为了方便，net/http 包提供了一个全局的 ServeMux 实例 DefaultServeMux
	// 和包级别的 http.Handle 和 http.HandleFunc 函数
	// 现在，为了使用 DefaultServeMux 作为服务器的主 handler，我们不需要将它传给 ListenAndServe 函数，nil 值就可以工作
	// 所以 http3.go 的 main 函数可以简写为：
	db := database3{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars3 float32
type database3 map[string]dollars3

func (db database3) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database3) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (d dollars3) String() string { return fmt.Sprintf("$%.2f", d) }
