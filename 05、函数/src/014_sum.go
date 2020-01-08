package main

import (
	"fmt"
	"os"
	"reflect"
)

// 014、可变参数
// go run 014_sum.go
// 输出：
// 0
// 3
// 10
// 15
// func(...int)
// func([]int)
// Line 12: undefined: count
// 0
// 0
// 2
// -2
// ""
// a
// a, b, c
func main() {
	// 参数数量可变的函数称为可变参函数
	// 典型的例子就是 fmt.Printf 和类似函数，Printf 首先接收一个必备的参数，之后接收任意个数的后续参数
	fmt.Println(sum())           // 0
	fmt.Println(sum(3))          // 3
	fmt.Println(sum(1, 2, 3, 4)) // 10
	// 在上面的代码中，调用者隐式的创建一个数组，并将原始参数复制到数组中，再把数组的一个切片作为参数传给被调函数
	// 如果原始参数已经是切片类型，需要在最后一个参数后加上省略符
	values := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(values...)) // 15

	// 虽然在可变参函数内部，...int 型参数的行为看起来很像切片类型，但实际上，可变参函数和以切片作为参数的函数是不同的
	fmt.Printf("%T\n", f)          // func(...int)
	fmt.Println(reflect.TypeOf(g)) // func([]int)

	linenum, name := 12, "count"
	errorf(linenum, "undefined: %s", name) // Line 12: undefined: count

	// 练习
	values = []int{-2, -1, 0, 1, 2}
	fmt.Println(max())          // 0
	fmt.Println(min())          // 0
	fmt.Println(max(values...)) // 2
	fmt.Println(min(values...)) // -2

	str := []string{"a", "b", "c"}
	fmt.Printf("%q\n", stringsJoin(", ", []string{}...)) // ""
	fmt.Println(stringsJoin(", ", "a"))                  // a
	fmt.Println(stringsJoin(", ", str...))               // a, b, c
}

// sum 函数返回任意个 int 型参数的和
// 在函数体中，vals 被看做是 []int 类型的切片
// sum 可以接收任意数量的 int 型参数
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func f(...int) {}
func g([]int)  {}

// interface{} 表示函数的最后一个参数可以接收任意类型
func errorf(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}

func max(vals ...int) int {
	result := 0
	for i, val := range vals {
		if i == 0 || result < val {
			result = val
		}
	}
	return result
}

func min(vals ...int) int {
	result := 0
	for i, val := range vals {
		if i == 0 || result > val {
			result = val
		}
	}
	return result
}

func stringsJoin(symbol string, vals ...string) string {
	var result string
	for i, val := range vals {
		if i == 0 {
			result = val
		} else {
			result += symbol + val
		}
	}
	return result
}
