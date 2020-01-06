package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// 003、递归 outline
// go run ../../01、入门/src/008_fetch.go https://golang.org | go run 003_outline.go
// 输出：
// [html]
// [html head]
// [html head meta]
// [html head meta]
// [html head meta]
// [html head meta]
// [html head title]
// [html head link]
// [html head link]
// [html head link]
// [html head link]
// [html head script]
// [html head script]
// [html head script]
// [html head script]
// [html head script]
// [html head script]
// [html body]
// [html body header]
// [html body header nav]
// [html body header nav a]
// [html body header nav a img]
// [html body header nav button]
// [html body header nav button div]
// [html body header nav ul]
// [html body header nav ul li]
// [html body header nav ul li a]
// [html body header nav ul li]
// [html body header nav ul li a]
// [html body header nav ul li]
// [html body header nav ul li a]
// [html body header nav ul li]
// [html body header nav ul li a]
// [html body header nav ul li]
// [html body header nav ul li a]
// [html body header nav ul li]
// [html body header nav ul li a]
// [html body header nav ul li]
// [html body header nav ul li form]
// [html body header nav ul li form input]
// [html body header nav ul li form button]
// [html body header nav ul li form button svg]
// [html body header nav ul li form button svg title]
// [html body header nav ul li form button svg path]
// [html body header nav ul li form button svg path]
// [html body main]
// [html body main div]
// [html body main div div]
// [html body main div div]
// [html body main div div section]
// [html body main div div section h1]
// [html body main div div section h1 strong]
// [html body main div div section h1 strong]
// [html body main div div section h1 strong]
// [html body main div div section i]
// [html body main div div section a]
// [html body main div div section a img]
// [html body main div div section p]
// [html body main div div section p br]
// [html body main div div section]
// [html body main div div section div]
// [html body main div div section div h2]
// [html body main div div section div a]
// [html body main div div section div]
// [html body main div div section div textarea]
// [html body main div div section div]
// [html body main div div section div pre]
// [html body main div div section div pre noscript]
// [html body main div div section div]
// [html body main div div section div select]
// [html body main div div section div select option]
// [html body main div div section div select option]
// [html body main div div section div select option]
// [html body main div div section div select option]
// [html body main div div section div select option]
// [html body main div div section div select option]
// [html body main div div section div select option]
// [html body main div div section div select option]
// [html body main div div section div div]
// [html body main div div section div div button]
// [html body main div div section div div div]
// [html body main div div section div div div button]
// [html body main div div section div div div a]
// [html body main div div section]
// [html body main div div section h2]
// [html body main div div section div]
// [html body main div div section div a]
// [html body main div div section]
// [html body main div div section h2]
// [html body main div div section div]
// [html body main div div section div iframe]
// [html body main div script]
// [html body footer]
// [html body footer div]
// [html body footer div img]
// [html body footer div ul]
// [html body footer div ul li]
// [html body footer div ul li a]
// [html body footer div ul li]
// [html body footer div ul li a]
// [html body footer div ul li]
// [html body footer div ul li a]
// [html body footer div ul li]
// [html body footer div ul li a]
// [html body footer div a]
// [html body script]
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}

// outline 有入站操作，但没有相对应的出栈操作
// 当 outline 调用自身时，被调用者接收的是 stack 的拷贝
// 被调用者对 stack 的元素追加操作，修改的是 stack 的拷贝，其可能会修改 slice 底层的数组甚至是申请一块新的内存空间进行扩容
// 但这个过程并不会修改调用方的 stack
// 因此当函数返回时，调用方的 stack 与其调用自身之前完全一致
func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

// 大部分编程语言使用固定大小的函数调用栈，常见的大小从 64KB 到 2MB 不等
// 固定大小栈会限制递归的深度，当你用递归处理大量数据时，需要避免栈溢出；除此之外，还会导致安全性问题
// Go 语言使用可变栈，栈的大小按需增加（初始很小），这使我们使用递归时不必考虑溢出和安全问题
