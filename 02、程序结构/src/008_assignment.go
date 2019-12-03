package main

import "fmt"

// 008、赋值
// go run 008_assignment.go
// 输出：
// 一、赋值
// x = 1                       // 命名变量的赋值
// *p = true                   // 通过指针间接赋值
// person.name = "bob"         // 结构体字段赋值
// count[x] = count[x] * scale // 数组、slice 或 map 的元素赋值，可以简写为 count[x] *= scale

// 1
// 2
// 1

// 二、元组赋值
// 1 2
// 2 1
// [1 2]
// [2 1]

// 3
// 46368
// 2 3 5
// f, err = os.Open(file) // 当表达式有多个值时，左边变量数必须和右边一致
// 通常这些函数会用额外的返回值来表示某种错误类型，但也有一些会返回布尔值，通常被称为 ok，例如：
// v, ok = m[key] // map 查找（map lookup）
// v, ok = x.(T)  // 类型断言（type assertion）
// v, ok = <-ch   // 通道接收（channel receive）
// map 查找，类型断言和通道接收也可能只产生一个结果：
// v = m[key] // map 查找，失败时返回零值
// v = x.(T)  // 类型断言，失败时 panic 异常
// v = <-ch   // 管道接收，失败时返回零值（阻塞不算是失败）
// _, ok = m[key]        // map 返回 2 个值
// _, ok = mm[""], false // map 返回 1 个值
// _ = mm[""]            // map 返回 1 个值
// 和变量声明一样，我们可以用下划线空白标识符来丢弃不需要的值：
// _, err = io.Copy(dst, src) // 丢弃字节数
// _, ok = x.(T)              // 只检测类型，忽略具体值

// 三、可赋值性
// 隐式地对 slice 赋值：
// medals := []string{"gold", "silver", "bronze"}
// 以上相当于：
// medals[0] = "gold"
// medals[1] = "silver"
// medals[2] = "bronze"
// map 和 chan 的元素，虽然不是普通变量，但也有类似的隐式赋值行为
// 可赋值性的规则：
// 1、类型必须完全匹配；
// 2、nil 可以赋值给任何指针或引用类型的变量；
// 3、常量有更灵活的赋值规则，可以避免不必要的显式类型转换。
func main() {
	// 一、赋值
	fmt.Println("一、赋值")
	example := []string{
		"x = 1                       // 命名变量的赋值",
		"*p = true                   // 通过指针间接赋值",
		"person.name = \"bob\"         // 结构体字段赋值",
		"count[x] = count[x] * scale // 数组、slice 或 map 的元素赋值，可以简写为 count[x] *= scale",
	}
	for _, arg := range example {
		fmt.Println(arg)
	}
	fmt.Println()

	v := 1
	fmt.Println(v) // 1
	v++
	fmt.Println(v) // 等价方式 v = v + 1，v 变成 2
	v--
	fmt.Println(v) // 等价方式 v = v - 1，v 变成 1
	fmt.Println()

	// 二、元组赋值
	fmt.Println("二、元组赋值")

	// 交换赋值
	x := 1
	y := 2
	fmt.Println(x, y) // 1 2
	x, y = y, x
	fmt.Println(x, y) // 2 1
	z := []int{1, 2}
	fmt.Println(z) // [1, 2]
	z[0], z[1] = z[1], z[0]
	fmt.Println(z) // [2, 1]
	fmt.Println()

	// 计算最大公约数
	fmt.Println(gcd(24, 9))

	// 计算斐波那契数列的第 n 个数
	fmt.Println(fib(24))

	// 元祖赋值也可使一系列的赋值更加紧凑，若过长请单独赋值以保持可读性
	i, j, k := 2, 3, 5
	fmt.Println(i, j, k)

	example = []string{
		"f, err = os.Open(file) // 当表达式有多个值时，左边变量数必须和右边一致",
		"通常这些函数会用额外的返回值来表示某种错误类型，但也有一些会返回布尔值，通常被称为 ok，例如：",
		"v, ok = m[key] // map 查找（map lookup）",
		"v, ok = x.(T)  // 类型断言（type assertion）",
		"v, ok = <-ch   // 通道接收（channel receive）",
		"map 查找，类型断言和通道接收也可能只产生一个结果：",
		"v = m[key] // map 查找，失败时返回零值",
		"v = x.(T)  // 类型断言，失败时 panic 异常",
		"v = <-ch   // 管道接收，失败时返回零值（阻塞不算是失败）",
		"_, ok = m[key]        // map 返回 2 个值",
		"_, ok = mm[\"\"], false // map 返回 1 个值",
		"_ = mm[\"\"]            // map 返回 1 个值",
		"和变量声明一样，我们可以用下划线空白标识符来丢弃不需要的值：",
		"_, err = io.Copy(dst, src) // 丢弃字节数",
		"_, ok = x.(T)              // 只检测类型，忽略具体值",
	}
	for _, arg := range example {
		fmt.Println(arg)
	}
	fmt.Println()

	// 三、可赋值性
	fmt.Println("三、可赋值性")
	example = []string{
		"隐式地对 slice 赋值：",
		"medals := []string{\"gold\", \"silver\", \"bronze\"}",
		"以上相当于：",
		"medals[0] = \"gold\"",
		"medals[1] = \"silver\"",
		"medals[2] = \"bronze\"",
		"map 和 chan 的元素，虽然不是普通变量，但也有类似的隐式赋值行为",
		"可赋值性的规则：",
		"1、类型必须完全匹配；",
		"2、nil 可以赋值给任何指针或引用类型的变量；",
		"3、常量有更灵活的赋值规则，可以避免不必要的显式类型转换。",
	}
	for _, arg := range example {
		fmt.Println(arg)
	}
}

// 计算最大公约数 greatest common divisor
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// 计算斐波那契数列的第 n 个数 fibonacci
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
