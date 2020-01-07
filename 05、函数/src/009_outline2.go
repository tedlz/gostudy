package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// 009、函数值 outline2
// go run 009_outline2.go https://golang.org
// 输出：
// HTML 源码
func main() {
	url := os.Args[1]
	doc, err := findLinks2(url)
	if err != nil {
		log.Fatalf("get doc error: %v\n", err)
	}
	// 使用函数值，我们可以将遍历节点的逻辑和操作节点的逻辑分离
	forEachNode(doc, startElement, endElement)
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

var depth int

func startElement(n *html.Node) {
	if n.Type == html.TextNode {
		text := strings.Trim(strings.TrimSpace(n.Data), "\r\n")
		if text != "" {
			fmt.Printf("%*s%s\n", multi(), "", text)
		}
	} else if n.Type == html.CommentNode {
		fmt.Printf("%*s<!-- %s -->\n", multi(), "", n.Data)
	} else if n.Type == html.ElementNode {
		attrs := []string{}
		for _, c := range n.Attr {
			attrStr := ""
			if c.Val == "" {
				attrStr = fmt.Sprintf("%s", c.Key)
			} else {
				attrStr = fmt.Sprintf("%s=\"%s\"", c.Key, c.Val)
			}
			attrs = append(attrs, attrStr)
		}
		format := "%*s<%s%s>\n"
		if noChildElement(n) {
			format = "%*s<%s%s />\n"
		}
		add := ""
		if len(attrs) > 0 {
			add = " " + strings.Join(attrs, " ")
		}
		fmt.Printf(format, multi(), "", n.Data, add) // %*s 中的 * 会在字符串之前填充一些空格
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if noChildElement(n) {
			fmt.Printf("")
		} else {
			fmt.Printf("%*s</%s>\n", multi(), "", n.Data)
		}
	}
}

func findLinks2(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("get %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parse %s as HTML: %v", url, err)
	}
	return doc, nil
}

func noChildElement(n *html.Node) bool {
	return n.Data == "meta" ||
		n.Data == "link" ||
		n.Data == "img"
}

func multi() int {
	return depth * 2
}
