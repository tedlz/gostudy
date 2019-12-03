package main

import "fmt"

// 007、new 函数
// go run 007_new.go
// 输出：
// 0
// 2
// 0
// 3
// 0
// 4
// false
// 1
func main() {
	// 另一个创建类型变量的方法是调用内建的 new 函数
	// new(T) 将创建一个 T 类型的匿名变量，初始化为 T 类型的零值
	// 然后返回变量地址，返回的指针类型为 *T

	// 一
	p := new(int)   // p, *int 类型，指向匿名的 int 变量
	fmt.Println(*p) // 0
	*p = 2          // 设置 int 匿名变量的值为 2
	fmt.Println(*p) // 2

	// 二
	a := newInt1()
	fmt.Println(*a) // 0
	*a = 3
	fmt.Println(*a) // 3

	// 三
	b := newInt2()
	fmt.Println(*b) // 0
	*b = 4
	fmt.Println(*b) // 4

	// 四、每次调用 new 函数都返回新的变量地址，因此下面两个地址是不同的
	x := new(int)
	y := new(int)
	fmt.Println(x == y) // false

	// 五
	fmt.Println(delta(*b, *a))
}

// 下面两个函数拥有相同的行为，例子见上面二、三
func newInt1() *int {
	return new(int)
}

func newInt2() *int {
	var dummy int
	return &dummy
}

// new 只是一个预定义函数，非关键字，因此可以将 new 名字重新定义为别的类型
// 由于 new 被定义为 int 变量名，所以在函数 delta 中无法使用内置的 new 函数
func delta(old, new int) int {
	return new - old
}
