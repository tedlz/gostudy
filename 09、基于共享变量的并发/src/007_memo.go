package main

import (
	"io/ioutil"
	"net/http"
)

// 007、示例：并发的非阻塞缓存
func main() {

}

// 本节中我们会做一个无阻塞的缓存
// 这种工具可以帮助我们来解决现实世界中并发程序出现但没有现成的库可以解决的问题
// 这个问题叫做缓存（memoizing）函数，也就是说，我们需要缓存函数的返回结果
// 这样在对函数进行调用的时候，我们就只需要一次计算，之后只要返回计算的结果就可以了
// 我们的解决方案会是并发安全的，且会避免对整个缓存加锁而导致所有操作都去争一个锁的设计

// 我们将使用下面的 httpGetBody 函数作为我们需要缓存的函数的一个样例
// 这个函数会去进行 HTTP GET 请求并且获取 http 响应 body
// 对这个函数的调用本身开销是比较大的，所以我们尽量避免在不必要的时候反复调用
func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// 最后一行稍微隐藏了一些细节
// ReadAll 会返回两个结果，一个 []byte 数组和一个错误，
// 不过这两个对象可以被赋值给 httpGetBody 返回声明里的 interface 和 error 类型，
// 所以我们也就可以这样返回结果并且不需要额外的工作了
// 我们在 httpGetBody 中选用这种返回类型是为了使其可以与缓存匹配

// 下面是我们要设计的 cache 的第一个草稿：
// （见 files/memo1）

// Memo 实例会记录需要缓存的函数 f（类型为 Func），以及缓存内容（里面是一个 string 到 result 映射的 map）
// 每一个 result 都是简单的函数返回的值对儿（一个值和一个错误值）
// 我们将要展示一些 Memo 的变体，不过所有的例子都会遵循这些基本方面

// 下面是一个使用 Memo 的例子
// 对于流入的 URL 的每一个元素我们都会调用 Get，并打印调用延时以及其返回的数据大小的 log：
// （见 files/memotest.Sequential()）

// 我们可以使用测试包（第 11 章主题）来系统地鉴定缓存的效果
// 从下面的测试输出，我们可以看到 URL 流包含了一些重复的情况，
// 尽管我们第一次对每一个 URL 的 (*Memo).Get 的调用都会花上几百毫秒，
// 但第二次就只需要花一毫秒就可以返回完整的数据了：
// （见 files/memo1/memo_test.go 的 go test -v 命令及执行结果）

// ↑ 这个测试是顺序地去做所有的调用的

// 由于这种彼此独立的 HTTP 请求可以很好的并发，我们可以把这个测试改成并发形式
// 可以使用 sync.WaitGroup 来等待所有的请求都完成之后再返回：
// （见 files/memotest.Concurrent()）

// 这次测试跑起来更快了，然而不幸的是这个测试不是每次都能正常工作
// 我们注意到有一些预料之外的缓存未能命中（cache miss），或者命中了缓存但却返回了错误的值，或者甚至会直接崩溃

// 但更糟糕的是，有时候这个程序还是能正确的运行，所以我们甚至都不会意识到这个程序可能有 bug
// 但是我们可以使用 -race 这个 flag 来运行程序，竞争检测器（9.6 节）会打印像下面这样的报告：
// （见 files/memo1/memo_test.go 的 go test -run=TestConcurrent -v -race 命令及执行结果）

// files/memo1/memo.go:27 出现了两次，说明有两个 goroutine 在没有同步干预的情况下更新了 cache map
// 这表明 Get 不是并发安全的，存在数据竞争

// 最简单的使 cache 并发安全的方式是使用基于监控的同步
// 只要给 Memo 加上一个 mutex，在 Get 的一开始获取互斥锁，return 的时候释放锁，就可以让 cache 的操作发生在临界区内了：
// （见 files/memo2 的 Get 方法）

// 测试依然并发进行，但这回竞争检查器沉默了
// （见 files/memo2/memo_test.go 的 go test -run=TestConcurrent -v -race 命令及执行结果）
// 不幸的是对于 Memo 的这些改变使我们完全丧失了并发的性能优点
// 每次对 f 的调用期间都会持有锁，Get 将本来并行运行的 I/O 操作串行化了
// 我们本章的目的是完成一个无锁缓存，而不是现在这样的将所有请求串行化的函数的缓存

// 下一个 Get 的实现，调用 Get 的 goroutine 会两次获取锁：
// 查找阶段获取一次，如果查找没有返回任何内容，那么进入更新阶段会再次获取
// 在这两次获取锁的中间阶段，其它 goroutine 可以随意使用 cache
// （见 files/memo3 的 Get 方法）

// 这些修改使性能再次得到了提升，但有一些 URL 被获取了两次
// 这种情况在两个以上的 goroutine 同一时刻调用 Get 来请求同样的 URL 时发生
// 多个 goroutine 一起查询 cache，发现没有值，然后一起调用 f 这个慢不啦叽的函数
// 在得到结果后，也都会去更新 map，其中一个的结果会覆盖掉另一个的结果

// 理想情况下是应该避免掉多余的工作的
// 而这种 “避免” 工作一般会称为 duplicate suppression（重复抑制 / 避免）
// 下面版本的 Memo 每个 map 元素都是指向一个条目的指针
// 每一个条目包含对函数 f 调用结果的内容缓存
// 与之前不同的是这次 entry 还包含了一个叫 ready 的 channel
// 在条目的结果被设置之后，这个 channel 就会被关闭，以向其它 goroutine 广播（8.9 节）去读取该条目内的结果是安全的了
// （见 files/memo4）

// 现在 Get 函数包括下面这些步骤了：
// 获取互斥锁来保护共享变量 cache map，查询 map 中是否存在指定条目，如果没有找到那么分配空间插入一个新条目，释放互斥锁
// 如果存在条目的话且其值没有写入完成（也就是有其它的 goroutine 在调用 f 这个慢函数）时，
// goroutine 必须等待值 ready 之后才能读到条目的结果
// 而想知道是否 ready 的话，可以直接从 ready channel 中读取，因为这个读取操作在 channel 关闭前一直是阻塞

// 如果没有条目的话，需要向 map 中插入一个没有准备好的条目，
// 当前正在调用的 goroutine 就需要负责调用慢函数、更新条目以及向其它所有 goroutine 广播 [条目已经 ready 可读] 的消息了
// 条目中的 e.res.value 和 e.res.err 变量是在多个 goroutine 之间共享的
// 创建条目的 goroutine 同时也会设置条目的值，其它 goroutine 在收到 ready 的广播消息后立刻会去读取条目的值
// 尽管会被多个 goroutine 同时访问，但却并不需要互斥锁
// ready channel 的关闭一定会发生在其它 goroutine 收到广播事件之前，
// 因此第一个 goroutine 对这些变量的写操作是一定发生在这些读操作之前的，不会发生数据竞争

// 这样并发、无重复、不阻塞的 cache 就完成了

// 上面这样 Memo 的实现使用了一个互斥量来保护多个 goroutine 调用 Get 时的共享 map 变量
// 不妨把这种设计和前面提到的把 map 变量限制在一个单独的 monitor goroutine 的方案做一些对比，后者在调用 Get 时需要发消息

// Func、result 和 entry 的声明和之前保持一致（见 files/memo4/memo.go）

// 然而 Memo 类型现在包含了一个叫做 requests 的 channel，Get 的调用方用这个 channel 来和 minitor goroutine 通信
// requests channel 中的元素类型是 request
// Get 的调用方会把这个结构中的两组 key 都填充好，实际上用这两个变量来对函数进行缓存的
// 另一个叫 response 的 channel 会被拿来发送响应结果，这个 channel 只会传回一个单独的值
// （见 files/memo5）

// Get 方法，会创建一个 response channel，把它放进 request 结构中，然后发送给 monitor goroutine，然后马上又会接收它
// cache 变量被限制在了 monitor goroutine (*Memo).server 中
// monitor 会在循环中一直读取请求，直到 request channel 被 Close 方法关闭
// 每一个请求都会去查询 cache，如果没有找到条目的话，那么就会创建 / 插入一个新的条目
// （见 files/memo5/memo.go 中被注释 !+ monitor、!- monitor 中包裹的内容）

// 和基于互斥量的版本类似，第一个对某个 key 的请求需要负责去调用函数 f 并传入这个 key，
// 将结果存在条目里，并关闭 ready channel 来广播条目的 ready 消息
// 使用 (*entry).call 来完成上述工作

// 紧接着对同一个 key 的请求会发现 map 中已经有了存在的条目，
// 然后会等待结果变为 ready，并将结果从 response 发送给客户端的 goroutine
// 上述工作是用 (*entry).deliver 来完成的
// 对 call 和 deliver 方法的调用，
// 必须让它们在自己的 goroutine 中进行以确保 monitor goroutine 不会因此被阻塞而无法处理新的请求

// 这个例子说明我们无论用上锁，还是通信来建立并发程序都是可行的

// 上面两种方案并不好说特定场景下哪种方案更好，不过了解它们还是有价值的
// 有时候从一种方式切换到另一种可以使你的代码更为简洁
