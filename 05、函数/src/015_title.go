package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// 015、延迟函数调用
// go run 015_title.go
// 输出：
// title1: Effective Go - The Go Programming Language
// title2: Effective Go - The Go Programming Language
func main() {
	url := os.Args[1]
	// resp.Body.Close 调用了多次，这是确保 title 在所有执行路径下（即使函数运行失败）都关闭了网络连接
	title1(url)
	// 随着函数变的复杂，需要处理的错误也变多，维护清理逻辑变的越来越困难，而 Go 语言独有的 defer 机制可以让事情变简单
	title2(url)
}

func title1(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		resp.Body.Close()
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	flag := false // 标记，只输出第一个 title
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode &&
			n.Data == "title" &&
			n.FirstChild != nil &&
			!flag {
			flag = true
			fmt.Println("title1:", n.FirstChild.Data)
		}
	}
	forEachNode(doc, visitNode, nil)
	return nil
}

func title2(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		return fmt.Errorf("%s has type %s", url, ct)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	flag := false
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode &&
			n.Data == "title" &&
			n.FirstChild != nil &&
			!flag {
			flag = true
			fmt.Println("title2:", n.FirstChild.Data)
		}
	}
	forEachNode(doc, visitNode, nil)
	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
