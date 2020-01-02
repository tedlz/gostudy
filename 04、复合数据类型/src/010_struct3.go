package main

import "fmt"

// 010、结构体嵌入和匿名成员
// go run 010_struct3.go
// 输出：
// main.Wheel{Circle:main.Circle{Point:main.Point{X:8, Y:8}, Radius:5}, Spokes:20}
// main.Wheel{Circle:main.Circle{Point:main.Point{X:2, Y:8}, Radius:5}, Spokes:20}
func main() {
	struct1()
	struct2()
	struct3()
}

func struct1() {
	// Circle 圆
	type Circle struct {
		X, Y, Radius int // 圆心的 X, Y 坐标，和 Radius 半径
	}

	// Wheel 轮
	type Wheel struct {
		X, Y, Radius, Spokes int // 除了 X, Y, Radius，还包含了 Spokes 径向辐条的数量
	}

	// 创建一个 Wheel 变量
	var w Wheel
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20
}

func struct2() {
	// 随着库中几何形状增多，为了便于维护可以将相同的属性独立出来
	type Point struct {
		X, Y int
	}
	type Circle struct {
		Center Point
		Radius int
	}
	type Wheel struct {
		Circle Circle
		Spokes int
	}
	// 这样改动之后结构变清晰了，但也使访问每个成员变的繁琐
	var w Wheel
	w.Circle.Center.X = 8
	w.Circle.Center.Y = 8
	w.Circle.Radius = 5
	w.Spokes = 20
}

func struct3() {
	// Go 语言有个特性让我们只声明一个成员对应的数据类型而不指明成员的名字，这类成员就叫匿名成员
	// 匿名成员的数据类型必须是命名的类型，或指向一个命名的类型的指针
	// 下面的代码中，Circle 和 Wheel 各自都有一个匿名成员
	// 我们可以说 Point 类型被嵌入到了 Circle 结构体，Circle 类型被嵌入到了 Wheel 结构体
	type Point struct {
		X, Y int
	}
	type Circle struct {
		Point
		Radius int
	}
	type Wheel struct {
		Circle
		Spokes int
	}
	// 得益于匿名嵌入的特性，我们可以直接访问叶子属性，而不需要给出完整的路径
	var w Wheel
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20
	// 但显式访问叶子成员的语法依然有效
	w.Circle.Point.X = 8
	w.Circle.Point.Y = 8
	w.Circle.Radius = 5
	w.Spokes = 20
	// 不幸的是，结构体字面值并没有简短表示匿名成员的语法，因此以下写法都会导致编译不通过
	// cannot use 8 (type int) as type Circle in field value
	// w = Wheel{8, 8, 5, 20}
	// cannot use promoted field Circle.Point.X in struct literal of type Wheel
	// w = Wheel{X: 8, Y: 8, Radius: 5, Spokes: 20}

	// 结构体字面值必须遵循形状类型声明时的结构
	// 所以我们只能用下面的两种语法，它们彼此是等价的
	w = Wheel{Circle{Point{8, 8}, 5}, 20}
	w = Wheel{
		Circle: Circle{
			Point: Point{
				X: 8,
				Y: 8,
			},
			Radius: 5,
		},
		Spokes: 20,
	}
	// main 为 package
	fmt.Printf("%#v\n", w) // main.Wheel{Circle:main.Circle{Point:main.Point{X:8, Y:8}, Radius:5}, Spokes:20}

	w.X = 2
	// main 为 package
	fmt.Printf("%#v\n", w) // main.Wheel{Circle:main.Circle{Point:main.Point{X:2, Y:8}, Radius:5}, Spokes:20}
}
