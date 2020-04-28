package main

// 008、获取结构体字段标识
func main() {
	// 在 4.5 节我们使用结构字段标签来修改 Go 结构值的 JSON 编码
	// 其中 JSON 成员标签让我们可以选择成员的名字和抑制零值成员的输出
	// 在本节，我们将看到如何通过反射机制类获取成员标签

	// 对于一个 web 服务，大部分 HTTP 处理函数要做的第一件事就是展开请求中的参数到本地变量中
	// 我们定义了一个工具函数，叫 params.Unpack，通过使用结构体成员标签机制来让 HTTP 处理函数解析请求参数更方便

	// 首先，我们看看如何使用它
	// 下面的 search 函数是一个 HTTP 请求处理函数，它定义了一个匿名结构体类型的变量，用结构体的每个成员表示 HTTP 请求的参数
	// 其中结构体成员的标签指定了参数名称，为了减少 URL 的长度这些参数名通常都是缩略词
	// Unpack 将请求参数填充到合适的结构体成员中，这样我们可以方便地通过合适的类型来访问这些参数
	// （见 files/search/main.go 的 search 函数）

	// 下面的 Unpack 函数主要完成三件事情
	// 第一，它调用 req.ParseForm() 来解析 HTTP 请求
	// 然后，req.Form 将包含所有的请求参数，不管 HTTP 客户端使用的是 GET 还是 POST 请求方法
	// 下一步，Unpack 函数将构建每个结构体成员有效参数名字到成员变量的映射
	// 如果结构体成员有成员标签的话，有效参数名字可能和实际的成员名字不同
	// reflect.Type 的 Field 方法将返回一个 reflect.StructField，里面含有每个成员的名字、类型和可选的成员标签等信息
	// 其中成员标签信息对应 reflect.StructTag 类型的字符串，并且提供了 Get 方法用于解析和根据特定 key 提取的子串，
	// 例如这里的 http:"..." 形式的子串
	// （见 files/params/params.go 的 Unpack 函数）

	// 最后，Unpack 遍历 HTTP 请求的 name/value 参数键值对，并且根据更新相应的结构体成员
	// 回想一下，同一个名字的参数可能出现多次，
	// 如果发生这种情况，并且对应的结构体成员是一个 slice，那么就将所有的参数添加到 slice 中
	// 其它情况，对应的成员值将被覆盖，只有最后一次出现的参数值才是起作用的

	// populate 函数小心用请求的字符串类型参数值来填充单一的成员 v（或者是 slice 类型成员中的单一的元素）。
	// 目前，它仅支持字符串、有符号的整数和布尔型。其中其它的类型将留作练习任务
	// （见 files/params/params.go 的 populate 函数）

	// 如果我们将上面的程序添加到 web 服务器，则可以产生以下的会话：
	// $ go build gopl.io/ch12/search
	// $ ./search &
	// $ ./fetch 'http://localhost:12345/search'
	// Search: {Labels:[] MaxResults:10 Exact:false}
	// $ ./fetch 'http://localhost:12345/search?l=golang&l=programming'
	// Search: {Labels:[golang programming] MaxResults:10 Exact:false}
	// $ ./fetch 'http://localhost:12345/search?l=golang&l=programming&max=100'
	// Search: {Labels:[golang programming] MaxResults:100 Exact:false}
	// $ ./fetch 'http://localhost:12345/search?x=true&l=golang&l=programming'
	// Search: {Labels:[golang programming] MaxResults:10 Exact:true}
	// $ ./fetch 'http://localhost:12345/search?q=hello&x=123'
	// x: strconv.ParseBool: parsing "123": invalid syntax
	// $ ./fetch 'http://localhost:12345/search?q=hello&max=lots'
	// max: strconv.ParseInt: parsing "lots": invalid syntax
}
