package main

import (
	"bytes"
	"fmt"
)

// 008、示例：Bit 数组
// go run 008_intset.go
// 输出：
// {1 9 144}
// {9 42}
// {1 9 42 144}
// true false
// {1 2 3 4 9 42 144}
// 7
// {1 2 3 9 42 144}
// {}
// {1 2 3 9 42 144}
func main() {
	// Go 语言里的集合一般会用 map[T]bool 这种形式来表示，T 代表元素类型
	// 集合用 map 类型来表示虽然非常灵活，但我们可以以一种更好的形式来表示它

	// 例如在数据流分析领域，集合元素通常是一个非负整数，集合会包含很多元素，并且集合会经常进行并集、交集操作
	// 这种情况下，bit 数组会比 map 表现更加理想（译注：比如我们执行一个 http 下载任务，把文件按照 16KB 一块划分为很多块
	// 需要有一个全局变量来标识哪些块下载完成了，这时也需要用到 bit 数组）

	// 一个 bit 数组通常会用一个无符号数或者称之为 “字” 的 slice 来表示，每一个元素的每一位都表示集合里的一个值
	// 当集合的第 i 位被设置时，我们才说这个集合包含元素 i
	// 下面的这个程序展示了一个简单的 bit 数组类型，并且实现了三个函数来对这个 bit 数组进行操作

	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // {1 9 144}

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // {9 42}

	x.UnionWith(&y)
	fmt.Println(x.String())           // {1 9 42 144}
	fmt.Println(x.Has(9), x.Has(123)) // true false

	x.AddAll(2, 3, 4)
	fmt.Println(x.String()) // {1 2 3 4 9 42 144}
	fmt.Println(x.Len())    // 7
	x.Remove(4)
	fmt.Println(x.String()) // {1 2 3 9 42 144}

	z := x.Copy()
	x.Clear()
	fmt.Println(x.String()) // {}
	fmt.Println(z.String()) // {1 2 3 9 42 144}
}

// IntSet 是一组小的非负整数
// 它的零值代表空集
type IntSet struct {
	words []uint64
}

// Has 报告集合是否包含非负值 x
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add 将非负值 x 添加到集合中
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith 将 s 设置为 s 和 t 的并集
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String 以字符串形式返回集合
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// AddAll 将多个非负值 x 添加到集合中
func (s *IntSet) AddAll(xs ...int) {
	for _, x := range xs {
		s.Add(x)
	}
}

// Len 集合元素数量
func (s *IntSet) Len() int {
	ret := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				ret++
			}
		}
	}
	return ret
}

// Remove 从集合中删除某元素
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	s.words[word] &= ^(1 << bit)
}

// Clear 清空集合
func (s *IntSet) Clear() {
	s.words = s.words[:0]
}

// Copy 拷贝
func (s *IntSet) Copy() *IntSet {
	ret := make([]uint64, len(s.words))
	copy(ret, s.words)
	return &IntSet{ret}
}
