package main

import "fmt"

const boilingF = 212.0

// 002、声明
// go run 002_boiling.go
// 输出：
// boiling point: = 212°F or 100°C
func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point: = %g°F or %g°C\n", f, c)
}
