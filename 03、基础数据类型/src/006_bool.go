package main

import "fmt"

// 006、布尔型
// go run 006_bool.go
// 输出：
// false
// false
// 1
// 1
// false
func main() {
	var s string = "abc"
	var c int8 = 127
	// bool 类型的值仅两种，true 或 false
	// 布尔值可以和 && 与 || 操作符结合，并且有短路行为：
	// 如果运算符左边的值已经可以确定整个布尔表达式的值，那么运算符右边的值将不再被求值
	// 因此下面的表达式总是安全的
	fmt.Println(s != "" && s[0] == 'x') // false, 其中 s[0] 操作如果应用于空字符串时会导致 panic 异常

	// 因为 && 的优先级比 || 高，下面的布尔表达式是不需要加小括号的
	if 'a' <= c && c <= 'z' ||
		'A' <= c && c <= 'Z' ||
		'0' <= c && c <= '9' {
		fmt.Println(true)
	} else {
		fmt.Println(false) // false
	}

	// 布尔值并不会隐式转换为数字 0 和 1，因此需要显式转换
	// 如果经常使用，建议封装（btoi、itob）
	i := 0
	if 1 == 1 {
		i = 1
	}
	fmt.Println(i) // 1

	fmt.Println(btoi(1 == 1)) // 1
	fmt.Println(itob(0))      // false
}

// bool 转 int
func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

// int 转 bool
func itob(i int) bool {
	return i != 0
}
