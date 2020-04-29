package main

import (
	"fmt"
	"unsafe"
)

// 002、unsafe.Sizeof、Alignof 和 Offsetof
func main() {
	// unsafe.Sizeof 函数返回操作数在内存中的字节大小，参数可以是任意类型的表达式，但是它并不会对表达式进行求值
	// 一个 Sizeof 函数调用是一个对应 uintptr 类型的常量表达式，
	// 因此返回的结果可以用作数组类型的长度大小，或者用作计算其它的常量

	fmt.Println(unsafe.Sizeof(float64(0))) // 8

	// Sizeof 函数返回的大小只包括数据结构中固定的部分，例如字符串对应结构体中的指针和字符串长度部分，
	// 但是并不包含指针指向的字符串的内容
	// Go 语言中非聚合类型通常有一个固定的大小，尽管在不同工具链下生成的实际大小可能会有所不同
	// 考虑到可移植性，引用类型或包含引用类型的大小在 32 位平台上是 4 个字节，在 64 位平台上是 8 个字节

	// 计算机在加载和保存数据时，如果内存地址合理地对齐将会更有效率
	// 例如 2 字节大小的 int16 类型的变量地址应该是偶数，一个 4 字节大小的 rune 类型变量的地址应该是 4 的倍数，
	// 一个 8 字节大小的 float64、uint64 或 64-bit 指针类型变量的地址应该是 8 个字节对齐的
	// 但是对于再大的地址对齐倍数则是不需要的，即使是 complex128 等较大的数据类型最多也只是 8 字节对齐

	// 由于地址对齐这个因素，一个聚合类型（struct 或 array）的大小至少是所有字段或元素大小的总和或者更大，因为可能存在内存空洞
	// 内存空洞是编译器自动添加的、没有被使用的内存空间，用于保证后面每个字段或元素的地址相对于结构或数组的开始地址能够合理地对齐
	// （译注：内存空洞可能会存在一些随机数据，可能会对用 unsafe 包直接操作内存的处理产生影响）

	// | 类型                            | 大小                          |
	// | ----------------------------- | --------------------------- |
	// | bool                          | 1 个字节                       |
	// | intN, uintN, floatN, complexN | N/8 个字节（例如 float64 是 8 个字节） |
	// | int, uint, uintptr            | 1 个机器字                      |
	// | *T                            | 1 个机器字                      |
	// | string                        | 2 个机器字(data, len)           |
	// | []T                           | 3 个机器字(data, len, cap)      |
	// | map                           | 1 个机器字                      |
	// | func                          | 1 个机器字                      |
	// | chan                          | 1 个机器字                      |
	// | interface                     | 2 个机器字(type, value)         |

	// Go 语言的规范并没有要求一个字段的声明顺序和内存中的顺序是一致的，
	// 所以理论上一个编译器可以随意地重新排列每个字段的内存位置，虽然在写作本书的时候编译器还没有这么做
	// 下面三个结构体虽然有着相同的字段，但是第一种写法比另外两个要多 50% 的内存
	//                                    64bit   32bit
	// struct { bool, float64, int16 } // 3words  4words
	// struct { float64, int16, bool } // 2words  3words
	// struct { bool, int16, float64 } // 2words  3words

	// 关于内存地址对齐算法的细节超出了本书的范围，也不是每一个 struct 都需要担心这个问题，不过有效的包装可以使数据结构更加紧凑
	// （译注：未来 Go 语言的编译器应该会默认优化 struct 的顺序，当然应该也能够指定具体的内存布局，
	//   相同讨论请参考：https://github.com/golang/go/issues/10014），
	// 内存使用率和性能都可能会受益

	// unsafe.Alignof 函数返回对应参数的类型需要对齐的倍数
	// 和 Sizeof 类似，Alignof 也是返回一个常量表达式，对应一个常量
	// 通常情况下布尔和数字类型需要对齐到它们本身的大小（最多 8 个字节），其它的类型对齐到机器字大小

	// unsafe.Offsetof 函数的参数必须是一个字段 x.f，然后返回 f 字段相对于 x 起始地址的偏移量，包括可能的空洞
	// 图 13.1 显示了一个结构体变量 x 以及其在 32 位和 64 位机器上的典型的内存，灰色区域（#）是空洞
	//
	// var x struct {
	// 	   a bool
	// 	   b int16
	// 	   c []int
	// }
	//
	// 下面显示了对 x 和它的三个字段调用 unsafe 包相关函数的计算结果
	//
	//           |
	//  a ### b  |  a ### b ############
	//  c(data)  |        c(data)
	//  c(len)   |        c(len)
	//  c(cap)   |        c(cap)
	//           |
	//  32 bit   |        64 bit
	//  图 13.1、Holes in a struct

	// 32 位系统：
	// Sizeof(x)   = 16  Alignof(x)   = 4
	// Sizeof(x.a) = 1   Alignof(x.a) = 1  Offsetof(x.a) = 0
	// Sizeof(x.b) = 2   Alignof(x.b) = 2  Offsetof(x.b) = 2
	// Sizeof(x.c) = 12  Alignof(x.c) = 4  Offsetof(x.c) = 4

	// 64 位系统：
	// Sizeof(x)   = 32  Alignof(x)   = 8
	// Sizeof(x.a) = 1   Alignof(x.a) = 1  Offsetof(x.a) = 0
	// Sizeof(x.b) = 2   Alignof(x.b) = 2  Offsetof(x.b) = 2
	// Sizeof(x.c) = 24  Alignof(x.c) = 8  Offsetof(x.c) = 8

	// 虽然这几个函数在不安全的 unsafe 包，但是这几个函数调用并不是真的不安全，
	// 特别在需要优化内存空间时，它们返回的结果对于理解原生的内存布局很有帮助
}
