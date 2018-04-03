package main

import (
	"fmt"
	"log"
	"net/http"
)

// 012、Web 服务 3，打印请求头和参数
// go run 012_server3.go
// 浏览器访问 localhost:8000/?query=abc&other=123
// 输出：
// GET /?query=abc&other=123 HTTP/1.1
// Header["Accept"] = ["text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8"]
// Header["Dnt"] = ["1"]
// Header["Accept-Encoding"] = ["gzip, deflate, br"]
// Header["Accept-Language"] = ["zh-CN,zh;q=0.9"]
// Header["Connection"] = ["keep-alive"]
// Header["Upgrade-Insecure-Requests"] = ["1"]
// Header["User-Agent"] = ["Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36"]
// Host = "localhost:8000"
// RemoteAddr = "127.0.0.1:7610"
// Form["query"] = ["abc"]
// Form["other"] = ["123"]
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
