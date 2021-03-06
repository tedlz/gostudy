package main

// 001、引言
func main() {

}

// Go 语言提供了一种机制，能够在运行时更新变量和检查它们的值、调用它们的方法和它们支持的内在操作，
// 而不需要在编译时就知道这些变量的具体类型，这种机制被称为反射
// 反射也可以让我们将类型本身作为第一类的值类型处理

// 在本章，我们将探讨 Go 语言的反射特性，看看它可以给语言增加哪些表达力，以及在两个至关重要的 API 中是如何使用反射机制的：
// 一个是 fmt 包提供的字符串格式功能，另一个是类似 encoding/json 和 encoding/xml 提供的针对特定协议的编解码功能

// 对于我们在 4.6 节中看到过的 text/template 和 html/template 包，它们的实现也是依赖反射技术的
// 然后，反射是一个复杂的内省技术，不应该随意使用
// 因此，尽管上面这些包内部都是用反射技术实现的，但是它们自己的 API 都没有公开反射相关的接口
