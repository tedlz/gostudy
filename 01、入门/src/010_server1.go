package main

import (
	"fmt"
	"log"
	"net/http"
)

// 010、Web 服务 1
// go run 010_server1.go
// 浏览器访问 localhost:8000，输出：
// URL.Path
// 1) 可使用 ../build/010_server1& 把进程放后台执行
// 2) 通过 jobs 命令查看后台进程编号 {num}
// 3) 使用 fg %{num} 把进程放前台执行
// 4) 输入 Ctrl+Z 暂停进程
// 5) 再输入 bg %{num} 把进程放后台执行
// 6) 可以通过 kill %{num} 杀掉后台进程
// 7) 也可以通过 jobs -l 命令看到进程的 pid，之后用 kill {pid} 杀死进程
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
