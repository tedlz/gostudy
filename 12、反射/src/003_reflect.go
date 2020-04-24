package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

// 003、reflect.Type 和 reflect.Value
func main() {
	// 反射是由 reflect 包提供的
	// 它定义了两个重要的类型，Type 和 Value
	// ! Type 类型
	// 一个 Type 表示一个 Go 类型，
	// 它是一个接口，有许多方法来区分类型以及检查它们的组成部分，例如一个结构体的成员或一个函数的参数等
	// 唯一一个能反映 reflect.Type 实现的是接口的类型描述信息（7.5 节），也正是这个实体标识了接口值的动态类型

	// 函数 reflect.TypeOf 接受任意的 interface{} 类型，并以 reflect.Type 形式返回其动态类型
	t := reflect.TypeOf(3)  // a reflect.Type
	fmt.Println(t.String()) // "int"
	fmt.Println(t)          // "int"

	// 其中 TypeOf(3) 调用将值 3 传给 interface{} 参数
	// 回到 7.5 节，将一个具体的值转为接口类型会有一个隐式的接口转换操作，它会创建一个包含两个信息的接口值：
	// 操作数的动态类型（这里是 int）和它的动态的值（这里是 3）
	// 因为 reflect.TypeOf 返回的是一个动态类型的接口值，它总是返回具体的类型
	// 因此，下面的代码将打印 "*os.File" 而不是 "io.Writer"，稍后，我们将看到能够表达接口类型的 reflect.Type
	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w)) // "*os.File"

	// 要注意的是 reflect.Type 接口是满足 fmt.Stringer 接口的
	// 因为打印一个接口的动态类型对于调试和日志是有帮助的，
	// fmt.Printf 提供了一个缩写 %T 参数，内部使用 reflect.TypeOf 来输出：
	fmt.Printf("%T\n", 3) // "int"

	// ! Value 类型：
	// reflect 包中另一个重要的类型是 Value，一个 reflect.Value 可以装载任意类型的值
	// 函数 reflect.ValueOf 接受任意的 interface{} 类型，并返回一个装载着其动态值的 reflect.Value
	// 和 reflect.TypeOf 类似，reflect.ValueOf 返回的结果也是具体的类型，但是 reflect.Value 也可以持有一个接口值
	v := reflect.ValueOf(3) // a reflect.Value
	fmt.Println(v)          // "3"
	fmt.Printf("%v\n", v)   // "3"
	fmt.Println(v.String()) // NOTE: "<int value>"

	// 和 reflect.Type 类似，reflect.Value 也满足 fmt.Stringer 接口，
	// 但是除非 Value 持有的是字符串，否则 String 方法只返回其类型
	// 而使用 fmt 包的 %v 标志参数会对 reflect.Values 特殊处理
	// 对 Value 调用 type 方法将返回具体类型所对应的 reflect.Type：
	s := v.Type()           // a reflect.Type
	fmt.Println(s.String()) // "int"

	// reflect.ValueOf 的腻操作是 reflect.Value.Interface 方法
	// 它返回一个 interface{} 类型，装载着与 reflect.Value 相同的具体值：
	a := reflect.ValueOf(3) // a reflect.Value
	x := a.Interface()      // an interface
	i := x.(int)            // an int
	fmt.Printf("%d\n", i)   // "3"

	// reflect.Value 和 interface{} 都能装载任意的值
	// 所不同的是，一个空的接口隐藏了值内部的表示方式和所有方法，
	// 因此只有我们知道具体的动态类型才能使用类型断言来访问内部的值（就像上面那样），内部值我们没法访问
	// 相比之下，一个 Value 则有很多方法来检查其内容，无论它的具体类型是什么
	// 让我们再次尝试实现我们的格式化函数 format.Any

	// 我们使用 reflect.Value 的 Kind 方法来替代之前的类型 switch
	// 虽然还是有无穷多的类型，但它们的 kinds 类型却是有限的：
	// - Bool，String 和所有数字类型的基础类型；
	// - Array 和 Struct 对应的聚合类型；
	// - Chan、Func、Ptr、Slice 和 Map 对应的引用类型；
	// - Interface 类型；
	// - 表示空值的 Invalid 类型（空的 reflect.Value 的 kind 即为 Invalid）
	// （见 files/format）

	// 到目前为止，我们的函数将每个值视作一个不可分割没有内部结构的物品，因此它叫 formatAtom
	// 对于聚合类型（Struct、Array）和接口，只是打印值的类型，
	// 对于引用类型（channels、functions、pointers、slices、maps），打印类型和十六进制的引用地址
	// 虽然还不够理想，但依然是一个重大的进步，并且 Kind 只关心底层表示，format.Any 也支持具名类型
	// 例如：
	// （见 files/format/format_test.go 的 Test 方法）
}
