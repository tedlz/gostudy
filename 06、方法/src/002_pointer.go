package main

import "fmt"

// 002、基于指针对象的方法
// go run 002_pointer.go
// 输出：
// &{2 4}
// {2 4}
// &{2 4}
// {2 4}
// {2 4}
func main() {
	// 想要调用指针类型方法 (*Point2).ScaleBy，只要提供一个 Point2 类型的指针即可，像下面这样：
	r := &Point2{1, 2}
	r.ScaleBy(2)
	fmt.Println(r) // &{2 4}
	// 或者这样：
	p := Point2{1, 2}
	pptr := &p
	pptr.ScaleBy(2)
	fmt.Println(p)    // {2 4}
	fmt.Println(pptr) // &{2 4}
	// 或者这样：
	q := Point2{1, 2}
	(&q).ScaleBy(2)
	fmt.Println(q) // {2 4}

	// 不过后面两种方法有些笨拙，幸运的是，Go 语言本身在这种地方会帮到我们
	// 如果接收器 x 是一个 Point2 类型的变量，并且其方法需要一个 Point2 指针作为接收器，我们可以用下面这种简短的写法：
	x := Point2{1, 2}
	x.ScaleBy(2)
	fmt.Println(x) // {2 4}
	// 编译器会隐式的帮我们去用 &x 调用 ScaleBy 这个方法
	// 这种简写方法只适用于变量，包括 struct 里的字段比如 p.X，以及 array 和 slice 里的元素比如 perim[0]
	// 我们不能通过一个无法取到地址的接收器来调用指针方法，比如临时变量的内存地址就无法获取的到：
	// Point2{1, 2}.ScaleBy(2) // 编译错误：cannot take the address of composite literal

	// 但是我们可以用一个 *Point2 这样的接收器来调用 Point2 方法，因为我们可以通过地址来找到这个变量
	// 只要用解引用符号（*）来取到该变量即可，编译器在这里也会隐式的帮我们插入 * 这个操作符
	// 因此以下写法是等价的：
	// pptr.Distance(q)
	// (*pptr).Distance(q)
}

// 当调用一个函数时，会对其每一个参数值进行拷贝
// 如果一个函数需要更新一个变量，或者函数的其中一个参数实在太大，我们希望避免这样默认的拷贝
// 这种情况下我们就需要用到指针了

// Point2 *
type Point2 struct{ X, Y float64 }

// ScaleBy *
// 这个方法的名字是 (*Point2).ScaleBy
// 这里的括号是必须的，没有括号的话这个表达式可能会被理解为 *(Point2.ScaleBy)
func (p *Point2) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

// 在现实的程序里，一般会约定如果 Point2 这个类有一个指针作为接收器的方法，
// 那么所有 Point2 的方法都必须有一个指针接收器，即使是那些并不需要这个指针接收器的函数
// 只有类型（Point2）和指向它们的指针（*Point2），才是可能会出现在接收器声明里的两种接收器

// 此外，为了避免歧义，在声明方法时，如果一个类型名本身是一个指针的话，是不允许其出现在接收器中的
// type P *int
// func (P) f() {} // 编译错误：invalid receiver type P (P is a pointer type)
