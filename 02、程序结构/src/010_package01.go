package main

import (
	"fmt"
	"gostudy/02、程序结构/files"
)

// 010、包和文件01
// go run 010_package01.go
// 输出：
// Brrrr! -273.15℃
// 212℉
func main() {
	fmt.Printf("Brrrr! %v\n", files.AbsoluteZeroC) // Brrrr! -273.15℃
	fmt.Println(files.CToF(files.BoilingC))        // 212℉
}
