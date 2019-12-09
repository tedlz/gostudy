package main

import (
	"fmt"
	"os"
	"strconv"

	"gostudy/02、程序结构/files"
)

// 011、包和文件02 - 导入包
// go run 011_package02.go 32
// 输出：
// 32℉ = 0℃, 32℃ = 89.6℉
func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := files.Fahrenheit(t)
		c := files.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, files.FToC(f), c, files.CToF(c))
	}
}
