package main

import (
	"fmt"
	"os"
)

// 002、输出命令行参数 1
// go run 002_echo1.go a b c
// 输出：
// a b c
func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
