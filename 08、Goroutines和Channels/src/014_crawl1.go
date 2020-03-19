package main

import (
	"fmt"
	"log"
	"os"

	links "gostudy/05、函数/files"
)

// 014、并发的 Web 爬虫 - crawl1 并行，但未限制并行数
// go run 014_crawl1.go http://gopl.io
func main() {
	worklist := make(chan []string)

	// 接收命令行参数
	go func() { worklist <- os.Args[1:] }()

	// 同时抓取网页
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
