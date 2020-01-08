package main

import "os"

// 013、匿名函数 - 捕获迭代变量
func main() {
	// 本节，将介绍 Go 词法作用域的一个陷阱，请务必仔细阅读，弄清除问题发生的原因
	// 即使是经验丰富的程序员也会在这个问题上犯错误

	// 考虑这样一个问题：你被要求首先创建一些目录，然后再将目录删除
	// 在下面的例子中我们用函数值来完成删除操作
	// 为了使代码简单，我们忽略了所有的异常处理
	// 程序 1：
	var rmdirs []func()
	for _, d := range tempDirs() {
		dir := d               // 注意：必要！
		os.MkdirAll(dir, 0755) // 创建父目录
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dir)
		})
	}
	// ...do some work…
	for _, rmdir := range rmdirs {
		rmdir() // clean up
	}

	// 你可能会感到困惑，为什么要在循环体中把变量 d 赋值给一个新的局部变量，而不是像下面代码一样直接使用循环变量 dir
	// 需要注意，下面的代码是错误的
	// 程序 2：
	var rmdirs []func()
	for _, dir := range tempDirs() {
		os.MkdirAll(dir, 0755)
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dir) // 注意：错误的！
		})
	}

	// 问题的原因在于循环变量的作用域
	// 在程序 1 中，for 循环语句引入了新的词法块，循环变量 dir 在这个词法块中被声明
	// 在该循环中生成的所有函数值都共享相同的循环变量
	// 需要注意，函数值中记录的是循环变量的内存地址，而不是循环变量某一时刻的值
	// 以 dir 为例，后续的迭代会不断更新 dir 的值
	// 当删除操作执行时，for 循环已完成，dir 存储的值等于最后一次迭代的值
	// 这意味着，每次对 os.RemoveAll 的调用删除的都是相同的目录
	// 通常为了解决这个问题，我们会引入一个与循环变量同名的局部变量，作为循环变量的副本
	// 比如下面的变量 dir，虽然这看起来很奇怪，但却很有用
	for _, dir := range tempDirs() {
		dir := dir // 声明内部的变量 dir，初始化为循环变量 dir
		// ...
	}

	// 这个问题不仅存在于基于 range 的循环，在下面的例子中，对循环变量 i 的使用也存在同样的问题：
	var rmdirs []func()
	dirs := tempDirs()
	for i := 0; i < len(dirs); i++ {
		os.MkdirAll(dirs[i], 0755)
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dirs[i]) // 注意：错误的！
		})
	}
	// 如果你使用 go 语句（第 8 章）或 defer 语句（5.8 节）会经常遇到此类问题
	// 这不是 go 或 defer 本身导致的，而是因为它们都会等待循环结束后，再执行函数值
}
