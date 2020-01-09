package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

// 017、延迟函数调用
// go run 017_fetch.go https://golang.org/
// 输出：
// 向 files 目录输出 index.html 文件
func main() {
	fmt.Println(fetch(os.Args[1]))
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	if !strings.HasSuffix(local, ".html") {
		local += ".html"
	}
	f, err := os.Create("../files/" + local)
	if err != nil {
		return "", 0, err
	}
	defer f.Close()

	n, err = io.Copy(f, resp.Body)
	if err != nil {
		return "", 0, err
	}
	return local, n, err
}
