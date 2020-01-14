package main

import "fmt"

// 003、基于指针对象的方法 2
// go run 003_nil.go
// 输出：
// 1 + 0 = 1
// 0 + 0 = 0
func main() {
	// nil 也是一个合法的接收器类型
	i := IntList{1, nil} // {1, 0}
	fmt.Println(i.Sum()) // 1 + 0 = 1
	j := IntList{}       // {0, 0}
	fmt.Println(j.Sum()) // 0 + 0 = 0
}

// 就像一些函数允许 nil 指针作为参数一样，方法理论上也可以用 nil 指针作为其接收器
// 尤其当 nil 对于对象来说是合法的零值时，比如 map 或者 slice
// 在下面简单的 int 链表的例子里，nil 代表的是空链表

// IntList *
type IntList struct {
	Value int
	Tail  *IntList
}

// Sum *
func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	fmt.Printf("%d + %d = ", list.Value, list.Tail.Sum())
	return list.Value + list.Tail.Sum()
}
