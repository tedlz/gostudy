package main

import (
	"crypto/sha256"
	"fmt"
)

// 001、数组
// go run 001_array.go
// 输出：
// 0
// 0
// 0 0
// 1 0
// 2 0
// 0
// 0
// 0
// 1 2 0
// [3]int
// 0 1 2 3
// $
// [0 0 0 0 0 0 0 0 0 -1]
// true false false
// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
// false
// [32]uint8
// [1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1]
// [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
// [2 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 2]
// [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
func main() {
	// 数组是由一个固定长度的特定类型元素组成的序列
	// 默认情况下，数组的每个元素都被初始化为元素类型对应的零值，对数字类型来说就是 0
	var a [3]int
	fmt.Println(a[0])        // 0
	fmt.Println(a[len(a)-1]) // 0
	// 打印键和值
	for i, v := range a {
		fmt.Println(i, v) // 0 0, 0 1, 0 2
	}
	// 仅打印值
	for _, v := range a {
		fmt.Println(v) // 0, 0, 0
	}

	// 初始化数组
	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(q[0], r[1], r[2]) // 1 2 0
	// 初始化数组，简写
	x := [...]int{1, 2, 3}
	fmt.Printf("%T\n", x) // [3]int

	// 数组的长度是数组类型的组成部分，因此 [3]int 和 [4]int 是两种不同的数组类型
	// 数组的长度必须是常量表达式，因为数组的长度需要在编译阶段确定
	// q = [4]int{1, 2, 3, 4} // cannot use [4]int literal (type [4]int) as type [3]int in assignment

	// 上面的形式是直接提供顺序初始化值序列
	// 以下为指定索引和值的方式初始化
	type Currency int
	const (
		USD = iota // 美元
		EUR        // 欧元
		GBP        // 英镑
		RMB        // 人民币
	)
	fmt.Println(USD, EUR, GBP, RMB) // 0 1 2 3
	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}
	fmt.Println(symbol[USD]) // $
	// 在这种方式中，初始化索引和值的顺序是无关紧要的
	// 例如定义了含有 10 个元素的数组 n，最后一个元素初始化的值为 -1，其余都是用 0 初始化
	n := [...]int{9: -1}
	fmt.Println(n) // [0 0 0 0 0 0 0 0 0 -1]

	// 如果一个数组的元素类型是可以相互比较的，那么数组的类型也是可以相互比较的
	d := [2]int{1, 2}
	e := [...]int{1, 2}
	f := [2]int{1, 3}
	fmt.Println(d == e, d == f, e == f) // true false false
	// g := [3]int{1, 2}
	// fmt.Println(d == g) // invalid operation: d == g (mismatched types [2]int and [3]int)

	sha256Test()

	ptr := [32]byte{0: 1, 31: 1}
	fmt.Println(ptr) // [1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1]
	zero(&ptr)
	fmt.Println(ptr) // [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	ptr = [32]byte{0: 2, 31: 2}
	fmt.Println(ptr) // [2 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 2]
	zero2(&ptr)
	fmt.Println(ptr) // [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
}

func sha256Test() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	// 输出：
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8

	// %x 以 16 进制方式打印
	// %t 打印 bool 类型数据
}

// 给 [32]byte 的数组类型清零
func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

// 给 [32]byte 的数组类型清零，简写
func zero2(ptr *[32]byte) {
	*ptr = [32]byte{}
}
