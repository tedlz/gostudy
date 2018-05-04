package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

// 006、输出命令行参数 4
// go run 006_echo4.go -h
// 输出：
// Usage of /tmp/go-build771787477/b001/exe/006_echo4:
//   -n    omit trailing newline
//   -s string
//         separator (default " ")
// exit status 2
func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

// 其它示例：
// go run 006_echo4.go a b c
// 输出：
// a b c
// go run 006_echo4.go -s / a b c
// 输出：
// a/b/c
// go run 006_echo4.go -n a b c
// 输出：
// a b c%
