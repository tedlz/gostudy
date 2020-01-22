package main

// 016、基于类型断言区别错误类型
func main() {
	// 思考在 os 包中文件操作返回的错误集合
	// I/O 可以因为任何数量的原因失败，但是有三种经常的错误必须进行不同的处理：
	// 文件已经存在（创建文件时）、找不到文件（读取文件时）、和权限拒绝
	// os 包中提供了这三个帮助函数来对给定的错误值表示的失败进行分类：
	//
	// package os
	// func IsExist(err error) bool
	// func IsNotExist(err error) bool
	// func IsPermission(err error) bool
	//
	// 对这些判断的一个缺乏经验的实现可能会去检查错误消息是否包含了特定的子字符串：
	//
	// func IsNotExist(err error) bool {
	// 	   警告：不稳定
	// 	   return strings.Contains(err.Error(), "file does not exist")
	// }
	//
	// 但是处理 I/O 错误的逻辑可能一个和另一个平台非常的不同，所以这种方案并不健壮
	// 并且对相同的失败可能会报出各种不同的错误消息
	// 在测试的过程中，通过检查错误消息的子字符串来保证特定的函数以期望的方式失败是非常有用的，但对于线上的代码是不够的

	// 一个更可靠的方式是使用专门的类型来描述结构化的错误
	// os 包中定义了一个 PathError 类型来描述在文件路径操作中涉及到的失败，像 Open 和 Delete 操作
	// 并且定义了一个叫 LinkError 的变体来描述涉及到两个文件路径的操作，像 Symlink 和 Rename
	// 下面是 os.PathError：
	//
	// package os
	// type PathError struct { // PathError 记录错误以及导致该错误的操作和文件路径
	//     Op string
	//     Path string
	//     Err error
	// }
	// func (e *PathError) Error() string {
	// 	   return e.Op + " " + e.Path + ": " + e.Err.Error()
	// }
	//
	// 大多数调用方都不知道 PathError 并且通过调用错误本身的 Error 方法来统一处理所有的错误
	// 尽管 PathError 的 Error 方法简单的把这些字段连起来生成错误消息，PathError 的结构保护了内部的错误组件
	// 调用方需要使用类型断言来检测错误的具体类型以便将一种失败和另一种区分开；具体的类型比字符串可以提供更多的细节
	//
	// _, err := os.Open("/no/such/file")
	// fmt.Println(err) // open /no/such/file: No such file or directory
	// fmt.Printf("%#v\n", err)
	// 输出：
	// &os.PathError{Op: "open", Path: "/no/such/file", Err: 0x2}
	//
	// 这就是三个帮助函数是怎么工作的
	// 例如下面展示的 IsNotExist，它会报出是否一个错误和 syscall.ENOENT 或者和有名的错误 os.ErrNotExist 相等
	// 或者是一个 PathError，它内部的错误是 syscall.ENOENT 和 os.ErrNotExist 其中之一
	//
	// import (
	//     "errors"
	//     "syscall"
	// )
	// var ErrNotExist = errors.New("file does not exist")
	// func IsNotExist(err error) bool {
	// 	   if pe, ok := err.(*PathError); ok {
	// 		   err = pe.Err
	// 	   }
	// 	   return err == syscall.ENOENT || err == ErrNotExist
	// }
	//
	// IsNotExist 返回一个布尔值，该布尔值指示是否已知该错误，以报告文件或目录不存在
	// ErrNotExist 以及一些系统调用错误都可以满足要求
	// 下面是它的实际使用：
	//
	// _, err := os.Open("/no/such/file")
	// fmt.Println(os.IsNotExist(err)) // true
	//
	// 如果错误消息结合成一个更大的字符串，当然 PathError 的结构就不再为人所知，例如通过一个对 fmt.Errorf 函数的调用
	// 区别错误通常必须在操作失败后，错误传回调用者前进行
}
