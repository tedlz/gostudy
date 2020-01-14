package main

import (
	"fmt"
	"image/color"
	"math"
)

// 005、通过嵌入结构体来扩展类型
// go run 005_coloredpoint.go
// 输出：
// 1
// 2
// 5
// 10
// 5
// {2 2} {2 2}
func main() {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point3.X) // 1
	cp.Point3.Y = 2
	fmt.Println(cp.Y) // 2

	// Point3 类的方法也被引入了 ColoredPoint
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Point3{1, 1}, red}
	var q = ColoredPoint{Point3{5, 4}, blue}
	fmt.Println(p.Distance(q.Point3)) // 5
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point3)) // 10

	// Distance 有一个参数是 Point3 类型，但 q 并不是一个 Point3 类
	// 所以尽管 q 有着 Point3 这个内嵌类型，我们也必须要显式的选择它。直接传 q 你会看到下面的错误：
	// p.Distance(q)
	// 编译错误：cannot use q (type ColoredPoint) as type Point3 in argument to p.Point3.Distance

	// 一个 ColoredPoint 并不是一个 Point3，但它 “has a” Point3
	// 并且它有从 Point3 类里引入的 Distance 和 ScaleBy 方法
	// 如果你喜欢从实现的角度来考虑问题，内嵌字段会指导编译器去生成额外的包装方法来委托已经声明好的方法，和下面的形式等价：
	// func (p ColoredPoint) Distance(q Point3) float64 {
	// 	   return p.Point3.Distance(q)
	// }
	// func (p *ColoredPoint) ScaleBy(factor float64) {
	// 	   p.Point3.ScaleBy(factor)
	// }
	// 当 Point3.Distance 被第一个包装方法调用时，它的接收器值是 p.Point3 而不是 p
	// 当然在 Point3 类的方法里，你是访问不到 ColoredPoint 里的任何字段的

	// 在类型中内嵌的匿名字段也可能是一个命名类型的指针，这种情况下字段和方法会被间接的引入到当前的类型中
	// 添加这一层间接关系让我们可以共享通用的结构并动态的去改变对象之间的关系
	// 下面这个 ColoredPoint2 的声明内嵌了一个 Point3 类型的指针

	x := ColoredPoint2{&Point3{1, 1}, red}
	y := ColoredPoint2{&Point3{5, 4}, blue}
	fmt.Println(x.Distance(*y.Point3)) // 5
	y.Point3 = x.Point3
	x.ScaleBy(2)
	fmt.Println(*x.Point3, *y.Point3) // {2 2} {2 2}

	// 一个 struct 类型也可能会有多个匿名字段，我们将 ColoredPoint3 定义为下面这样：
	type ColoredPoint3 struct {
		Point3
		color.RGBA
	}
	// 然后这种类型的值便会拥有 Point3 和 color.RGBA 类型的所有方法，以及直接定义在 ColoredPoint3 中的方法
	// 当编译器解析一个选择器到方法时，比如 p.ScaleBy，它会首先去找直接定义在这个类型里的 ScaleBy 方法
	// 然后找被 ColoredPoint3 内嵌字段们引入的方法，然后去找 Point3 和 RGBA 的内嵌字段引入的方法
	// 然后一直递归向下找。如果选择器有二义性的话编译器会报错，比如你在同一级里有两个同名的方法

	// 方法只能在命名类型（像 Point3）或者指向类型的指针上定义
	// 但是多亏了内嵌，有些时候我们给匿名 struct 类型来定义方法也有了手段
}

// Point3 *
type Point3 struct{ X, Y float64 }

// ColoredPoint *
type ColoredPoint struct {
	Point3
	Color color.RGBA
}

// Distance *
func (p Point3) Distance(q Point3) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// ScaleBy *
func (p *Point3) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

// ColoredPoint2 *
type ColoredPoint2 struct {
	*Point3
	Color color.RGBA
}
