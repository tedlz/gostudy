package main

// 015、类型断言
func main() {
	// 类型断言是一个使用在接口值上的操作，语法上它看起来像 x.(T) 被称为断言类型（x 表示接口类型，T 表示类型）
	// 一个类型断言检查它操作对象的动态类型是否和断言的类型匹配

	// 这里有两种可能
	// 第一种，如果断言的类型 T 是一个具体类型，然后类型断言检查 x 的动态类型是否和 T 相同
	// 如果检查成功了，类型断言的结果是 x 的动态值，当然它的类型是 T
	// 换句话说，具体类型的类型断言，从它的操作对象中获得具体的值
	// 如果检查失败，接下来的操作会抛出 panic，例如：
	// var w io.Writer
	// w = os.Stdout
	// f := w.(*os.File)      // success: f == os.Stdout
	// c := w.(*bytes.Buffer) // panic: interface holds *os.File, not *bytes.Buffer

	// 第二种，如果相反断言的类型 T 是一个接口类型，然后类型断言检查是否 x 的动态类型满足 T
	// 如果这个检查成功了，动态值没有获取到；这个结果仍然是一个有相同类型和值部分的接口值，但是结果有类型 T
	// 换句话说，对一个接口类型的类型断言改变了类型的表述方式，改变了可以获取的方法集合（通常更大），
	// 但它保护了接口值内部的动态类型和值的部分

	// 在下面的第一个类型断言后，w 和 rw 都持有 os.Stdout，因此它们每个有一个动态类型 *os.File
	// 但是变量 w 是一个 io.Writer 类型只对外公开出文件的 Write 方法，然而 rw 变量也只公开它的 Read 方法
	// var w io.Writer
	// w = os.Stdout
	// rw := w.(io.ReadWriter) // success: *os.File has both Read and Write
	// w = new(ByteCounter)
	// rw = w.(io.ReadWriter) // panic: *ByteCounter has no Read method

	// 如果断言操作的对象是一个 nil 接口值，那么无论被断言的类型是什么，这个类型断言都会失败
	// 我们几乎不需要对一个更少限制性的接口类型（更少的方法集合）做断言，因为它表现的就像赋值操作一样
	// 除了对于 nil 接口值的情况
	// w = rw             // io.ReadWriter is assignable to io.Writer
	// w = rw.(io.Writer) // fails only if rw == nil

	// 经常地我们对一个接口值的动态类型是不确定的，并且我们更愿意去检验它是否是一些特定的类型
	// 如果类型断言出现在一个预期有两个结果的赋值操作中，例如如下的定义，这个操作不会在失败的时候发生 panic，
	// 但是代替的返回一个额外的第二个结果，这个结果是一个标识成功的布尔值：
	// var w io.Writer = os.Stdout
	// f, ok := w.(*os.File)      // success:  ok, f == os.Stdout
	// b, ok := w.(*bytes.Buffer) // failure, !ok, b == nil
	// 第二个结果常规的赋值给一个命名为 ok 的变量，如果这个操作失败了，
	// 那么 ok 就是 false 值，第一个结果等于被断言类型的零值，在这个例子中就是一个 nil 的 *bytes.Buffer 类型

	// 这个 ok 结果经常立即用于决定程序下面做什么。if 语句的扩展格式让这个变的很简洁：
	// if f, ok := w.(*os.File); ok {
	//     ......
	// }

	// 当类型断言的操作对象是一个变量，你有时会看见原来的变量名重用而不是声明一个新的本地变量
	// 这个重用的变量会覆盖原来的值，如下面这样：
	// if w, ok := w.(*os.File); ok {
	// 	   ......
	// }
}
