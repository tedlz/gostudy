package main

import (
	"fmt"
	"math"
)

// 001、整型
// go run 001_integer.go
// 输出：
// 有符号类型：
// int8(Min=-128, Max=127)
// int16(Min=-32768, Max=32767)
// int32(Min=-2147483648, Max=2147483647)
// int64(Min=-9223372036854775808, Max=9223372036854775807)
// 无符号类型：
// uint8(Min=0, Max=255)
// uint16(Min=0, Max=65535)
// uint32(Min=0, Max=4294967295)
// uint64(Min=0, Max=18446744073709551615)
// 特殊类型：
// var x rune    // Unicode 类型，等价于 int32，通常用来表示一个 Unicode 码点
// var y byte    // 等价于 uint8
// var z uintptr // 没有具体的大小，但足以容纳指针，只在底层编程时需要
// var a int     // 可变类型
// var b uint    // 可变类型
// 二元运算符优先级（使用括号可提升优先级）：
// *       %       <<      >>      &       &^
// +       -       |       ^
// ==      !=      <       <=      >       >=
// &&
// ||
// 算术运算符 + - * / 可用于整数、浮点数和复数，% 仅能用于整数
// -5%3 == -5%-3 的结果为 true
// 255 0 1
// 127 -128 1
// 两个相同的整数类型可以用 == != > >= < <= 比较运算符比较，结果是 bool 类型
// 对于整数，+x 是 0+x 的简写，-x 是 0-x 的简写；对于浮点数和负数，+x 就是 x，-x 就是 x 的负数
// 位运算符：
// &  位运算 AND
// |  位运算 OR
// ^  位运算 XOR
// &^ 位清空
// << 左移
// >> 右移
// 00100010
// 00000110
// 00000010
// 00100110
// 00100100
// 00100000
// 1
// 5
// 01000100
// 00010001
// bronze  silver  gold
// 整数类型的转换不会改变值，只会告诉编译器如何解释该值
// 将一个大尺寸整数类型转换为小尺寸的整数类型，或者将一个浮点数转换为整数，可能会改变数值或者丢失精度
// 3.141 3
// 1
// -9223372036854775808
// 任何大小的整数字面值都可以用 0 开始的八进制书写，例如 0666
// 或用以 0x 或 0X 开头的十六进制书写，例如 0xdeadbeef
// 438 666 0666
// 3735928559 deadbeef 0xdeadbeef
// 97 a 'a'
// 22269 国 '国'
// 10 '\n'
func main() {
	// 有符号类型
	var int8Min, int8Max int8
	var int16Min, int16Max int16
	var int32Min, int32Max int32
	var int64Min, int64Max int64
	int8Min, int8Max = math.MinInt8, math.MaxInt8
	int16Min, int16Max = math.MinInt16, math.MaxInt16
	int32Min, int32Max = math.MinInt32, math.MaxInt32
	int64Min, int64Max = math.MinInt64, math.MaxInt64
	fmt.Println("有符号类型：")
	fmt.Printf("int8(Min=%v, Max=%v)\n", int8Min, int8Max)    // int8(Min=-128, Max=127)
	fmt.Printf("int16(Min=%v, Max=%v)\n", int16Min, int16Max) // int16(Min=-32768, Max=32767)
	fmt.Printf("int32(Min=%v, Max=%v)\n", int32Min, int32Max) // int32(Min=-2147483648, Max=2147483647)
	fmt.Printf("int64(Min=%v, Max=%v)\n", int64Min, int64Max) // int64(Min=-9223372036854775808, Max=9223372036854775807)

	// 无符号类型
	var uint8Min, uint8Max uint8
	var uint16Min, uint16Max uint16
	var uint32Min, uint32Max uint32
	var uint64Min, uint64Max uint64
	uint8Min, uint8Max = 0, math.MaxUint8
	uint16Min, uint16Max = 0, math.MaxUint16
	uint32Min, uint32Max = 0, math.MaxUint32
	uint64Min, uint64Max = 0, math.MaxUint64
	fmt.Println("无符号类型：")
	fmt.Printf("uint8(Min=%v, Max=%v)\n", uint8Min, uint8Max)    // uint8(Min=0, Max=255)
	fmt.Printf("uint16(Min=%v, Max=%v)\n", uint16Min, uint16Max) // uint16(Min=0, Max=65535)
	fmt.Printf("uint32(Min=%v, Max=%v)\n", uint32Min, uint32Max) // uint32(Min=0, Max=4294967295)
	fmt.Printf("uint64(Min=%v, Max=%v)\n", uint64Min, uint64Max) // uint64(Min=0, Max=18446744073709551615)

	// 特殊类型
	fmt.Println("特殊类型：")
	fmt.Println("var x rune    // Unicode 类型，等价于 int32，通常用来表示一个 Unicode 码点")
	fmt.Println("var y byte    // 等价于 uint8")
	fmt.Println("var z uintptr // 没有具体的大小，但足以容纳指针，只在底层编程时需要")
	fmt.Println("var a int     // 可变类型")
	fmt.Println("var b uint    // 可变类型")

	// 优先级，二元运算符，优先级从上到下，从左到右递减
	fmt.Println("二元运算符优先级（使用括号可提升优先级）：")
	fmt.Println("*\t%\t<<\t>>\t&\t&^")
	fmt.Println("+\t-\t|\t^")
	fmt.Println("==\t!=\t<\t<=\t>\t>=")
	fmt.Println("&&")
	fmt.Println("||")

	fmt.Println("算术运算符 + - * / 可用于整数、浮点数和复数，% 仅能用于整数")
	fmt.Println("-5%3 == -5%-3 的结果为", -5%3 == -5%-3)

	// 计算溢出，超出高位的 bit 部分会被丢弃
	var u uint8 = 255
	fmt.Println(u, u+1, u*u) // 255 0 1
	var i int8 = 127
	fmt.Println(i, i+1, i*i) // 127 -128 1
	fmt.Println("两个相同的整数类型可以用 == != > >= < <= 比较运算符比较，结果是 bool 类型")
	fmt.Println("对于整数，+x 是 0+x 的简写，-x 是 0-x 的简写；对于浮点数和负数，+x 就是 x，-x 就是 x 的负数")
	fmt.Println("位运算符：")
	fmt.Println("&  位运算 AND")
	fmt.Println("|  位运算 OR")
	fmt.Println("^  位运算 XOR")
	fmt.Println("&^ 位清空")
	fmt.Println("<< 左移")
	fmt.Println(">> 右移")

	// 使用位操作解释 uint8 类型值的 8 个独立的 bit 位
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("%08b\n", x)    // 00100010
	fmt.Printf("%08b\n", y)    // 00000110
	fmt.Printf("%08b\n", x&y)  // 00000010
	fmt.Printf("%08b\n", x|y)  // 00100110
	fmt.Printf("%08b\n", x^y)  // 00100100
	fmt.Printf("%08b\n", x&^y) // 00100000
	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Println(i) // 1 5
		}
	}
	fmt.Printf("%08b\n", x<<1) // 01000100
	fmt.Printf("%08b\n", x>>1) // 00010001

	// len 函数会返回一个有符号的 int，例如用来处理逆序循环
	medals := []string{"gold", "silver", "bronze"}
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Print(medals[i], "\t") // bronze silver gold
	}
	fmt.Println()

	// 转换
	fmt.Println("整数类型的转换不会改变值，只会告诉编译器如何解释该值")
	fmt.Println("将一个大尺寸整数类型转换为小尺寸的整数类型，或者将一个浮点数转换为整数，可能会改变数值或者丢失精度")
	tf := 3.141
	ti := int(tf)
	fmt.Println(tf, ti) // 1.141 3
	tf = 1.99
	fmt.Println(int(tf)) // 1
	tf = 1e100           // float
	fmt.Println(int(tf)) // 结果依赖于具体实现
	fmt.Println("任何大小的整数字面值都可以用 0 开始的八进制书写，例如 0666")
	fmt.Println("或用以 0x 或 0X 开头的十六进制书写，例如 0xdeadbeef")
	oo := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", oo) // 438 666 0666
	xx := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x\n", xx) // 3735928559 deadbeef 0xdeadbeef

	// 通过转义的数值来表示任意 Unicode 码点对应的字符
	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)   // 97 a 'a'
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // 22269 国 '国'
	fmt.Printf("%d %[1]q\n", newline)       // 10 '\n'
}
