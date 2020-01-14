package main

import (
	"fmt"
	"gostudy/06、方法/files"
)

// 004、基于指针对象的方法 3
// go run 004_urlvalues.go
// 输出：
// en
// ""
// 1
// [1 2]
// ""
func main() {
	m := files.Values{"lang": {"en"}}
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang"))     // en
	fmt.Printf("%q\n", m.Get("q")) // ""
	fmt.Println(m.Get("item"))     // 1
	fmt.Println(m["item"])         // [1 2]

	m = nil
	fmt.Printf("%q\n", m.Get("item")) // ""
	// m.Add("item", "3") // panic: assignment to entry in nil map

	// 在对 Get 的最后一次调用中，nil 接收器的行为即是一个空 map 的行为
	// 我们可以等价的将这个操作写成 Values(nil).Get("item")
	// 如果直接写 nil.Get("item") 是无法通过编译的，因为 nil 的字面量编译器无法判断其准备类型
}
