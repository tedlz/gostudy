package main

import (
	"bufio"
	"fmt"
	"os"
)

// 005、查找重复的行 1
// cat ../files/005_dup1.txt | go run 005_dup1.go
// 如果用命令行方式调用，需要按 Ctrl+D 终止输入才有输出
// 输出：
// 3        a
// 2        c
func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
