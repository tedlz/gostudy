package main

import (
	"fmt"
	"math"
)

// 002、浮点数
// go run 002_float.go
// 输出：
// 3.4028235e+38 1.7976931348623157e+308
// true
// 0.707 1
// x = 0 e^x =    1.000
// x = 1 e^x =    2.718
// x = 2 e^x =    7.389
// x = 3 e^x =   20.086
// x = 4 e^x =   54.598
// x = 5 e^x =  148.413
// x = 6 e^x =  403.429
// x = 7 e^x = 1096.633
// 0 -0 +Inf -Inf NaN
// false false false
// 0 false
func main() {
	var x float32
	var y float64
	x = math.MaxFloat32
	y = math.MaxFloat64
	fmt.Println(x, y) // 3.4028235e+38 1.7976931348623157e+308

	// float32 类型的浮点数可以提供大约 6 个十进制数的精度，float64 可以提供 15 个
	// 通常优先使用 float64 类型，float32 能精确表示的正整数不是很大
	// 因为 float32 有效 bit 位仅 23 个，其它的 bit 位用于指数和符号
	// 当整数大于 23bit 能表达的范围时，float32 的表现会出现误差
	var f float32 = 16777216
	fmt.Println(f == f+1) // true

	// 浮点数的字面值可以直接写小数部分
	const e = 2.71828 // 近似值
	// 小数点前面或后面的值都可能被省略
	a := .707
	b := 1.
	fmt.Println(a, b) // 0.707 1
	// 很小或很大的数最好用科学计数法书写，通过 e 或 E 来指定指数部分
	const Avogadro = 6.02214129e23 // 阿伏伽德罗常数
	const Planck = 6.62606957e-34  // 普朗克常数
	// 打印 e 的幂，精度为小数点后 3 位，以及 8 个字符宽度
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}

	// 特殊值
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) // 0 -0 +Inf -Inf NaN

	// 函数 math.isNaN 用于测试一个数是否是非数（NaN），math.NaN 则返回非数对应的值
	// 虽然可以用 math.NaN 来表示一个非法的结果，但测试结果是否是非数是充满风险的，因为 NaN 和任何数都不相等
	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan) // false false false

	fmt.Println(compute()) // 0 false
}

// 如果一个函数返回的浮点数结果可能失败，最好的做法是用单独的标志报告失败
func compute() (value float64, ok bool) {
	result, failed := 1., true
	if failed {
		return 0, false
	}
	return result, true
}
