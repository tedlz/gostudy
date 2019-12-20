package main

import "fmt"

// 002、切片
// go run 002_slice.go
// 输出：
// "" January December
// [April May June]
// [June July August]
// [April May xxx]
// [April May June]
// June
// June appears in both
// [June July August September October]
// [5 4 3 2 1 0]
// [0 1 2 3 4 5]
// [1 0 2 3 4 5]
// [1 0 5 4 3 2]
// [2 3 4 5 0 1]
// true false
// 0 true
// 0 true
// 0 true
// 0 false
// [0 0] 2 3
// [0 0] 2 2
// ['h' 'e' 'l' 'l' 'o' '，' '世' '界']
// ['h' 'e' 'l' 'l' 'o' '，' '世' '界']
// 0 cap=01 [0]
// 1 cap=02 [0 1]
// 2 cap=04 [0 1 2]
// 3 cap=04 [0 1 2 3]
// 4 cap=08 [0 1 2 3 4]
// 5 cap=08 [0 1 2 3 4 5]
// 6 cap=08 [0 1 2 3 4 5 6]
// 7 cap=08 [0 1 2 3 4 5 6 7]
// 8 cap=16 [0 1 2 3 4 5 6 7 8]
// 9 cap=16 [0 1 2 3 4 5 6 7 8 9]
// [1 2 3 4 5 6 1 2 3 4 5 6]
// [1 2 3 4 5 6 1 2 3 4 5 6 100 200 300]
// [one three]
// [one three three]
// [one two four]
// [one two four four]
// [1 2]
// 2
// [1]
// [5 6 8 9]
// [5 6 8 9 9]
// [5 6 9 9 9] [5 6 9 9]
func main() {
	// 切片代表变长的序列，序列中每个元素都有相同的类型
	// 一个切片是一个轻量级的数据结构，由三个部分组成：指针、长度和容量
	// 指针：指向第一个切片元素对应的底层数组元素的地址（切片的第一个元素不一定是数组的第一个元素）
	// 长度：对应切片中元素的数目，长度不能超过容量
	// 容量：从切片开始的位置到底层数据的结尾位置
	// 内置的 len 和 cap 函数分别返回切片的长度和容量

	months := [...]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}
	// months[0] 会被初始化为空字符串 ""
	fmt.Printf("%q %v %v\n", months[0], months[1], months[12]) // "" January December

	// 切片操作 s[i:j]，其中 0 <= i <= j <= cap(s)，用于创建一个新的 slice
	// 新 slice 从 i 个元素到第 j-1 个元素，共 j-i 个元素
	// 可以写做 s[:j]，省略的 i 会被 0 代替
	// 可以写做 s[i:]，省略的 j 会被 len(s) 代替
	// 因此 months[1:13] 将引用全部有效的月份，months[:] 则是引用整个数组
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)     // [April May June]
	fmt.Println(summer) // [June July August]
	// 切片是对原始数据的引用，当原始数据改变时，引用的数据也会变，反之亦然
	months[6] = "xxx"
	fmt.Println(Q2) // [April May xxx]
	Q2[2] = "June"
	fmt.Println(Q2)        // [April May June]
	fmt.Println(months[6]) // June

	outputSameMonth(Q2, summer) // June appears in both

	// fmt.Println(summer[20]) // panic: runtime error: index out of range [20] with length 3
	endlessSummer := summer[:5]
	fmt.Println(endlessSummer) // [June July August September October]

	// 字符串类型的切片操作和 []byte 字节类型切片的切片操作是类似的，都写做 x[m:n]
	// 并且都是返回一个字节系列的子序列，底层都是共享之前的底层数组，因此这种操作都是常量时间复杂度
	// x[m:n] 对字符串则生成一个新的字符串，对 []byte 则生成一个新 []byte

	// 反转
	s := []int{0, 1, 2, 3, 4, 5}
	reverse(s[:])  // 切片反转
	fmt.Println(s) // [5 4 3 2 1 0]
	reverse2(&s)   // 指针反转
	fmt.Println(s) // [0 1 2 3 4 5]
	reverse(s[:2])
	fmt.Println(s) // [1 0 2 3 4 5]
	reverse(s[2:])
	fmt.Println(s) // [1 0 5 4 3 2]
	reverse(s)
	fmt.Println(s) // [2 3 4 5 0 1]

	// 和数组不同的是，slice 之间不能比较
	// 标准库提供了 bytes.Equal 来判断两个 []byte 字节型 slice 是否相等
	// 对于其它类型的 slice，我们必须自己展开进行比较
	b := []int{2, 3, 4, 5, 0, 1}
	c := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(equal(s, b), equal(s, c)) // true false

	// slice 唯一合法的比较是和 nil 比较
	var z []int
	fmt.Println(len(z), z == nil) // 0 true
	z = nil
	fmt.Println(len(z), z == nil) // 0 true
	z = []int(nil)
	fmt.Println(len(z), z == nil) // 0 true
	z = []int{}
	fmt.Println(len(z), z == nil) // 0 false
	// 如果要测试一个 slice 是否为空，可以用 len(s) == 0 判断，不能用 s == nil 判断

	// 内置的 make 函数可以创建一个指定元素类型、长度和容量的 slice
	// 容量部分可以省略，若省略，此时容量 = 长度
	m := make([]int, 2, 3)
	n := make([]int, 2)
	fmt.Println(m, len(m), cap(m)) // [0 0] 2 3
	fmt.Println(n, len(n), cap(n)) // [0 0] 2 2

	// 内置的 append 函数用于向 slice 添加元素
	var runes []rune
	for _, r := range "hello，世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)    // ['h' 'e' 'l' 'l' 'o' '，' '世' '界']
	runes2 := []rune("hello，世界") // append 操作可以通过内置转换
	fmt.Printf("%q\n", runes2)   // ['h' 'e' 'l' 'l' 'o' '，' '世' '界']

	// 每次调用 appendInt 函数，必须先检测 slice 底层数组是否有足够的容量来保存新添加的元素
	// 如果有足够空间，在原有的底层数组之上直接扩展 slice，将新添加的元素 y 复制到新扩展的空间，并返回 slice
	// 此时输入的 x 和输出的 z 共享相同的底层数组
	// 如果没有足够的空间，appendInt 函数会先分配一个足够大的 slice 用来保存新的结果，先将输入的 x 复制到新空间再添加 y 元素
	// 此时输入的 x 和输出的 z 将引用不同的底层数组
	// 虽然通过循环复制元素更直接，不过内置的 copy 函数可以将一个 slice 复制到另一个相同类型的 slice
	// copy(z, x) 第一个参数是要复制的目标 slice，第二个参数是源 slice
	// 目标和源的位置和 dst = src 赋值语句是一致的，两个 slice 可以共享一个底层数组，甚至重叠也没关系
	// copy 函数将返回成功复制的元素个数，等于两个 slice 中较小的长度，所以不担心覆盖会超出目标 slice 的范围

	// 为了提高内存使用效率，新分配的数组一般略大于保存 x 和 y 所需要的最低大小
	// 通过每次扩展数组时直接将长度翻倍，从而避免了多次内存分配，也确保了添加单个元素的平均时间（原文见下）
	// （ensures that appending a single element takes constant time on average）
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%02d %v\n", i, cap(y), y)
		x = y
	}
	// 以上输出，可以看到每次容量变化都会导致重新分配内存和 copy 操作
	// 0 cap=01 [0]
	// 1 cap=02 [0 1]
	// 2 cap=04 [0 1 2]
	// 3 cap=04 [0 1 2 3]
	// 4 cap=08 [0 1 2 3 4]
	// 5 cap=08 [0 1 2 3 4 5]
	// 6 cap=08 [0 1 2 3 4 5 6]
	// 7 cap=08 [0 1 2 3 4 5 6 7]
	// 8 cap=16 [0 1 2 3 4 5 6 7 8]
	// 9 cap=16 [0 1 2 3 4 5 6 7 8 9]

	// int 型的 slice 类似下面的结构
	type IntSlice struct {
		ptr      *int
		len, cap int
	}

	// 内置的 append 可以追加多个元素，甚至追加一个 slice
	var xx []int
	xx = append(xx, 1)
	xx = append(xx, 2, 3)
	xx = append(xx, 4, 5, 6)
	xx = append(xx, xx...) // append the slice xx
	fmt.Println(xx)        // [1 2 3 4 5 6 1 2 3 4 5 6]
	xx = appendInt(xx, 100, 200, 300)
	fmt.Println(xx) // [1 2 3 4 5 6 1 2 3 4 5 6 100 200 300]

	// 过滤数组里的空值（输出的 slice 和 输入的 slice 共享一个底层数组）
	data := []string{"one", "", "three"}
	fmt.Println(nonempty(data)) // [one three]
	fmt.Println(data)           // [one three three]
	data2 := []string{"one", "two", "", "four"}
	fmt.Println(nonempty2(data2)) // [one two four]
	fmt.Println(data2)            // [one two four four]

	// 一个 slice 可以用来模拟一个 stack
	stack := []int{}
	stack = append(stack, 1)     // push 1
	stack = append(stack, 2)     // push 2
	fmt.Println(stack)           // [1 2]
	top := stack[len(stack)-1]   // stack 的顶部位置对应 slice 的最后一个元素
	fmt.Println(top)             // 2
	stack = stack[:len(stack)-1] // pop 2
	fmt.Println(stack)           // [1]

	// 删除 slice 中间的某个元素
	d := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(d, 2))     // [5 6 8 9]（保持顺序）
	fmt.Println(d)                // [5 6 8 9 9]
	fmt.Println(d, remove2(d, 2)) // [5 6 9 9 9], [5 6 9 9]（不保持顺序）
}

// 打印相同的月份
func outputSameMonth(a []string, b []string) {
	for _, x := range a {
		for _, y := range b {
			if x == y {
				fmt.Printf("%s appears in both\n", x)
			}
		}
	}
}

// 反转（切片）
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// 反转（指针）
func reverse2(s *[]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

// 比较切片是否相等
func equal(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

// 处理 int 类型的 slice
func appendInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		// 有成长的空间，扩展切片
		z = x[:zlen]
	} else {
		// 空间不足，分配一个新数组
		// 通过加倍增长，摊销线性复杂度
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // 把 z 复制到 x
	}
	copy(z[len(x):], y)
	return z
}

// 过滤数组里的空值
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

// 过滤数组里的空值 2
func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

// 删除 slice 中间的某个元素，并保持原有的元素顺序
func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:]) // 通过把后面的子 slice 向前移动一位
	return slice[:len(slice)-1]
}

// 删除 slice 中间的某个元素，且不需保持原有元素顺序
func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1] // 用最后的元素覆盖被删除的元素
	return slice[:len(slice)-1]
}
