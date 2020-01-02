package main

import "fmt"

// 009、结构体字面值与结构体比较
// go run 009_struct2.go
// 输出：
// {1 2} {11 0}
// 2 4
// 4 2
// 4 4
// true
// false
// false true
func main() {
	// 结构体值也可以用结构体字面值表示，结构体字面值可以指定每个成员的值
	p := Point{1, 2} // 第一种写法
	q := Point{      // 第二种写法，更常用，避免写法一和二混用
		X: 11,
		// 忽略了 Y，被忽略的将默认用该值类型的零值
	}
	fmt.Println(p, q) // {1 2} {11 0}

	p = Scale(p, 2)
	fmt.Println(p.X, p.Y) // 2 4

	x := Bonus(&p, 200)
	fmt.Println(x, p.X) // 4 2

	AwardAnnualRaise(&p)
	fmt.Println(p.X, p.Y) // 4 4

	// 结构体通常通过指针处理，可以用下面的写法来创建并初始化一个结构体变量，并返回结构体的地址
	// pp := &Point{1, 2}
	// 它和下面的语句是等价的
	// qq := new(Point)
	// *qq = Point{1, 2}
	// 不过 &Point{1, 2} 写法可以直接在表达式中使用，比如一个函数调用

	// 如果结构体全部成员都是可以比较的，那么结构体也是可以比较的；这样的话两个结构体可以使用 == 或 != 运算符进行比较
	a := Point{1, 2}
	b := Point{2, 1}
	fmt.Println(a.X == b.Y && b.X == a.Y) // true
	fmt.Println(a.X == b.X && a.Y == b.Y) // false
	fmt.Println(a == b, a != b)           // false true

	// 可比较的结构体类型和其它可比较的类型一样，可以用于 map 的 key 类型
	type address struct {
		hostname string
		port     int
	}
	hits := make(map[address]int)
	hits[address{"golang.org", 443}]++
}

// Point *
type Point struct{ X, Y int }

// Scale - 结构体可以作为函数的参数和返回值
func Scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

// Bonus - 较大的结构体通常会以指针的方式传入和返回
func Bonus(p *Point, percent int) int {
	return p.X * percent / 100
}

// AwardAnnualRaise - 如果要在函数内部修改结构体成员，指针传入是必须的；因为在 Go 语言中，所有函数参数都是值拷贝
func AwardAnnualRaise(p *Point) {
	p.X = p.X * 200 / 100
}
