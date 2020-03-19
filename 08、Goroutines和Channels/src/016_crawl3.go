package main

import (
	"fmt"
	links "gostudy/05、函数/files"
	"log"
	"os"
)

// 016、并发的 Web 爬虫 - crawl3 并行，限制并行数 20 的另一种思路
// go run 016_crawl3.go http://gopl.io
func main() {
	worklist := make(chan []string)  // URL 列表，可能有重复
	unseenLinks := make(chan string) // 重复的 URL

	// 将命令行参数加入到 worklist
	go func() { worklist <- os.Args[1:] }()

	// 创建 20 个 goroutines 爬虫来获取每个看不到的链接
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl3(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// 主 goroutine 对 worklist 的重复数据进行删除，并将看不到的链接发给爬虫
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}

func crawl3(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
