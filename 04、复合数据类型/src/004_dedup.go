package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 004、MAP 例子（接收输入，并且输入相同的内容，只输出第一行）
// go run 004_dedup.go
// 输出：
// 相同输入内容的第一次输入
func main() {
	// Go 语言没有 set 类型，可以用 map 实现类似 set 的功能
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin) // 获取输入流
	for input.Scan() {                  // 循环接收输入
		line := input.Text() // 输入内容
		if !seen[line] {     // 如果相同内容之前未输入过
			seen[line] = true       // 标记为输入过
			fmt.Println("输出", line) // 并且输出
		} // 否则就接收下一次输入

		slice := strings.Split(line, "") // 把字符串（string）转为切片（slice）
		add(slice)
		fmt.Printf("输入过 %s [%d] 次\n", line, count(slice))
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}

// 有时候我们需要一个 map 的 key 是 slice 类型，但 map 的 key 必须是可比较类型，slice 并不满足这个条件
// 不过我们可以通过两个步骤绕过这个限制：
// 1、定义一个辅助函数 k，将 slice 转为 map 对应的 string 类型的 key，确保 x 和 y 相等时 k(x) == k(y) 才成立
// 2、然后创建一个 key 为 string 类型的 map，在每次对 map 操作时先用 k 函数将 slice 转换为 string 类型
var m = make(map[string]int)

func k(list []string) string  { return fmt.Sprintf("%q", list) }
func add(list []string)       { m[k(list)]++ }
func count(list []string) int { return m[k(list)] }

// 使用同样的方式可以处理任何不可比较的 key 类型，而不仅仅是 slice
// 这种方式对于想使用自定义 key 比较函数的时候也很有用，例如在比较字符串的时候忽略大小写
// 同时，辅助函数 k(x) 也不一定是字符串类型，它可以返回任何可比较的类型，例如整数、数组或结构体等
