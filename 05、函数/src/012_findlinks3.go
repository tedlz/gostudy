package main

import (
	"fmt"
	"gostudy/05、函数/files"
	"log"
	"os"
)

// 012、匿名函数
// go run 012_findlinks3.go https://golang.org
// 输出：
// 页面上所有的链接
func main() {
	// 当所有发现的链接都已经被访问或内存耗尽时，程序运行结束
	breadthFirst(crawl, os.Args[1:])
}

// 广度优先算法
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...) // f(item)... 会将返回的一组元素一个个添加到 worklist 中
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := files.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
