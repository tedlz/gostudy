package main

import "fmt"

// 005、指针
// go run 005_pointer.go
// 输出：
// 1
// 2

// true false false
// 0xc420012130 0xc420012138 false

// 3
func main() {
	x := 1
	p := &x         // p, of type *int, points to x
	fmt.Println(*p) // 1
	*p = 2          // equivalent to x = 2
	fmt.Println(x)  // 2
	fmt.Println()

	// 任何指针类型的零值都是 nil，指针之间可以进行相等测试，当它们指向同一个变量或全部是 nil 时才相等
	var a, b int
	fmt.Println(&a == &a, &a == &b, &a == nil) // true, false, false

	// 返回函数中局部变量的地址也是安全的
	// 调用 f() 函数时创建局部变量 v，在局部变量地址被返回后依然有效，因为指针 z 依然引用这个变量
	fmt.Println(f(), f(), f() == f()) // false
	fmt.Println()

	v := 1
	incr(&v)              // 2
	fmt.Println(incr(&v)) // 3
}

var z = f()

func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++
	return *p // 引用相当于对原变量创建了新别名，此处 p 为 v 的别名
}
