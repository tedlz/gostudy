package main

// 005、剖析
func main() {

}

// 基准测试（Benchmark）对于衡量特定操作的性能是有帮助的
// 但是当我们试图让程序跑的更快的时候，通常我们并不知道应该从哪里优化

// 当我们想仔细观察我们程序的运行速度的时候，最好的方法是性能剖析
// 剖析技术是基于程序执行期间的一些自动抽样，然后在收尾时进行推断；最后产生的统计结果就称为剖析数据

// Go 语言支持多种类型的剖析性能分析，每一种关注不同的方面，
// 但它们都涉及到每个采样记录的感兴趣的一系列事件消息，每个事件都包含函数调用时函数调用堆栈的信息
// 内建的 go test 工具对几种分析方式都提供了支持

// CPU 剖析数据标识了最耗 CPU 时间的函数
// 在每个 CPU 上运行的线程在每隔几毫秒都会遇到操作系统的中断事件，每次中断时都会记录一个剖析数据然后恢复正常的运行

// 堆剖析则标识了最耗内存的语句
// 剖析库会记录调用内部内存分配的操作，平均每 512KB 的内存申请会触发一个剖析数据

// 阻塞剖析则记录阻塞 goroutine 最久的操作，例如系统调用，管道发送和接收，还有获取锁等
// 每当 goroutine 被这些操作阻塞时，剖析库都会记录相应的事件

// 只需要开启下面其中一个标志参数就可以生成各种分析文件
// 当同时使用多个标志参数时需要当心，因为一项分析操作可能会影响其它项的分析结果

// $ go test -cpuprofile=cpu.out
// $ go test -blockprofile=block.out
// $ go test -memprofile=mem.out

// 对于一些非测试程序也很容易进行剖析，具体的实现方式，与程序是短时间运行的小工具还是长时间运行的服务会有很大不同
// 剖析对于长期运行的程序尤其有用，因此可以通过调用 Go 的 runtime API 来启用运行时剖析

// 一旦我们已经收集到了用于分析的采样数据，我们就可以使用 pprof 来分析这些数据
// 这是 Go 工具箱自带的一个工具，但并不是一个日常工具，它对应 go tool pprof 命令
// 该命令有许多特性和选项，但是最基本的是两个参数：生成这个概要文件的可执行程序和对应的剖析数据

// 为了提高分析效率和减少空间，分析日志本身并不包含函数的名字，它只包含函数对应的地址
// 也就是说 pprof 需要对应的可执行程序来解读剖析数据
// 虽然 go test 通常在测试完成后就丢弃临时用的测试程序，
// 但是在启用分析的时候会将测试程序保存为 foo.test 文件，其中 foo 部分对应待测包的名字

// 下面的命令演示了如何收集并展示一个 CPU 分析文件
// 我们选择 net/http 包的一个基准测试为例
// 通常最好是对业务关键代码的部分设计专门的基准测试
// 因为简单的基准测试没法代表业务场景，因此我们用 run=NONE 参数禁止那些简单测试

// $ go test -run=NONE -bench=ClientServerParallelTLS64 \
//     -cpuprofile=cpu.log net/http
//  PASS
//  BenchmarkClientServerParallelTLS64-8  1000
//     3141325 ns/op  143010 B/op  1747 allocs/op
// ok       net/http       3.395s

// $ go tool pprof -text -nodecount=10 ./http.test cpu.log
// 2570ms of 3590ms total (71.59%)
// Dropped 129 nodes (cum <= 17.95ms)
// Showing top 10 nodes out of 166 (cum >= 60ms)
//     flat  flat%   sum%     cum   cum%
//   1730ms 48.19% 48.19%  1750ms 48.75%  crypto/elliptic.p256ReduceDegree
//    230ms  6.41% 54.60%   250ms  6.96%  crypto/elliptic.p256Diff
//    120ms  3.34% 57.94%   120ms  3.34%  math/big.addMulVVW
//    110ms  3.06% 61.00%   110ms  3.06%  syscall.Syscall
//     90ms  2.51% 63.51%  1130ms 31.48%  crypto/elliptic.p256Square
//     70ms  1.95% 65.46%   120ms  3.34%  runtime.scanobject
//     60ms  1.67% 67.13%   830ms 23.12%  crypto/elliptic.p256Mul
//     60ms  1.67% 68.80%   190ms  5.29%  math/big.nat.montgomery
//     50ms  1.39% 70.19%    50ms  1.39%  crypto/elliptic.p256ReduceCarry
//     50ms  1.39% 71.59%    60ms  1.67%  crypto/elliptic.p256Sum

// 参数 -text 用于指定输出格式，在这里每行是一个函数，根据使用 CPU 的时间长短来排序
// 其中 -nodecount=10 参数限制了只输出前 10 行的结果
// 对于严重的性能问题，这个文本格式基本可以帮助查明原因了

// 这个概要文件告诉我们，HTTPS 基准测试中 crypto/elliptic.p256ReduceDegree 函数占用了将近一半的 CPU 资源，对性能占很大比重
// 相比之下，如果一个概要文件中主要是 runtime 包的内存分配的函数，那么减少内存消耗可能是一个值得尝试的优化策略

// 对于一些更微妙的问题，你可能需要 pprof 的图形显示功能
// 这个需要 GraphViz 工具，可以从 http://www.graphviz.org 下载
// 参数 -web 用于生成函数的有向图，标注有 CPU 的使用和最热点的函数等信息

// 这一节我们只是简单看了下 Go 语言的分析工具
// 如果想了解更多，可以阅读 Go 官方博客的 Profiling Go Programs 一文
