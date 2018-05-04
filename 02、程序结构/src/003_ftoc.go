package main

import "fmt"

// 003、华氏度转摄氏度
// go run 003_ftoc.go
// 输出：
// 32°F = 0°C
// 212°F = 100°C
func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF))
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
