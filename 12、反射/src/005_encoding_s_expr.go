package main

// 005、示例：编码 S 表达式
func main() {
	// Display 是一个用于显示结构化数据的调试工具，但是它并不能将任意的 Go 语言对象编码为通用消息，然后用于进程间的通信

	// 正如我们在 4.5 节中看到的，Go 语言的标准库支持了包括 JSON、XML 和 ASN.1 等多种编码格式
	// 还有另一种依然被广泛使用的格式是 S 表达式，采用 Lisp 语言的语法
	// 但是和其它编码格式不同的是，Go 语言自带的标准库并不支持 S 表达式，主要是它没有一个公认的标准规范

	// 在本节中，我们将定义一个包，用于将任意的 Go 语言对象编码为 S 表达式格式，它支持以下结构：
	// 42        integer
	// "hello"   string （带有 Go 风格的引号）
	// foo       symbol （未用引号括起来的名字）
	// (1, 2, 3) list   （括号包起来的 0 个或多个元素）

	// 布尔型习惯上用 t 符号表示 true，空列表或 nil 符号表示 false，但是为了简单起见，我们暂时忽略布尔类型
	// 同时忽略的还有 chan 管道和函数，因为通过反射并无法知道它们的确切状态
	// 我们忽略的还有浮点数、复数和 interface，支持它们是练习 12.3 的任务

	// 我们将 Go 语言的类型编码为 S 表达式的方法如下
	// 整数和字符串以显而易见的方式编码，空值编码为 nil 符号，数组和 slice 被编码为列表
	// 结构体被编码为成员对象的列表，每个成员对象对应一个有两个元素的子列表，子列表的第一个元素是成员的名字，第二个元素是成员的值
	// Map 被编码为键值对的列表，传统上，S 表达式使用点状符号列表（key . value）结构来表示 key/value 对，
	// 而不是用一个含双元素的列表，不过为了简单我们忽略了点状符号列表

	// 编码是由一个 encode 递归函数完成，如下所示。它的结构本质上和 Display 函数类似
	// （见 files/sexpr/encode.go）

	// Marshal 函数是对 encode 的包装，以保持和 encoding/... 下其它包有着相似的 API：
	// （见 files/sexpr/encode.go）

	// 下面是 Marshal 对 12.3 节的 strangelove 变量编码后的结果：
	// （见 files/sexpr/sexpr_test.go 的 // 编码）
	// Marshal() = ((Title "Dr. Strangelove") (Subtitle "How I Learned to Stop Worrying and Love the Bomb") (Year 1964) (Actor (("Dr. Strangelove" "Peter Sellers") ("Grp. Capt. Lionel Mandrake" "Peter Sellers") ("Pres. Merkin Muffley" "Peter Sellers") ("Gen. Buck Turgidson" "George C. Scott") ("Brig. Gen. Jack D. Ripper" "Sterling Hayden") ("Maj. T.J. \"King\" Kong" "Slim Pickens"))) (Oscars ("Best Actor (Nomin.)" "Best Adapted Screenplay (Nomin.)" "Best Director (Nomin.)" "Best Picture (Nomin.)")) (Sequel nil))

	// 整个输出编码为一行以减少输出的大小，但是也很难阅读
	// 下面是对 S 表达式手动格式化的结果：
	// ((Title "Dr. Strangelove")
	// (Subtitle "How I Learned to Stop Worrying and Love the Bomb")
	// (Year 1964)
	// (Actor (("Grp. Capt. Lionel Mandrake" "Peter Sellers")
	// 		("Pres. Merkin Muffley" "Peter Sellers")
	// 		("Gen. Buck Turgidson" "George C. Scott")
	// 		("Brig. Gen. Jack D. Ripper" "Sterling Hayden")
	// 		("Maj. T.J. \"King\" Kong" "Slim Pickens")
	// 		("Dr. Strangelove" "Peter Sellers")))
	// (Oscars ("Best Actor (Nomin.)"
	// 		 "Best Adapted Screenplay (Nomin.)"
	// 		 "Best Director (Nomin.)"
	// 		 "Best Picture (Nomin.)"))
	// (Sequel nil))

	// 编写一个 S 表达式的美化格式化函数将作为一个具有挑战性的练习任务；不过 gopl.io 也提供了一个简单版本
	// （见 files/sexpr/pretty.go）
	// MarshalIndent 输出：
	// MarshalIndent() = ((Title "Dr. Strangelove")
	// (Subtitle "How I Learned to Stop Worrying and Love the Bomb") (Year 1964)
	// (Actor
	//  (("Dr. Strangelove" "Peter Sellers")
	//   ("Grp. Capt. Lionel Mandrake" "Peter Sellers")
	//   ("Pres. Merkin Muffley" "Peter Sellers")
	//   ("Gen. Buck Turgidson" "George C. Scott")
	//   ("Brig. Gen. Jack D. Ripper" "Sterling Hayden")
	//   ("Maj. T.J. \"King\" Kong" "Slim Pickens")))
	// (Oscars
	//  ("Best Actor (Nomin.)" "Best Adapted Screenplay (Nomin.)"
	//   "Best Director (Nomin.)" "Best Picture (Nomin.)")) (Sequel nil))

	// 和 fmt.Print、json.Marshal、Display 函数类似，sexpr.Marshal 函数处理带环的数据结构也会陷入死循环

	// 在 12.6 节中，我们将给出 S 表达式解码器的实现步骤，
	// 但是在那之前，我们还需要先了解如何通过反射技术来更新程序的变量
}
