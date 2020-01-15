package main

import (
	"fmt"
	"math"
	"time"
)

// 007、方法值和方法表达式
// go run 007_method.go
// 输出：
// 5
// 2.23606797749979
// {2 4}
// {6 12}
// {60 120}
// 5
// func(main.Point4, main.Point4) float64
// {2 4}
// func(*main.Point4, float64)
// [{5 9} {7 11}]
func main() {
	// 我们经常选择一个方法，并且在同一个表达式里执行，比如常见的 p.Distance 形式，实际上将其分成两步来执行也是可能的
	// p.Distance 叫做选择器，选择器会返回一个方法值，一个将方法（Point.Distance）绑定到特定接收器变量的函数
	// 这个函数可以不通过指定其接收器即可被调用，即调用时不需要指定接收器，只要传入函数的参数即可
	p := Point4{1, 2}
	q := Point4{4, 6}
	distanceFromP := p.Distance        // method value 方法值
	fmt.Println(distanceFromP(q))      // 5
	var origin Point4                  // {0, 0}
	fmt.Println(distanceFromP(origin)) // 2.23606797749979, sqrt(5)

	scaleP := p.ScaleBy // method value 方法值
	scaleP(2)
	fmt.Println(p) // {2 4}
	scaleP(3)
	fmt.Println(p) // {6 12}
	scaleP(10)
	fmt.Println(p) // {60 120}

	// 在一个包的 API 需要一个函数值、且调用方希望操作的是某一个绑定了对象的方法的话，方法值会非常实用
	// 举例来说，下面例子中的 time.AfterFunc 这个函数的功能是在指定的延迟时间之后来执行一个另外的函数
	// 且这个函数操作的是一个 Rocket 对象 r
	r := new(Rocket)
	time.AfterFunc(3*time.Second, func() { r.Launch() })
	// 直接使用方法值传入 AfterFunc 的话可以更为简短
	time.AfterFunc(3*time.Second, r.Launch)

	// 和方法值相关的还有方法表达式
	// 当调用一个方法时，与调用一个普通的函数相比，我们必须要用选择器（p.Distance）语法来指定方法的接收器
	// 当 T 是一个类型时，方法表达式可能会写作 T.f 或 (*T).f，会返回一个函数值
	// 这种函数会将其第一个参数用作接收器，所以可以用通常（不写选择器）的方式来对其进行调用
	x := Point4{1, 2}
	y := Point4{4, 6}
	distance := Point4.Distance  // method expression 方法表达式
	fmt.Println(distance(x, y))  // 5
	fmt.Printf("%T\n", distance) // func(main.Point4, main.Point4) float64

	scale := (*Point4).ScaleBy
	scale(&x, 2)
	fmt.Println(x)            // {2 4}
	fmt.Printf("%T\n", scale) // func(*main.Point4, float64)

	z := Path2{x, y}
	z.TranslateBy(Point4{3, 5}, true)
	fmt.Println(z) // [{5 9} {7 11}]
}

// Point4 *
type Point4 struct{ X, Y float64 }

// Distance *
func (p Point4) Distance(q Point4) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y) // 相当于 math.Sqrt(a*a + b*b)
}

// ScaleBy *
func (p *Point4) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

// Rocket *
type Rocket struct{}

// Launch *
func (r *Rocket) Launch() {}

// 当你根据一个变量来决定调用同一个类型的哪个函数时，方法表达式就显得很有用了，你可以根据选择来调用接收器各不相同的方法
// 下面的例子，变量 op 代表 Point4 类型的 addition 或者 subtraction 方法，
// Path2.TranslateBy 方法会为其 Path2 数组中的每一个 Point4 来调用对应的方法

// Add *
func (p Point4) Add(q Point4) Point4 { return Point4{p.X + q.X, p.Y + q.Y} }

// Sub *
func (p Point4) Sub(q Point4) Point4 { return Point4{p.X - q.X, p.Y - q.Y} }

// Path2 *
type Path2 []Point4

// TranslateBy *
func (path Path2) TranslateBy(offset Point4, add bool) {
	var op func(p, q Point4) Point4
	if add {
		op = Point4.Add
	} else {
		op = Point4.Sub
	}
	for i := range path {
		// 调用 path[i].Add(offset) 或 path[i].Sub(offset)
		path[i] = op(path[i], offset)
	}
}
