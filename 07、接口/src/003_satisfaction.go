package main

import (
	"bytes"
	"io"
	"os"
)

// 003、实现接口的条件
func main() {
	// 一个类型如果拥有一个接口需要的所有方法，那么这个类型就实现了这个接口
	// 例如，*os.File 类型实现了 io.Reader、Writer、Closer 和 ReadWriter 接口，
	// *bytes.Buffer 实现了 Reader、Writer、ReadWriter 接口，但它没有实现 Closer 接口因为它不具有 Close 方法
	// Go 的程序员经常会简要的把一个具体类型描述成一个特定的接口类型
	// 举个例子，*bytes.Buffer 是 io.Writer，*os.File 是 io.ReadWriter

	// 接口的指定规则非常简单，表达一个类型属于某个接口，只要这个类型实现这个接口
	var w io.Writer
	w = os.Stdout         // OK，*os.File 具有 Write 方法
	w = new(bytes.Buffer) // OK，*os.Buffer 具有 Write 方法
	// w = time.Second       // 编译错误，time.Duration 没有 Write 方法

	var rwc io.ReadWriteCloser
	rwc = os.Stdout // OK，*os.File 具有 Read、Write、Close 方法
	// rwc = new(bytes.Buffer) // 编译错误，*bytes.Buffer 没有 Close 方法

	// 这个规则甚至适用于等式右边本身也是一个接口类型
	w = rwc // OK，io.ReadWriteCloser 具有 Write 方法
	// rwc = w // 编译错误，io.Writer 没有 Close 方法

	// 因为 ReadWriter 和 ReadWriteCloser 包含所有的 Writer 方法，
	// 所以任何实现了 ReadWriter 和 ReadWriteCloser 的类型必定也实现了 Writer 接口

	// 我们可以将任意一个值赋给空接口类型
	var any interface{}
	any = true
	any = 12.34
	any = "hello"
	any = map[string]int{"one": 1}
	any = new(bytes.Buffer)
}
