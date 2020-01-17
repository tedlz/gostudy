package main

// 002、接口类型
func main() {
	// 接口类型具体描述了一系列方法的集合，一个实现了这些方法的具体类型是这个接口类型的实例
	// io.Writer 类型是用的最广泛的接口之一，因为它提供了所有类型写入 bytes 的抽象
	// 包括文件类型、内存缓冲区、网络连接、HTTP 客户端、压缩工具、哈希等等，io 包中还定义了很多其它有用的接口类型
	// Reader 代表可以任意读取 bytes 的类型，Closer 是任意可以关闭的值，例如一个文件或是网络连接：

	// package io
	// type Reader interface {
	// 	   Read(p []byte) (n int, err error)
	// }
	// type Closer interface {
	// 	   Close() error
	// }

	// 有些接口类型通过组合已有接口来定义
	// type ReadWriter interface {
	// 	   Reader
	// 	   Writer
	// }
	// type ReadWriterCloser interface {
	// 	   Reader
	// 	   Writer
	// 	   Closer
	// }

	// 上面用到的语法和结构内嵌相似，我们可以用这种方式以一个简写命名另一个接口，而不用声明它所有的方法，这种方式称为接口内嵌
	// 尽管略失简洁，我们可以像下面这样，不使用内嵌来声明 io.Writer 接口
	// type ReadWriter interface {
	// 	   Read(p []byte) (n int, err error)
	// 	   Write(p []byte) (n int, err error)
	// }
	// 或者使用这种混合的风格
	// type ReadWriter interface {
	// 	   Read(p []byte) (n int, err error)
	// 	   Write
	// }

	// 上面三种定义方式都是一样的效果，方法的顺序变化也没有影响，唯一重要的就是这个集合里面的方法
}
