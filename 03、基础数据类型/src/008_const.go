package main

import (
	"fmt"
	"math"
	"reflect"
	"time"
)

// 008、常量
// go run 008_const.go
func main() {
	// 常量表达式的值在编译期计算，而不是运行期
	const pi = 3.14159

	// 批量声明常量
	const (
		x = 1.23
		y = 4.56
	)

	// 常量间的所有算术运算、逻辑运算和比较运算的结果也是常量
	// 对常量的类型转换操作或进行以下函数调用返回的仍然是常量
	// len、cap、real、imag、complex 和 unsafe.Sizeof

	// 常量可以是构成类型的一部分
	// const IPv4Len = 4
	// var p [IPv4Len]byte

	// 常量的声明可以包含一个类型和一个值
	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %[1]v\n", noDelay)     // time.Duration 0s
	fmt.Printf("%T %[1]v\n", timeout)     // time.Duration 5m0s
	fmt.Printf("%T %[1]v\n", time.Minute) // time.Duration 1m0s

	// 如果是批量声明的常量，除了第一个外，其它常量右边的初始化表达式都可以省略
	// 如果省略初始化表达式，则表示使用前面常量的初始化表达式写法，对应的常量类型也一样
	const (
		a = 1
		b
		c = 2
		d
	)
	fmt.Println(a, b, c, d) // 1 1 2 2

	// iota 常量生成器
	type Weekday int
	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday) // 0 1 2 3 4 5 6

	fmt.Println(FlagUp, FlagBroadcast, FlagLoopback, FlagPointToPoint, FlagMulticast) // 1 2 4 8 16

	// netflag
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, isUp(v)) // 10001 true
	turnDown(&v)
	fmt.Printf("%b %t\n", v, isUp(v)) // 10000 false
	setBroadcast(&v)
	fmt.Printf("%b %t\n", v, isUp(v))   // 10010 false
	fmt.Printf("%b %t\n", v, isCast(v)) // 10010 true

	// 更复杂的例子，每个常量都是 1024 的幂
	const (
		_ = 1 << (10 * iota)
		Kib
		Mib
		Gib
		Tib
		Pib
		Eib
		Zib
		Yib
	)
	fmt.Println(Kib)          // 1024
	fmt.Println(Mib)          // 1048576
	fmt.Println(Gib)          // 1073741824
	fmt.Println(Tib)          // 1099511627776
	fmt.Println(Pib)          // 1125899906842624
	fmt.Println(Eib)          // 1152921504606846976
	fmt.Println(float64(Zib)) // 1.1805916207174113e+21
	fmt.Println(float64(Yib)) // 1.2089258196146292e+24

	// 无类型常量
	fmt.Println(Yib / Zib) // 1024，编译期计算出来，可以有效表示
	// math.Pi 的无类型浮点数常量，可以直接用于任意需要浮点数或复数的地方
	var xx float32 = math.Pi
	var yy float64 = math.Pi
	var zz complex128 = math.Pi
	fmt.Println(xx, yy, zz) // 3.1415927 3.141592653589793 (3.141592653589793+0i)
	// 如果 math.Pi 被确定为特定类型，比如 float64，那么结果精度可能会不一样
	// 同时对于需要 float32 或 complex128 类型值的地方则会强制需要一个明确的类型转换
	const Pi64 float64 = math.Pi
	var xxx float32 = float32(Pi64)
	var yyy float64 = Pi64
	var zzz complex128 = complex128(Pi64)
	fmt.Println(xxx, yyy, zzz) // 3.1415927 3.141592653589793 (3.141592653589793+0i)

	// 对于常量面值，不同的写法会对应不同的类型
	// 例如 0、0.0、0i 和 \u0000 虽然有相同的常量值
	// 但它们分别对应无类型的整数、无类型的浮点数、无类型的复数和无类型的字符
	// true 和 false 也是无类型的布尔类型，字符串面值常量是无类型的字符串类型

	// 除法运算符会根据操作数的类型，生成对应类型的结果
	// 因此，不同写法的除法常量表达式可能对应不同的结果
	var f float64 = 212
	fmt.Println((f-32)*5/9, reflect.TypeOf((f-32)*5))    // 100 float64
	fmt.Println(5/9*(f-32), reflect.TypeOf(5/9))         // 0 int
	fmt.Println(5.0/9.0*(f-32), reflect.TypeOf(5.0/9.0)) // 100 float64

	// 只有常量可以是无类型的
	// 当无类型的常量被赋值给变量时，无类型的常量将会被隐式转换为对应的类型（如果可以合法转换）
	// var ff float64 = 3 + 0i // untyped complex -> float64
	// ff = 2                  // untyped integer -> float64
	// ff = 1e123              // untyped floating-point -> float64
	// ff = 'a'                // untyped rune -> float64
	// 上面的语句相当于
	// var gg float64 = float64(3 + 0i)
	// gg = float64(2)
	// gg = float64(1e123)
	// gg = float64('a')

	// 无论是隐式或显式转换，将一种类型转换为另一种类型都要求目标可以表示原始值
	// 对于浮点数和复数，可能会有舍入处理
	const (
		deadbeef = 0xdeadbeef
		aaa      = uint32(deadbeef)
		bbb      = float32(deadbeef)
		ccc      = float64(deadbeef)
		// ddd      = int32(deadbeef) // constant 3735928559 overflows int32
		// eee      = float64(1e309)  // constant 1e+309 overflows float64
		// fff      = uint(-1)        // constant -1 overflows uint
	)
	fmt.Println(aaa) // 3735928559
	fmt.Println(bbb) // 3.7359286e+09
	fmt.Println(ccc) // 3.735928559e+09

	// 对于一个没有显式类型的变量声明，常量的形式将隐式决定变量的默认类型
	ii := 0                         // untyped integer
	rr := '\000'                    // untyped rune
	ff := 0.0                       // untyped floating-point
	cc := 0i                        // untyped complex
	fmt.Println(reflect.TypeOf(ii)) // implicit int，无类型整数常量转为 int 内存大小是不确定的
	fmt.Println(reflect.TypeOf(rr)) // implicit int32（rune）
	fmt.Println(reflect.TypeOf(ff)) // implicit float64，无类型的浮点数转为内存大小明确的 float64
	fmt.Println(reflect.TypeOf(cc)) // implicit complex128，无类型的复数转为内存大小明确的 complex128
	// 如果要给变量一个不同类型，必须显式指定类型，例如：
	// var xxx = int8(0)
	// var yyy int8 = 0
}

// Flags 类型定义
type Flags uint

// Flags 类型的常量
const (
	FlagUp Flags = 1 << iota
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

func isUp(v Flags) bool     { return v&FlagUp == FlagUp }
func turnDown(v *Flags)     { *v &^= FlagUp }
func setBroadcast(v *Flags) { *v |= FlagBroadcast }
func isCast(v Flags) bool   { return v&(FlagBroadcast|FlagMulticast) != 0 }
