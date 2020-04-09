package main

// 004、基准测试
func main() {

}

// 基准测试是测量一个程序在固定工作负载下的性能
// 在 Go 语言中，基准测试函数和普通测试函数的写法类似，但是以 Benchmark 为前缀名，并且带有一个 *testing.B 类型的参数
// ，*testing.B 参数除了提供和 *testing.T 类似的方法，还有额外的一些和性能测量相关的方法
// 它还提供了一个整数 N，用于指定操作执行的循环次数

// 下面是 IsPalindrome 函数的基准测试，其中循环将执行 N 次：
// （见 files/word2/word_test.go 的 BenchmarkIsPalindrome 函数）

// 我们用下面的命令运行基准测试。和普通测试不同的是，默认情况下不运行任何基准测试
// 我们需要通过 -bench 命令行标志参数手工指定要运行的基准测试函数
// 该参数是一个正则表达式，用于匹配要执行的基准测试函数的名字，默认值是空的
// 其中 "." 模式将可以匹配所有基准测试函数，但因为这里只有一个基准测试函数，因此和 -bench=IsPalindrome 参数是等价的
// $ cd $GOPATH/src/gopl.io/ch11/word2
// $ go test -bench=.
// PASS
// BenchmarkIsPalindrome-8 1000000                1035 ns/op
// ok      gopl.io/ch11/word2      2.179s

// 结果中基准测试名的数字后缀部分，这里是 8，表示运行时对应的 GOMAXPROCS 的值，这对于一些与并发相关的基准测试是重要的信息
// 报告显示每次调用 IsPalindrome 函数花费 1.035 微秒，是执行 100 万次的平均时间
// 因为基准测试驱动器开始时并不知道每个基准测试函数运行所花的时间，它会尝试在真正运行基准测试前，
// 先尝试用较小的 N 运行测试来估算基准测试函数所需要的时间，然后推断一个较大的时间保证稳定的测量结果

// 循环在基准测试函数内实现，而不是放在基准测试框架内实现，
// 这样可以让每个基准测试函数有机会在循环启动前执行初始化代码，这样并不会显著影响每次迭代的平均运行时间
// 如果还是担心初始化代码部分对测量时间带来干扰，那么可以通过 *testing.B 参数提供的方法来临时关闭或重置计时器，
// 不过这些一般很少会用到

// 现在我们有了一个基准测试和普通测试，我们可以很容易测试改进程序运行速度的想法
// 也许最明显的优化是在 IsPalindrome 函数中第二个循环的停止检查，这样可以避免每个比较都做两次
// （见 files/word2/word.go 的 11.4 节优化 1）

// 不过很多情况下，一个显而易见的优化未必能带来预期的效果
// 这个改进在基准测试中只带来了 4% 的性能提升
// $ go test -bench=.
// PASS
// BenchmarkIsPalindrome-8 1000000              992 ns/op
// ok      gopl.io/ch11/word2      2.093s

// 另一个改进想法是在开始为每个字符预先分配一个足够大的数组，这样就可以避免在 append 调用时可能会导致内存的多次重新分配
// 声明一个 letters 数组变量，并指定合适的大小，像下面这样：
// （见 files/word2/word.go 的 11.4 节优化 2）

// 这个改进提升性能约 35%，报告结果是基于 200 万次迭代的平均运行时间统计：
// $ go test -bench=.
// PASS
// BenchmarkIsPalindrome-8 2000000                      697 ns/op
// ok      gopl.io/ch11/word2      1.468s

// 如这个例子所示，快的程序往往是伴随着较少的内存分配
// -benchmem 命令行标志参数将在报告中包含内存的分配数据统计，我们可以比较优化前后的内存分配情况：
// （优化前）
// $ go test -bench=. -benchmem
// PASS
// BenchmarkIsPalindrome    1000000   1026 ns/op    304 B/op  4 allocs/op

// 这是优化之后的结果：
// $ go test -bench=. -benchmem
// PASS
// BenchmarkIsPalindrome    2000000    807 ns/op    128 B/op  1 allocs/op
// 用一次内存分配替代多次内存分配节省了 75% 的分配调用次数和减少近一半的内存需求

// 这个基准测试告诉了我们某个具体操作所需的绝对时间，但我们往往想知道的是两个不同操作的时间对比
// 例如，如果一个函数需要 1ms 处理 1000 个元素，那么处理 10,000 或 1,000,000 将需要多少时间呢？
// 这样的比较揭示了渐进增长函数的运行时间
// 另一个例子：I/O 缓存该设置为多大呢？基准测试可以帮助我们选择在性能达标的情况下所需的最小内存
// 第三个例子：对于一个确定的工作哪种算法更好？基准测试可以评估两种不同算法对于相同输入在不同场景和负载下的优缺点

// 比较型的基准测试就是普通程序代码
// 它们通常是单参数的函数，由几个不同数量级的基准测试函数调用，就像这样：
// func benchmark(b *testing.B, size int) { /* ... */ }
// func Benchmark10(b *testing.B)         { benchmark(b, 10) }
// func Benchmark100(b *testing.B)        { benchmark(b, 100) }
// func Benchmark1000(b *testing.B)       { benchmark(b, 1000) }

// 通过函数参数来指定输入的大小，但是参数变量对于每个具体的基准测试都是固定的
// 要避免直接修改 b.N 来控制输入的大小
// 除非你将它作为一个固定大小的迭代计算输入，否则基准测试的结果将毫无意义

// 比较型的基准测试反映出的模式在程序设计阶段是很有帮助的，但是即使程序完工了也应当保留基准测试代码
// 因为随着项目的发展或者是输入的增加，或者是部署到新的操作系统或不同处理器，我们可以再次用基准测试来帮助我们改进设计
