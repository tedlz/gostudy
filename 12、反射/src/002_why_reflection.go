package main

import "strconv"

// 002、为何需要反射
func main() {

}

// 有时候我们需要编写一个函数能够处理一类并不满足普通公共接口的类型的值，也可能是因为它们并没有确定的表示方式，
// 或者是在我们设计该函数的时候这些类型可能还不存在

// 一个大家熟悉的例子是 fmt.Fprintf 提供的字符串格式化处理逻辑，
// 它可以用来对任意类型的值格式化并打印，甚至支持用户自定义的类型
// 让我们也来尝试实现一个类似功能的函数
// 为了简单起见，我们的函数只接收一个参数，然后返回和 fmt.Sprint 类似的格式化后的字符串，我们实现的函数名也叫 Sprint

// 我们首先用 switch 类型分支来测试输入参数是否实现了 String 方法，如果是的话就调用该方法
// 然后继续增加类型测试分支，检查这个值的动态类型是否为 string、int、bool 等基础类型，并在每种情况下执行相应的格式化操作

// Sprint *
func Sprint(x interface{}) string {
	type stringer interface {
		String() string
	}
	switch x := x.(type) {
	case stringer:
		return x.String()
	case string:
		return x
	case int:
		return strconv.Itoa(x)
	// int16、uint32 等类似情况……
	case bool:
		if x {
			return "true"
		}
		return "false"
	default:
		// array, chan, func, map, pointer, slice, struct
		return "???"
	}
}

// 但是我们如何处理其它类似 []float64、map[string][]string 等类型呢？
// 我们当然可以添加更多的测试分支，但是这些组合类型的数目基本是无穷的
// 还有如何处理类似 url.Values 这样的具名类型呢？
// 即使类型分支可以识别出底层的基础类型是 map[string][]string，但是它并不匹配 url.Values 类型，
// 因为它们是两种不同的类型，而且 switch 类型分支也不可能包含每个类似 url.Values 的类型，这会导致对这些库的依赖

// 没有办法检查对未知类型的表示方式，我们被卡住了
// 这就是我们为何需要反射的原因
