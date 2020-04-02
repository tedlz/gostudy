package main

// 006、包和命名
func main() {

}

// 在本节中，我们将提供一些关于 Go 语言独特的包和成员命名的约定

// 创建包时，一般要使用短小的包名，但也不能太短导致难以理解
// 标准库中最常用的包有 bufio、bytes、flag、fmt、http、io、json、os、sort、sync 和 time 等包

// 它们的名字都简洁明了
// 例如，不要将一个类似 imageutil 或 ioutil 的通用包命名为 util，虽然它看起来很短小
// 要尽量避免包名使用可能被经常用于局部变量的名字，这样可能导致用户重命名导入包，例如前面看到的 path 包

// 包名一般采用单数的形式
// 标准库的 bytes、errors 和 strings 使用了复数形式，这是为了避免和预定义的类型冲突
// 同样的例子还有 go/types 则是为了避免和 type 关键字冲突

// 要避免包名有其它的含义
// 例如，2.5 节中我们的温度转换包最初使用了 temp 包名，虽然并没有持续多久，但这是一个糟糕的尝试，
// 因为 temp 几乎是临时变量的同义词。然后我们有一段时间使用了 temperature 作为包名，但是名字并没有表达包的真实用途，
// 最后我们改成了和 strconv 标准包类似的 tempconv 包名，这个名字比之前的就好多了

// 现在让我们看看如何命名包的成员
// 由于是通过导入的包的名字引用包里面的成员，例如 fmt.Println 同时包含了包名和成员名信息
// 因此，我们一般并不需要关注 Println 的具体内容，因为 fmt 包名已经包含了这个信息

// 当设计一个包的时候，需要考虑包名和成员名两个部分如何更好的配合，下面有一些例子：
//   - bytes.Equal
//   - flag.Int
//   - http.Get
//   - json.Marshal

// 我们可以看到一些常用的命名模式。strings 包提供了和字符串相关的诸多操作：

// package strings
//
// func Index(need, haystack string) int
//
// type Replacer struct { /* */ }
// func NewReplacer(oldnew ...string) *Replacer
//
// type Reader struct { /* */ }
// func NewReader(s string) *Reader

// 包名 strings 并没有出现在任何成员名字中
// 因为用户会这样引用这些成员 strings.Index、strings.Replacer 等

// 其它一些包，可能只描述了单一的数据类型，例如 html/template 和 math/rand 等
// 只暴露一个主要的数据结构和与它相关的方法，还有一个以 New 命名的函数用于创建实例：

// package rand // math rand
//
// type rand struct { /* */ }
// func New(source Source) *Rand

// 这可能导致一些名字重复，例如 template.Template 或 rand.Rand，这就是为什么这些种类的包名往往特别短的原因之一

// 在另一个极端，还有像 net/http 包那样，含有非常多的名字和种类不多的数据类型，因为它们都是要执行一个复杂的复合任务
// 尽管有将近 20 种类型和更多的函数，但是包中最重要的成员名字确是简单明了的：
// Get、Post、Handle、Error、Client、Server 等
