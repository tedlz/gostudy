package main

import (
	"fmt"
	"reflect"
)

// 006、通过 reflect.Value 修改值
func main() {
	// 到目前为止，反射还只是程序中变量的另一种读取方式
	// 然而，在本节中，我们将重点讨论如何通过反射机制来修改变量

	// 回想一下，Go 语言中类似 x、x.f[1] 和 *p 形式的表达式都可以表示变量，但是其它如 x + 1 和 f(2) 则不是变量
	// 一个变量就是一个可寻址的内存空间，里面存储了一个值，并且存储的值可以通过内存地址来更新

	// 对于 reflect.Values 也有类似的区别
	// 有一些 reflect.Values 是可取地址的，有一些不可以
	// 考虑以下的声明语句：
	x := 2                   // value  type  variable?
	a := reflect.ValueOf(2)  // 2      int   no
	b := reflect.ValueOf(x)  // 2      int   no
	c := reflect.ValueOf(&x) // &x     *int  no
	d := c.Elem()            // 2      int   yes (x)

	// 其中 a 对应的变量不可取地址，因为 a 中的值仅仅是整数 2 的拷贝副本
	// b 中的值也同样不可取地址，c 中的值还是不可取地址，
	// 但是对于 d，它是 c 的解引用方式生成的，指向另一个变量，因此是可取地址的
	// 我们可以通过调用 reflect.ValueOf(&x).Elem()，来获取任意变量 x 对应的可取地址的 Value

	// 我们可以通过调用 reflect.Value 的 CanAddr 方法来判断其是否可以被取地址
	fmt.Println(a.CanAddr()) // false
	fmt.Println(b.CanAddr()) // false
	fmt.Println(c.CanAddr()) // false
	fmt.Println(d.CanAddr()) // true

	// 每当我们通过指针间接地获取的 reflect.Value 都是可取地址的，即使开始的是一个不可取地址的 Value
	// 在反射机制中，所有关于是否支持取地址的规则都是类似的
	// 例如，slice 的索引表达式 e[i] 将隐式的包含一个指针，它就是可取地址的，即使是开始的 e 表达式不支持也没关系
	// 以此类推，reflect.ValueOf(e).Index(i) 对于值也是可取地址的，即使 reflect.ValueOf(e) 不支持也没关系

	// 要从变量对应的可取地址的 reflect.Value 来访问变量需要三个步骤
	// 第一步是调用 Addr() 方法，它返回一个 Value，里面保存了指向变量的指针
	// 最后，如果我们知道变量的类型，我们可以使用类型的断言机制将得到的 interface{} 类型的接口强制转为普通的类型指针
	// 这样我们就可以通过这个普通指针来更新变量了：
	h := 2
	i := reflect.ValueOf(&h).Elem()   // i 引用了变量 h
	ph := i.Addr().Interface().(*int) // ph := &h
	*ph = 3                           // h = 3
	fmt.Println(h)                    // 3

	// 或者不使用指针，而是通过调取可取地址的 reflect.Value 的 reflect.Value.Set 方法来更新对应的值
	i.Set(reflect.ValueOf(4))
	fmt.Println(h) // 4

	// Set 方法将在运行时执行和编译时进行类似的可赋值性约束的检查
	// 以上代码，变量和值都是 int 类型，但是如果变量是 int64 类型，那么程序将抛出一个 panic 异常，
	// 所以关键是要确保值可以分配给变量的类型
	// i.Set(reflect.ValueOf(int64(5))) // panic: int64 is not assignable to int

	// 同样，对于一个不可取地址的 reflect.Value 调用 Set 方法也会导致 panic 异常
	// x := 2
	// b := reflect.ValueOf(x)
	// b.Set(reflect.ValueOf(3)) // panic: Set using unaddressable value

	// 这里有很多用于基本类型的 Set 方法：SetInt、SetUint、SetString、SetFloat 等
	j := 2
	k := reflect.ValueOf(&j).Elem()
	k.SetInt(3)
	fmt.Println(j) // 3

	// 从某种程度上说，这些 Set 方法总是尽可能的完成任务
	// 以 SetInt 为例，只要变量是某种类型的有符号的整数就可以工作，
	// 即使是一些命名的类型、甚至只要底层数据类型是有符号的整数就可以，
	// 如果值太大，就会被悄悄地截断以适应
	// 但需要谨慎的是：对于一个引用 interface{} 类型的 reflect.Value 调用 SetInt 会导致 panic 异常，
	// 即使那个 interface{} 变量是整数类型也不行
	//
	// o := 1
	// ro := reflect.ValueOf(&o).Elem()
	// ro.SetInt(2)                     // OK, o = 2
	// ro.Set(reflect.ValueOf(3))       // OK, o = 3
	// ro.SetString("hello")            // panic: string is not assignable to int
	// ro.Set(reflect.ValueOf("hello")) // panic: string is not assignable to int

	// var p interface{}
	// rp := reflect.ValueOf(&p).Elem()
	// rp.SetInt(2)                     // panic: SetInt called on interface Value
	// rp.Set(reflect.ValueOf(3))       // OK, p = int(3)
	// rp.SetString("hello")            // panic: SetString called on interface Value
	// rp.Set(reflect.ValueOf("hello")) // OK, p = "hello"
	//

	// 当我们用 Display 显示 os.Stdout 结构时，我们发现反射可以越过 Go 语言导出规则的限制读取结构体中未导出的成员
	// 比如在类 Unix 的系统上 os.File 结构体中的 fd int 成员
	// 然而，利用反射机制并不能修改这些未导出的成员：
	// stdout := reflect.ValueOf(os.Stdout).Elem() // *os.Stdout, an os.File var
	// fmt.Println(stdout.Type())                  // os.File
	// fd := stdout.FieldByName("fd")
	// fmt.Println(fd.Int()) // 1
	// fd.SetInt(2)          // panic: unexported field

	// 一个可取地址的 reflect.Value 会记录一个结构体成员是否是未导出成员，如果是的话则拒绝修改操作
	// 因此，CanAddr 方法并不能正确反映一个变量是否可以被修改
	// 另一个相关方法 CanSet 用于检查对应的 reflect.Value 是否可取地址且可被修改：
	// fmt.Println(fd.CanAddr(), fd.CanSet()) // true false
}
