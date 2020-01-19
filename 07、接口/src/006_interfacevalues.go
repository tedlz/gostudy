package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// 006、接口值
// go run 006_interfacevalues.go
// 输出：
// hello
// hello
// <nil>
// *os.File
// *bytes.Buffer
// <nil>
func main() {
	// 接口值有两个部分组成，一个具体的类型和那个类型的值。它们被称为接口的动态类型和动态值
	// 对于像 Go 语言这样静态类型的语言，类型是编译期的概念；因此一个类型不是一个值
	// 在我们的概念模型中，一些提供每个类型信息的值被称为类型描述符，比如类型的名称和方法
	// 在一个接口值中，类型部分代表与之相关类型的描述符（开始和最后的值是相同的）

	// 下面语句中，变量 w 得到了 3 个不同的值
	// 在 Go 语言中，变量总是被一个定义明确的值初始化，即使接口类型也不例外
	// 对于一个接口的零值就是它的类型和值的部分都是 nil
	var w io.Writer

	// 一个接口值基于它的动态类型被描述为空或非空，所以这是一个空的接口值
	// 可以通过 w == nil 或 w != nil 来判断接口值是否为空，调用一个空接口值的任意方法都会产生 panic
	// w.Write([]byte("hello")) // panic: runtime error: invalid memory address or nil pointer dereference

	// 这个赋值过程调用了一个具体类型到接口类型的隐式转换，这和显式使用 io.Writer(os.Stdout) 是等价的
	// 这类转换不管是显式的还是隐式的，都会刻画出操作到的类型和值
	// 这个接口值的动态类型被设为 *os.Stdout 指针的类型描述符，它的动态值持有 os.Stdout 的拷贝
	// 这是一个代表处理标准输出的 os.File 类型变量的指针
	w = os.Stdout
	w.Write([]byte("hello\n")) // hello

	// 通常在编译期，我们不知道接口值的动态类型是什么，所以一个接口上的调用必须使用动态分配
	// 因为不是直接调用，所以编译器必须把代码生成在类型描述符的方法 Write 上，然后间接调用那个地址
	// 这个调用的接收者是一个接口动态值的拷贝，os.Stdout。和下面这个直接调用一样：
	os.Stdout.Write([]byte("hello\n")) // hello

	// 第三个语句给接口赋值了 *bytes.Buffer 类型的值
	// 现在动态类型是 *bytes.Buffer 并且动态值是一个指向新分配的缓冲区的指针
	w = new(bytes.Buffer)
	// Write 方法的调用也使用了和之前一样的机制：
	w.Write([]byte("hello")) // 将 hello 写入 bytes.Buffer
	// 这次类型描述符是 *bytes.Buffer，所以调用了 (*bytes.Buffer).Write 方法，并且接收者是该缓冲区的地址
	// 这个调用把字符串 hello 添加到缓冲区中

	// 第四个语句将 nil 赋给了接口值
	// 这个重置将它所有的部分都设为 nil 值，把变量 w 恢复到和他之前定义时相同的状态
	w = nil

	// 一个接口值可以持有任意大的动态值
	// 例如，表示时间实例的 time.Time 类型，这个类型有几个不对外公开的字段，我们从它上面创建一个接口值
	// var x interface{} = time.Now()

	// 接口值可以通过 == 和 != 来进行比较
	// 两个接口值相等仅当它们都是 nil 值，或它们的动态类型相同且动态值也根据这个动态类型的 == 操作相等
	// 因为接口值是可比较的，所以它们可以用在 map 的键或者作为 switch 语句的操作数

	// 然而，如果两个接口值的动态类型相同，但是这个动态类型是不可比较的，比如 slice 切片，将它们比较就会失败并且 panic：
	// var x interface{} = []int{1, 2, 3}
	// fmt.Println(x == x) // panic: runtime error: comparing uncomparable type []int

	// 考虑到这点，接口类型是非常与众不同的
	// 其它类型要么是安全的可比较类型（如基本类型和指针），要么是完全不可比较的类型（如切片、映射类型和函数）
	// 但是在比较接口值或者包含了接口值的聚合类型时，我们必须要意识到潜在的 panic
	// 同样的风险也存在于使用接口作为 map 的键或者 switch 的操作数，只能比较你非常确定它们的动态值是可比较类型的接口值

	// 当我们处理错误或者调试的过程中，得知接口值的动态类型，是非常有帮助的
	var q io.Writer
	fmt.Printf("%T\n", q) // <nil>
	q = os.Stdout
	fmt.Printf("%T\n", q) // *os.File
	q = new(bytes.Buffer)
	fmt.Printf("%T\n", q) // *bytes.Buffer
	q = nil
	fmt.Printf("%T\n", q) // <nil>

	// 警告：一个包含 nil 指针的接口不是 nil 接口
	// 一个不包含任何值的 nil 接口值和一个刚好包含 nil 指针的接口值是不同的
	// 这个细微的区别产生了一个容易绊倒 Go 程序员的陷阱
	// 思考下面的程序，当 debug 变量设置为 true 时，main 函数会将 f 函数的输出收集到一个 bytes.Buffer 的类型中
	var buf *bytes.Buffer
	if debug {
		buf = new(bytes.Buffer)
	}
	f(buf)
	if debug {
		// use buf...
	}
	// 我们可能会预计，当把变量 debug 设置为 false 时可以禁止对输出的收集
	// 但是实际上在 out.Write 方法调用时发生了 panic：
	// panic: runtime error: invalid memory address or nil pointer dereference

	// 当 main 函数调用函数时，它给 f 函数的 out 参数赋了一个 *bytes.Buffer 的空指针，所以 out 的动态值是 nil
	// 然而，它的动态类型是 *bytes.Buffer，意思就是 out 变量是一个包含空指针值的非空接口
	// 所以防御性检查 out != nil 的结果依然是 true

	// 动态分配机制依然决定 (*bytes.Buffer).Write 方法会被调用，但是这次的接受者的值是 nil
	// 对于一些如 *os.File 的类型，nil 是一个有效的接收者，但是 *bytes.Buffer 类型不在这些类型中
	// 这个方法会被调用，但它尝试去获取缓冲区时会发生 panic

	// 问题在于尽管一个 nil 的 *bytes.Buffer 指针有实现这个接口的方法，它也不满足这个接口具体行为上的要求
	// 特别是这个调用违反了 (*bytes.Buffer).Write 方法的接收者非空的隐含先决条件，所以将 nil 指针赋给这个接口是错误的
	// 解决方案就是将 main 函数中的变量 buf 的类型改为 io.Writer
	// var buf io.Writer
	// if debug {
	//     buf = new(bytes.Buffer)
	// }
	// f(buf)
}

const debug = true

func f(out io.Writer) {
	// do something...
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}
