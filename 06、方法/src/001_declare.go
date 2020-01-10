package main

import (
	"fmt"
	"gostudy/06、方法/files"
	"math"
)

// 001、方法声明
// go run 001_declare.go
// 输出：
// 5
// 5
// 12
// 12
// 12
func main() {
	// 在函数声明时，在其名字之前放上一个变量，即是一个方法
	// 这个附加的参数会将该函数附加到这种类型上，即相当于为这种类型定义了一个独占的方法

	// 在 Go 语言中，我们并不会像其它语言那样用 this 或者 self 作为接收器，我们可以任意选择接收器的名字
	// 由于接收器的名字经常会被使用到，所以保持其在方法间传递时的一致性和简短性是不错的主意
	// 建议使用其类型的第一个字母，例如 Distance 方法 Point 类型的 p

	// 在方法调用过程中，接收器参数一般会在方法名之前出现
	// 这和方法声明是一样的，都是接收器参数在方法名字之前，下面是例子：
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q)) // 5，函数调用，function call
	fmt.Println(p.Distance(q))  // 5，方法调用，method call

	// 可以看到，上面两个调用都是 Distance，但却没有冲突
	// 第一个 Distance 调用实际上用的是包级别的函数 main.Distance
	// 第二个 Distance 则是使用刚声明的 Point，调的是 Point.Distance

	// 这种 p.Distance 的表达式叫做选择器（Selector），因为它会选择合适的对应这个对象的 Distance 方法来执行
	// 选择器也会被用来选择一个 struct 类型的字段，比如 p.X
	// 由于方法和字段都是在同一命名空间，所以如果我们在这里声明一个 X 方法的话，编译器会报错，因为在调用 p.X 时会有歧义

	// 每种类型都有其方法的命名空间，我们在用 Distance 这个名字的时候，不同的 Distance 调用指向了不同类型里的 Distance 方法

	// Path 是一个命名的 slice 类型，而不是 Point 那样的 struct 类型，然而我们依然可以为它定义方法
	// 在能够给任意类型定义方法这一点上，Go 和很多其它的面向对象的语言不太一样
	// 因此在 Go 语言里，我们为一些简单的数值、字符串、slice、map 来定义一些附加行为很方便
	// 我们可以给同一个包内的任意命名的类型定义方法，只要这个命名类型的底层类型不是指针或者 interface
	// （这个例子里，底层类型就是指 []Point 这个 slice，Path 就是命名类型）

	// 两个 Distance 方法有不同的类型，它们两个方法之间没有任何关系，
	// 尽管 Path 的 Distance 方法会在内部调用 Point.Distance 方法来计算每个连接邻接点的线段的长度
	// 我们来调用一个新方法，计算三角形的周长：
	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance()) // 12

	// 在上面两个对 Distance 名字的方法的调用中，编译器会根据方法的名字以及接收器来决定具体调用的是哪一个函数
	// 第一个例子中，path[i-1] 数组中的类型是 Point，因此 Point.Distance 方法被调用
	// 第二个例子中，perim 的类型是 Path，因此 Distance 调用的是 Path.Distance

	// 对于一个给定的类型，其内部方法都必须有唯一的方法名，但是不同类型却可以有同样的方法名
	// 比如我们这里 Path 和 Point 都有 Distance 方法，所以我们没必要在方法名之前加类型名来消除歧义，例如 PathDistance
	// 这里我们已经看到了方法比函数的一些好处，方法名可以简短
	// 当我们在包外调用的时候这种好处就会被放大，因为我们可以使用短名字，而省略掉包的名字
	perim2 := files.Path{
		files.Point{X: 1, Y: 1},
		files.Point{X: 5, Y: 1},
		files.Point{X: 5, Y: 4},
		files.Point{X: 1, Y: 1},
	}
	fmt.Println(files.PathDistance(perim2)) // 12
	fmt.Println(perim2.Distance())          // 12
}

// Point *
type Point struct{ X, Y float64 }

// Distance 函数
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance 方法，和 Distance 函数做同样的事，却是一个 Point 类型的方法
// 附加的参数 p，叫做方法的接收器（receiver），早期的面向对象语言留下的遗产，将调用一个方法称为“向一个对象发送消息”
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// 我们来定义一个 Path 类型，这个 Path 代表一个线段的集合，并且也给这个 Path 定义一个叫 Distance 的方法

// Path 类型
type Path []Point

// Distance - Path 类型的方法
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}
