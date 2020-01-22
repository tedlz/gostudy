package main

// 017、通过类型断言询问行为
func main() {
	// 下面这段逻辑和 net/http 包中 web 服务器负责写入 HTTP 头字段（例如 Content-Type: text/html）的部分相似
	// io.Writer 接口类型的变量 w 代表 HTTP 响应，写入它的字节最终被发送到某个人的 web 浏览器上
	//
	// func writeHeader(w io.Writer, contentType string) error {
	// 	   if _, err := w.Write([]byte("Content-Type: ")); err != nil {
	// 		   return err
	// 	   }
	// 	   if _, err := w.Write([]byte(contentType)); err != nil {
	// 		   return err
	// 	   }
	// 	   ......
	// }
	//
	// 因为 Write 方法需要传入一个 byte 切片，而我们希望写入的值是一个字符串，所以我们需要使用 []byte(...) 进行转换
	// 这个转换分配内存并且做一个拷贝，但是这个拷贝在转换后几乎立马就被丢弃掉
	// 让我们假装这是一个 web 服务器的核心部分并且我们的性能分析表示这个内存分配使服务器的速度变慢
	// 这里我们可以避免掉内存分配么？

	// 这个 io.Writer 接口告诉我们关于 w 持有的具体类型的唯一东西：就是可以向它写入字节切片
	// 如果我们回顾 net/http 包中的内幕，我们知道在这个程序中的 w 变量持有的动态类型，
	// 也有一个允许字符串高效写入的 WriteString 方法；这个方法会避免取分配一个临时的拷贝
	// （这可能像在黑夜中射击一样，但是许多满足 io.Writer 接口的重要类型同时也有 WriteString 方法，
	//  包括 *bytes.Buffer、*os.File 和 *bufio.Writer）

	// 我们不能对任意 io.Writer 类型的变量 w，假设它也拥有 WriteString 方法
	// 但是我们可以定义一个只有这个方法的新接口并且使用类型断言来检测是否 w 的动态类型满足这个新接口

	// writeString 将 s 写入 w
	// 如果 w 具有 writeString 方法，则代替 w.Write 调用它
	// func writeString(w io.Writer, s string) (n int, err error) {
	// 	   type stringWriter interface {
	// 		   WriteString(string) (n int, err error)
	// 	   }
	// 	   if sw, ok := w.(stringWriter); ok {
	// 		   return sw.WriteString(s) // 避免拷贝
	// 	   }
	// 	   return w.Write([]byte(s)) // 分配临时副本
	// }
	// func writeHeader(w io.Writer, contentType string) error {
	// 	   if _, err := writeString(w, "Content-Type: "); err != nil {
	// 		   return err
	// 	   }
	// 	   if _, err := writeString(w, contentType); err != nil {
	// 		   return err
	// 	   }
	// 	   ......
	// }

	// 为了避免重复定义，我们将这个检查移入到一个实用工具函数 writeString 中，
	// 但是它太有用了以致标准库将它作为 io.WriteString 函数提供
	// 这是向一个 io.Writer 接口写入字符串的推荐方法

	// 这个例子的神奇之处，在于没有定义了 WriteString 方法的标准接口和没有指定它是一个需要行为的标准接口
	// 而且一个具体类型只会通过它的方法决定它是否满足 stringWriter 接口，而不是任何它和这个接口类型表明的关系
	// 它的意思就是上面的技术依赖于一个假设；
	// 这个假设就是，如果一个类型满足下面的这个接口，然后 WriteString(s) 方法就必须和 Write([]byte(s)) 有相同的效果：
	//
	// interface {
	//     io.Writer
	//     WriteString(s string) (n int, err error)
	// }
	//
	// 尽管 io.WriteString 记录了它的假设，但是调用它的函数极少有可能会去记录它们也做了同样的假设
	// 定义一个特定类型的方法隐式的获取了对特定行为的协约
	// 对于 Go 语言的新手，特别是那些来自有强类型语言使用背景的新手，可能会发现它缺乏显式的意图令人感到混乱
	// 但是在实战的过程中这几乎不是一个问题
	// 除了空接口 interface{}，接口类型很少意外巧合的被实现

	// 上面的 writeString 函数使用一个类型断言来知道一个普遍接口类型的值是否满足一个更加具体的接口类型
	// 并且如果满足，它会使用这个更具体接口的行为
	// 这个技术可以被很好的使用，无论这个被询问的接口是一个标准的如 io.ReadWriter 或者用户定义的如 stringWriter

	// 这也是 fmt.Fprintf 函数怎么从其它所有值中区分满足 error 或者 fmt.Stringer 接口的值
	// 在 fmt.Fprintf 内部，有一个将单个操作对象转换成一个字符串的步骤，像下面这样：
	//
	// package fmt
	// func formatOneValue(x interface{}) string {
	// 	   if err, ok := x.(error); ok {
	// 		   return err.Error()
	// 	   }
	// 	   if str, ok := x.(Stringer); ok {
	// 		   return str.String()
	// 	   }
	// 	   ...... all other types
	// }
	//
	// 如果 x 满足这两个接口类型中的一个，具体满足的接口决定对值的格式化方式
	// 如果都不满足，默认的 case 或多或少会统一的使用反射来处理所有的其它类型，我们可以在第 12 章知道具体是怎么实现的

	// 再一次的，它假设任何有 String 方法的类型满足 fmt.Stringer 中约定的行为，这个行为会返回一个适合打印的字符串
}
