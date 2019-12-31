package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

// 005、统计每个 unicode 码点的输入次数
// cat ../files/threekingdoms.txt | go run 005_charcount.go
// 输出：
// rune    count
// '不'    9
// '\n'    9
// '：'    5
// '，'    30
// '“'     6
// '操'    5
// '。'    13
// '义'    6
// '家'    5
// '”'     6

// len     count
// 1       13
// 2       0
// 3       320
// 4       0
func main() {
	counts := make(map[rune]int)    // unicode 字符计数
	var utflen [utf8.UTFMax + 1]int // utf8 长度计数（0-4）
	invalid := 0                    // 无效的 utf8 字符数

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // 返回解码的 rune 字符的值, 字符 utf8 编码后的长度, 和一个错误值
		if err == io.EOF {
			break // 如果读取到文件结尾，跳出
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 { // unicode.ReplacementChar 是无效字符且长度为 1
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\t\n")
	for c, n := range counts {
		if n >= 5 {
			fmt.Printf("%q\t%d\n", c, n) // 出现大于等于 5 次的才输出
		}
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
