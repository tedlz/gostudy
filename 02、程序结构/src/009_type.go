package main

import "fmt"

type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸点温度
)

// 009、类型转换
// go run 009_type.go
// 输出：
// 100
// 180
// true
// true
// true
// - 显式的调用了 Celsius.String() 方法： 100℃
// - 隐式的调用了 Celsius.String() 方法： 100℃
// - 隐式的调用了 Celsius.String() 方法： 100℃
// - 隐式的调用了 Celsius.String() 方法： 100℃
// 未调用 Celsius.String() 方法： 100
// 未调用 Celsius.String() 方法： 100
// - 显式的调用了 Fahrenheit.String() 方法： 212℉
// - 隐式的调用了 Fahrenheit.String() 方法： 212℉
// - 隐式的调用了 Fahrenheit.String() 方法： 212℉
// - 隐式的调用了 Fahrenheit.String() 方法： 212℉
// 未调用 Fahrenheit.String() 方法： 212
// 未调用 Fahrenheit.String() 方法： 212
func main() {
	fmt.Printf("%g\n", BoilingC-FreezingC)       // 100
	boilingF := CToF(BoilingC)                   // 摄氏度类型转为华氏度类型
	fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // 180
	// 以下语句会报错【无效的操作：boilingF - FreezingC（Fahrenheit 和 Celsius 类型不匹配）】
	// invalid operation: boilingF - FreezingC (mismatched types Fahrenheit and Celsius)
	// fmt.Printf("%g\n", boilingF-FreezingC)

	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0)          // true
	fmt.Println(f >= 0)          // true
	fmt.Println(c == Celsius(f)) // true
	// 以下语句会报错【无效的操作：c == f（Celsius 和 Fahrenheit 类型不匹配）】
	// invalid operation: c == f (mismatched types Celsius and Fahrenheit)
	// fmt.Println(c == f)

	c = FToC(212.0)
	fmt.Println("显式的调用了 Celsius.String() 方法：", c.String())
	fmt.Printf("隐式的调用了 Celsius.String() 方法： %v\n", c)
	fmt.Printf("隐式的调用了 Celsius.String() 方法： %s\n", c)
	fmt.Println("隐式的调用了 Celsius.String() 方法：", c)
	fmt.Printf("未调用 Celsius.String() 方法： %g\n", c)
	fmt.Println("未调用 Celsius.String() 方法：", float64(c))

	f = CToF(100.0)
	fmt.Println("显式的调用了 Fahrenheit.String() 方法：", f.String())
	fmt.Printf("隐式的调用了 Fahrenheit.String() 方法： %v\n", f)
	fmt.Printf("隐式的调用了 Fahrenheit.String() 方法： %s\n", f)
	fmt.Println("隐式的调用了 Fahrenheit.String() 方法：", f)
	fmt.Printf("未调用 Fahrenheit.String() 方法： %g\n", f)
	fmt.Println("未调用 Fahrenheit.String() 方法：", float64(f))
}

// 摄氏度转华氏度
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}
func (c Celsius) String() string {
	fmt.Print("- ")
	return fmt.Sprintf("%g℃", c)
}

// 华氏度转摄氏度
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
func (f Fahrenheit) String() string {
	fmt.Print("- ")
	return fmt.Sprintf("%g℉", f)
}
