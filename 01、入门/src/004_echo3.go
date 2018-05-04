package main

import (
	"fmt"
	"os"
	"strings"
)

// 003、输出命令行参数 3
// go run 004_echo3.go a b c
// 输出：
// a b c
func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
