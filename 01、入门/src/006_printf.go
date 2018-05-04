package main

import "fmt"

// 006、Printf 的转换类型
// go run 006_printf.go
// 输出：
// %d      十进制整数
// %x, %o, %b  十六进制、八进制、二进制整数
// %f, %g, %e  浮点数：3.141593、3.141592653589793、3.141593e+00
// %t      布尔值：true、false
// %c      字符（rune）（Unicode 码点）
// %s      字符串
// %q      带双引号的字符串 "abc" 或带单引号的字符 'c'
// %v      变量的自然形式（natural format）
// %T      变量的类型
// %%      字面上的 % 百分号标志（无操作数）
func main() {
	transfer := []string{
		"%d\t\t十进制整数",
		"%x, %o, %b\t十六进制、八进制、二进制整数",
		"%f, %g, %e\t浮点数：3.141593、3.141592653589793、3.141593e+00",
		"%t\t\t布尔值：true、false",
		"%c\t\t字符（rune）（Unicode 码点）",
		"%s\t\t字符串",
		"%q\t\t带双引号的字符串 \"abc\" 或带单引号的字符 'c'",
		"%v\t\t变量的自然形式（natural format）",
		"%T\t\t变量的类型",
		"%%\t\t字面上的 % 百分号标志（无操作数）",
	}

	for _, arg := range transfer {
		fmt.Println(arg)
	}
}
