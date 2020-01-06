package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// 004、多返回值 findlinks2
// go run 004_findlinks2.go https://golang.org
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
	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}

func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return visit2(nil, doc), nil
}

// 为了遍历节点 n 的所有后代节点，每次遇到 n 的子节点时，visit 递归调用自身，将子节点存放在 FirstChild 链表中
func visit2(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit2(links, c)
	}
	return links
}
