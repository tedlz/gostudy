package main

import (
	"fmt"
	"math/cmplx"
)

// 004、复数
// go run 004_complex.go
// 输出：
// (-5+10i)
// -5
// 10
// (-1+0i)
// (0+1i)
func main() {
	// go 提供了两种精度的复数类型（complex64, complex128）
	// 分别对应 float32 和 float64 两种浮点数精度
	// complex 函数用来构建复数，real 和 imag 函数分别返回复数的实部和虚部
	var x complex128 = complex(1, 2) // 1+2i
	var y complex128 = complex(3, 4) // 3+4i
	fmt.Println(x * y)               // -5+10i
	fmt.Println(real(x * y))         // -5
	fmt.Println(imag(x * y))         // 10

	// 如果一个浮点数面值或一个整数后面跟着 i，如 3.14i 或 2i，它将构成一个复数的虚部，复数的实部是 0
	fmt.Println(1i * 1i) // (-1+0i)

	// 在常量算术规则下，一个复数常量可以与一个普通常量相加（整数或浮点数、实部或虚部）
	// 我们可以用自然的方式书写复数，就像 1+2i 或与之等价的写法 2i+1，上面 x 和 y 的声明语句还可以简化
	x = 1 + 2i
	y = 3 + 4i

	// 复数可以用 == 和 != 来比较，只有两个复数的实部和虚部都相等时，两个复数才相等（浮点数的相等小心精度问题）
	// math.cmplx 包提供了复数处理的许多函数，例如求复数的平方根函数和求幂函数
	fmt.Println(cmplx.Sqrt(-1)) // (0+1i)
}
