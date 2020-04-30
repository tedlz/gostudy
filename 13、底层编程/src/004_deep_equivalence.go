package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

// 004、示例：深度相等判断
func main() {
	// 来自 reflect 包的 DeepEqual 函数可以对两个值进行深度相等判断
	// DeepEqual 函数使用内建的 == 比较操作符对基础类型进行相等判断，
	// 对于复合类型则递归该变量的每个基础类型然后做类似的比较判断
	// 因为它可以工作在任意的类型上，甚至对于一些不支持 == 操作运算符的类型也可以工作，因此在一些测试代码中广泛地使用该函数
	// 比如下面的代码是用 DeepEqual 函数比较两个字符串数组是否相等
	// （见下 TestSplit 方法）

	// 尽管 DeepEqual 函数很方便，而且可以支持任意的数据类型，但它也有不足之处
	// 例如，它将一个 nil 值的 map 和非 nil 值但是空的 map 视作不相等，
	// 同样 nil 值的 slice 和非 nil 值但是空的 slice 也视作不相等
	var a, b []string = nil, []string{}
	fmt.Println(reflect.DeepEqual(a, b)) // false

	var c, d map[string]int = nil, make(map[string]int)
	fmt.Println(reflect.DeepEqual(c, d)) // false

	// 我们希望在这里实现一个自己的 Equal 函数，用于比较类型的值
	// 和 DeepEqual 函数类似的地方是它也是基于 slice 和 map 的每个元素进行递归比较，
	// 不同之处是它将 nil 值的 slice（map 类似）和非 nil 值但是空的 slice 视为相等的值
	// 基础部分的比较可以基于 reflect 包完成，和 12.3 章的 Display 函数的实现方法类似
	// 同样，我们也定义了一个内部函数 equal，用于内部的递归比较
	// 读者目前不用关心 seen 参数的具体含义
	// 对于每一对需要比较的 x 和 y，equal 函数首先检测它们是否都有效（或都无效），然后检测它们是否是相同的类型
	// 剩下的部分是一个巨大的 switch 分支，用于相同基础类型的元素比较
	// 因为页面空间的限制，我们省略了一些相似的分支
	// （见 files/equal/equal.go 的 equal 函数）

	// 和前面的建议一样，我们并不公开 reflect 包相关的接口，所以导出的函数需要在内部自己将变量转为 reflect.Value 类型
	// （见 files/equal/equal.go 的 Equal 函数）

	// 为了确保算法对于有环的函数也能正常退出，我们必须记录每次已经比较的变量，从而避免进入第二次比较
	// Equal 函数分配了一组用于比较的结构体，包含每对比较对象的地址（unsafe.Pointer 形式保存）和类型
	// 我们要记录类型的原因是，有些不同的变量可能对应相同的地址
	// 例如，如果 x、y 都是数组类型，那么 x 和 x[0] 将对应相同的地址，y 和 y[0] 也是对应相同的地址，
	// 这可以用于区分 x 与 y 之间的比较或 x[0] 与 y[0] 之间的比较是否进行过了
	// （见 files/equal/equal.go 的 equal 函数的 [循环检查] 一段）

	// 这是 Equal 函数用法的例子：
	// （见 files/equal/equal_test.go 的 Example_equal 函数）

	// Equal 函数甚至可以处理类似 12.3 节中导致 Display 陷入死循环的带有环的数据：
	// （见 files/equal/equal_test.go 的 Example_equalCycle 函数）
}

// TestSplit *
func TestSplit(t *testing.T) {
	got := strings.Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		// ...
	}
}
