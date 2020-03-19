package main

import (
	"fmt"
	"log"
	"os"

	links "gostudy/05、函数/files"
)

// 015、并发的 Web 爬虫 - crawl2 并行，限制并行数 20
// go run 015_crawl2.go http://gopl.io
func main() {
	worklist := make(chan []string)
	var n int // 待发送到 worklist 的数量

	n++
	// 接收命令行参数
	go func() { worklist <- os.Args[1:] }()

	// 同时抓取网页
	seen := make(map[string]bool)

	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl2(link)
				}(link)
			}
		}
	}
}

// tokens 是计数信号量，用来强制限制 20 个并发请求
var tokens = make(chan struct{}, 20)

func crawl2(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // 获取一个 token
	list, err := links.Extract(url)
	<-tokens // 释放这个 token
	if err != nil {
		log.Print(err)
	}
	return list
}
