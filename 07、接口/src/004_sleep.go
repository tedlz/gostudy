package main

import (
	"flag"
	"fmt"
	"time"
)

// 004、flag.Value 接口
// go run 004_sleep.go -period 1s
// 输出：
// Sleeping for 1s
func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period) // Sleeping for 1s
	time.Sleep(*period)
	fmt.Println()
}

var period = flag.Duration("period", 1*time.Second, "sleep period")
