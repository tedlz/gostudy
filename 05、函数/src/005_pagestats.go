package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// 005、多返回值 pagestats
// go run 005_pagestats.go https://golang.org
// word count: 728
// image count: 3
func main() {
	words, images, err := CountWordsAndImages(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "stats: %v\n", err)
	}
	fmt.Println("word count:", words)
	fmt.Println("image count:", images)
}

// CountWordsAndImages *
// 如果一个函数所有的返回值都有变量名，那么该函数的 return 语句可以省略，这称之为 bare return
// 当一个函数中有多个 return 及许多返回值时，bare return 可以减少代码重复
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return // bare return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return // bare return
	}
	words, images = countWordsAndImages(doc)
	return // bare return，相当于 return words, images, err
}

func countWordsAndImages(n *html.Node) (words, images int) {
	texts, images := visit3(nil, 0, n)
	for _, v := range texts {
		v = strings.Trim(strings.TrimSpace(v), "\r\n")
		if v == "" {
			continue
		}
		words += strings.Count(v, "")
	}
	return // bare return
}

func visit3(texts []string, images int, n *html.Node) ([]string, int) {
	// 文本
	if n.Type == html.TextNode {
		texts = append(texts, n.Data)
	}
	// 图片
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Data == "script" || c.Data == "style" {
			continue
		}
		texts, images = visit3(texts, images, c)
	}
	return texts, images
}
