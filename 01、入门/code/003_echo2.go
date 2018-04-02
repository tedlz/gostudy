package main

import (
	"fmt"
	"os"
)

// 003、输出命令行参数 2
// go run 003_echo2.go a b c
// 输出：
// a b c
func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
