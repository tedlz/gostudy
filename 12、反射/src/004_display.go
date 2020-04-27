package main

import (
	"fmt"
	"gostudy/11、测试/files/eval"
	"gostudy/12、反射/files/display"
	"os"
	"reflect"
)

// 004、Display，一个递归的值打印器
func main() {
	// 接下来，让我们看看如何改善聚合数据类型的显示
	// 我们并不想完全克隆一个 fmt.Sprint 函数，我们只是构建一个用于调试用的 Display 函数：
	// 给定任意一个复杂类型 x，打印这个值对应的完整结构，同时标记每个元素的发现路径
	// 让我们从一个例子开始

	e, _ := eval.Parse("sqrt(A / pi)")
	display.Display("e", e)

	// 在上面的调用中，传入 Display 函数的参数是在 7.9 节一个表达式求值函数返回的语法树
	// Display 函数的输出如下：

	// Display e (eval.call)
	// e.fn = "sqrt"
	// e.args[0].type = eval.binary
	// e.args[0].value.op = 47
	// e.args[0].value.x.type = eval.Var
	// e.args[0].value.x.value = "A"
	// e.args[0].value.y.type = eval.Var
	// e.args[0].value.y.value = "pi"

	// 你应该尽量避免在一个包的 API 中暴露涉及反射的接口。我们将定义一个未导出的 display 函数用于递归处理工作，
	// 导出的是 display 函数，它只是 display 函数简单的包装以接受 interface{} 类型的参数
	// （见 files/display/display.go 的 Display 函数）

	// 在 display 函数中，我们使用了前面定义的打印基础类型 —— 基本类型、函数和 chan 等 —— 元素值的 formatAtom 函数，
	// 但是我们会使用 reflect.Value 方法来递归显示复杂类型的每一个成员
	// 在递归下降过程中，path 字符串，从最开始传入的起始值（这里是 "e"），
	// 将逐步增长来表示是如何到达当前值（例如 e.args[0].value）的

	// 因为我们不再模拟 fmt.Sprint 函数，我们将直接使用 fmt 包来简化我们的例子实现
	// （见 files/display/display.go 的 display 函数）

	// 让我们针对不同类型分别讨论：
	// * Slice 和 Array：
	// 两种的处理逻辑是一样的。Len 方法返回 slice 或 array 值中的元素个数，
	// Index(i) 活动索引 i 对应的元素，返回的也是一个 reflect.Value；
	// 如果索引 i 超出范围的话将导致 panic 异常，这与 array 或 slice 类型内建的 len(a) 和 a[i] 操作类似
	// display 针对序列中的每个元素递归调用自身处理，我们通过在递归处理时向 path 附加 [i] 来表示访问路径
	// 虽然 reflect.Value 类型带有很多方法，但是只有少数的方法能对任意值都安全调用
	// 例如，Index 方法只能对 slice、array 或 string 类型的值调用，如果对其它类型的值调用则会导致 panic 异常
	// * Struct：
	// NumField 方法报告结构体中成员的数量，Field(i) 以 reflect.Value 类型返回第 i 个成员的值
	// 成员列表也包括通过匿名字段提升上来的成员
	// 为了在 path 添加 ".f" 来表示成员路径，我们必须获得结构体对应的 reflect.Type 类型信息，
	// 然后访问结构体第 i 个成员的名字
	// * Map：
	// MapKeys 方法返回一个 reflect.Value 类型的 slice，每一个元素对应 map 的一个 key
	// 和往常一样，遍历 map 时顺序是随机的
	// MapIndex(key) 返回 map 中 key 对应的 value
	// 我们向 path 添加 "[key]" 来表示访问路径
	// （我们这里有一个未完成的工作，其实 map 中 key 的类型并不局限于 formatAtom 能完美处理的类型，
	//   array、struct、interface 都可以作为 map 的 key，
	//   针对这种类型，完善 key 的显示信息是练习 12.1 的任务）
	// * Pointer：
	// Elem 方法返回指针指向的变量，依然是 reflect.Value 类型
	// 即使指针是 nil，这个操作也是安全的，在这种情况下指针是 invalid 类型，
	// 但是我们可以用 IsNil 方法来显式测试一个空指针，这样我们可以打印更合适的信息
	// 我们在 path 前面加 "*"，并用括号包含避免歧义
	// * Interface：
	// 再一次，我们使用 IsNil 方法来测试接口是否是 nil，
	// 如果不是，我们可以调用 v.Elem() 来获取接口对应的动态值，并打印对应的类型和值

	// 现在我们的 Display 函数总算完工了，让我们看看它的表现吧
	// 下面的 Movie 类型是在 4.5 节的 Movie 类型上演变而来的
	// （见 main 函数下面的代码）
	// 让我们声明一个该类型的变量，然后看看 Display 函数如何显示它：
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}
	fmt.Println()
	display.Display("strangelove", strangelove)

	// Display("strangelove", strangelove) 调用将显示（strangelove电影对应的中文名是《奇爱博士》）：

	// Display strangelove (main.Movie):
	// strangelove.Title = "Dr. Strangelove"
	// strangelove.Subtitle = "How I Learned to Stop Worrying and Love the Bomb"
	// strangelove.Year = 1964
	// strangelove.Color = false
	// strangelove.Actor["Dr. Strangelove"] = "Peter Sellers"
	// strangelove.Actor["Grp. Capt. Lionel Mandrake"] = "Peter Sellers"
	// strangelove.Actor["Pres. Merkin Muffley"] = "Peter Sellers"
	// strangelove.Actor["Gen. Buck Turgidson"] = "George C. Scott"
	// strangelove.Actor["Brig. Gen. Jack D. Ripper"] = "Sterling Hayden"
	// strangelove.Actor["Maj. T.J. \"King\" Kong"] = "Slim Pickens"
	// strangelove.Oscars[0] = "Best Actor (Nomin.)"
	// strangelove.Oscars[1] = "Best Adapted Screenplay (Nomin.)"
	// strangelove.Oscars[2] = "Best Director (Nomin.)"
	// strangelove.Oscars[3] = "Best Picture (Nomin.)"
	// strangelove.Sequel = nil

	// 我们也可以使用 Display 函数来显示标准库中类型的内部结构，例如 *os.File 类型：
	fmt.Println()
	display.Display("os.Stderr", os.Stderr)
	// 输出：
	// Display os.Stderr (*os.File):
	// (*(*os.Stderr).file).pfd.fdmu.state = 0
	// (*(*os.Stderr).file).pfd.fdmu.rsema = 0
	// (*(*os.Stderr).file).pfd.fdmu.wsema = 0
	// (*(*os.Stderr).file).pfd.Sysfd = 2
	// (*(*os.Stderr).file).pfd.pd.runtimeCtx = 0
	// (*(*os.Stderr).file).pfd.iovecs = nil
	// (*(*os.Stderr).file).pfd.csema = 0
	// (*(*os.Stderr).file).pfd.isBlocking = 1
	// (*(*os.Stderr).file).pfd.IsStream = true
	// (*(*os.Stderr).file).pfd.ZeroReadIsEOF = true
	// (*(*os.Stderr).file).pfd.isFile = true
	// (*(*os.Stderr).file).name = "/dev/stderr"
	// (*(*os.Stderr).file).dirinfo = nil
	// (*(*os.Stderr).file).nonblock = false
	// (*(*os.Stderr).file).stdoutOrErr = true
	// (*(*os.Stderr).file).appendMode = false

	// 可以看出，反射能够访问到结构体中未导出的成员
	// 需要当心的是这个例子的输出在不同操作系统上可能是不同的，并且随着标准库的发展也可能导致结果不同
	// （这也是将这些成员定义为私有成员的原因之一）

	// 我们甚至可以用 Display 函数来显示 reflect.Value 的内部构造（在这里设置为 *os.File 的类型描述体）

	fmt.Println()
	display.Display("rV", reflect.ValueOf(os.Stderr))
	// Display("rV", reflect.ValueOf(os.Stderr)) 调用的输出如下，当然不同环境得到的结果可能有差异：

	// Display rV (reflect.Value):
	// (*rV.typ).size = 8
	// (*rV.typ).ptrdata = 8
	// (*rV.typ).hash = 871609668
	// (*rV.typ).tflag = 9
	// (*rV.typ).align = 8
	// (*rV.typ).fieldAlign = 8
	// (*rV.typ).kind = 54
	// (*rV.typ).equal = func(unsafe.Pointer, unsafe.Pointer) bool 0x403220
	// (*(*rV.typ).gcdata) = 1
	// (*rV.typ).str = 7676
	// (*rV.typ).ptrToThis = 0
	// rV.ptr = unsafe.Pointer value
	// rV.flag = 22

	fmt.Println()
	// 观察下面两个例子的区别：
	var i interface{} = 3
	display.Display("i", i)
	// Display i (int):
	// i = 3

	fmt.Println()
	display.Display("&i", &i)
	// Display &i (*interface {}):
	// (*&i).type = int
	// (*&i).value = 3

	// 在第一个例子中，Display 函数调用 reflect.ValueOf(i)，它返回一个 int 类型的值
	// 正如我们在 12.2 节中提到的，reflect.ValueOf 总是返回一个具体类型的 Value，因为它是从一个接口值提取的内容
	// 在第二个例子中，Display 函数调用的是 reflect.ValueOf(&i)，它返回一个指向 i 的指针，对应 Ptr 类型
	// 在 switch 的 Ptr 分支中，对这个值调用 Elem 方法，返回一个 Value 来表示变量 i 本身，对应 interface 类型
	// 像这样一个间接获得的 Value，可能代表任意类型的值，包括接口类型
	// display 函数递归调用自身，这次它分别打印了这个接口的动态类型和值

	// 对于目前的实现，如果遇到对象图中含有回环，Display 将会陷入死循环，例如下面这个首尾相连的链表：
	// a struct that points to itself
	// type Cycle struct{ Value int; Tail *Cycle }
	// var c Cycle
	// c = Cycle{42, &c}
	// Display("c", c)

	// Display 会永远不停地进行深度递归打印：
	// Display c (display.Cycle):
	// c.Value = 42
	// (*c.Tail).Value = 42
	// (*(*c.Tail).Tail).Value = 42
	// (*(*(*c.Tail).Tail).Tail).Value = 42
	// ...ad infinitum...

	// 许多 Go 语言程序都包含了一些循环的数据
	// 让 Display 支持这些带环的数据结构需要些技巧，需要额外记录迄今访问的路径，相应会带来成本
	// 通用的解决方案是采用 unsafe 的语言特性，我们将在 13.3 节中看到具体的解决方案

	// 带环的数据结构很少会对 fmt.Sprint 函数造成问题，因为它很少尝试打印完整的数据结构
	// 例如，当它遇到一个指针的时候，它只是简单地打印指针的数字值
	// 在打印包含自身的 slice 或 map 时可能卡住，但是这种情况很罕见，不值得付出为了处理回环所需的开销
}

// Movie *
type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}
