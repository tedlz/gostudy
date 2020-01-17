package main

import (
	"flag"
	"fmt"
	"gostudy/07、接口/files"
)

// 005、flag.Value 接口 2
// go run 005_tempflag.go -temp 212F
// 输出：
// 100℃
func main() {
	flag.Parse()
	fmt.Println(*temp)
}

var temp = files.CelsiusFlag("temp", 20.0, "the temperature")
