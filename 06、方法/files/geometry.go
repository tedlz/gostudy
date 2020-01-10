package files

import "math"

// Point *
type Point struct{ X, Y float64 }

// Distance 函数
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance 方法，和 Distance 函数做同样的事，却是一个 Point 类型的方法
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Path 类型
type Path []Point

// Distance - Path 类型 Distance 的方法
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

// PathDistance - Path 类型的 Distance 函数
func PathDistance(path Path) float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}
