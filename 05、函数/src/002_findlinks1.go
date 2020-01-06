package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// 002、递归 findlinks1
// go run ../../01、入门/src/008_fetch.go https://golang.org | go run 002_findlinks1.go
// 输出：
// /
// /doc/
// /pkg/
// /project/
// /help/
// /blog/
// https://play.golang.org/
// /dl/
// https://tour.golang.org/
// https://blog.golang.org/
// /doc/copyright.html
// /doc/tos.html
// http://www.google.com/intl/en/policies/privacy/
// http://golang.org/issues/new?title=x/website:
// https://google.com
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// 为了遍历节点 n 的所有后代节点，每次遇到 n 的子节点时，visit 递归调用自身，将子节点存放在 FirstChild 链表中
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
