package main

import (
	"fmt"
	"sort"
)

// 003、MAP
// go run 003_map.go
// 输出：
// 32
// map[charlie:34]
// 1
// charlie 34
// bob     1
// bob     1
// charlie 34
// true true
// 0
// ted 不存在
// 0
// ted 不存在
// true
// false
func main() {
	// 在 Go 语言中，一个 map 就是一个哈希表的引用
	// 内置的 make 函数可以创建一个 map
	ages := make(map[string]int)
	// 也可以用 map 字面值的语法创建 map
	ages = map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	// 相当于
	ages = map[string]int{}
	ages["alice"] = 31
	ages["charlie"] = 34
	// 访问 map 中的元素
	ages["alice"] = 32
	fmt.Println(ages["alice"]) // 32
	// 内置的 delete 函数可以删除元素
	delete(ages, "alice")
	fmt.Println(ages) // map[charlie:34]
	// 元素不在 map 中也可以，如果访问失败会返回一个 value 类型对应的零值
	ages["bob"] = ages["bob"] + 1
	fmt.Println(ages["bob"]) // 1

	// 简短赋值语法也可以用在 map 中，例如上面的例子可以写成 ages["bob"] += 1 或 ages["bob"]++
	// map 中的元素并不是一个变量，所以不可以这么写：
	// _ = &ages["bob"] // cannot take the address of ages["bob"]
	// 禁止对 map 元素取址的原因是 map 可能随着元素数量增长而分配更大的内存空间，从而导致之前的地址无效

	// map 的遍历
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}
	// 以上循环输出：
	// bob     1
	// charlie 34
	// 或
	// charlie 34
	// bob     1

	// map 的迭代顺序是不确定的，并且不同的哈希函数实现可能导致不同的遍历顺序
	// 在实践中，遍历的顺序是随机的，每次遍历的顺序都不同
	// 这是故意的，每次都使用随机的遍历顺序可以强制要求程序不会依赖具体的哈希函数实现
	// 如果要按顺序遍历 key/value 时，必须显式的对 key 进行排序，例如使用 sort.Strings 对字符串 slice 排序：
	var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
	// 以上循环输出：
	// bob     1
	// charlie 34

	// 因为我们一开始就知道 names 的大小，因此给 slice 分配一个合适的大小更有效
	// 下面的代码创建了一个空的 slice，但是 slice 的容量刚好可以放下 map 中全部的 key
	names = make([]string, 0, len(ages))

	// map 类型的零值是 nil，也就是没有引用任何哈希表
	var ages2 map[string]int
	fmt.Println(ages2 == nil, len(ages2) == 0) // true true
	// map 的大部分操作，包括查找、删除、len 和 range 都可以安全工作在 nil 值的 map 上，它们的行为和一个空的 map 类似
	// 但是向一个 nil 值的 map 存入元素将导致 panic 异常
	// ages2["carol"] = 21 // panic: assignment to entry in nil map
	// 在向 map 存储数据前，必须先创建 map

	// 通过 map[key] 访问将产生一个 value
	// 如果 key 存在于 map，map[key] 将得到与 key 对应的 value
	// 如果 key 不存在，map[key] 将得到 value 对应类型的零值，例如刚才的 ages["bob"]

	// 判断 map 中的元素是否存在
	// 在这种情况下，ages["ted"] 将产生两个值，第一个 value 对应类型的零值，第二个为布尔值
	age, ok := ages["ted"]
	if !ok {
		fmt.Println(age)       // 0
		fmt.Println("ted 不存在") // ted 不存在
	}
	// 通常会简写为
	if age, ok := ages["ted"]; !ok {
		fmt.Println(age)       // 0
		fmt.Println("ted 不存在") // ted 不存在
	}

	// 和 slice 一样，map 也不能进行相等比较；唯一例外的是和 nil 比较
	// 判断两个 map 是否包含相同的 value
	a := map[string]int{
		"bob": 1,
		"lol": 2,
	}
	var b = map[string]int{
		"bob": 1,
		"lol": 2,
	}
	fmt.Println(equal2(a, b)) // true
	delete(b, "bob")
	fmt.Println(equal2(a, b)) // false
}

// 判断两个 map 是否包含相同的 value
func equal2(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
