package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// 013、Channels - 带缓存的 Channels
func main() {
	// 带缓存的 channel 内部持有一个元素队列
	// 队列的最大容量是在调用 make 函数创建 channel 时通过第二个参数指定的
	// 下面的语句创建了一个可以持有三个字符串元素的带缓存的 channel
	ch := make(chan string, 3)
	// 向缓存 channel 的发送操作就是向内部缓存队列的尾部插入元素，接收操作则是从队列的头部删除元素
	// 如果 channel 内部缓存队列是满的，那么发送操作将阻塞直到因另一个 goroutine 执行接收操作而释放了新的队列空间
	// 如果 channel 内部缓存队列是空的，那么接收操作将阻塞直到有另一个 goroutine 执行发送操作而向队列插入了新元素

	// 我们可以在无阻塞的情况下连续向新创建的 channel 发送三个值：
	ch <- "A"
	ch <- "B"
	ch <- "C"
	// 此刻，channel 的内部缓存队列将是满的，如果有第四个发送操作将发生阻塞
	// 如果我们接收一个值
	fmt.Println(<-ch) // A
	// 那么 channel 的缓存队列将不是满的也不是空的
	// 因此，对该 channel 执行的发送或接收操作都不会发生阻塞
	// 通过这种方式，channel 的缓存队列解耦了接收和发送的 goroutine

	// 在某些特殊情况下，程序可能需要知道 channel 内部缓存的容量，可以用内置的 cap 函数获取
	fmt.Println(cap(ch)) // 3

	// 同样，对于内置的 len 函数，如果传入的是 channel，那么将返回 channel 内部缓存队列中有效元素的个数
	// 因为在并发程序中，该信息会随着接收操作而失效，但是它对某些故障诊断和性能优化会有帮助
	fmt.Println(len(ch)) // 2

	// 在继续执行两次接收操作后，channel 的内部缓存队列将又成为空的，如果有第四个接收操作将发生阻塞
	fmt.Println(<-ch) // B
	fmt.Println(<-ch) // C

	// 在这个例子中，发送和接收操作都发生在同一个 goroutine 中，但是在真实的程序中它们一般都是由不同的 goroutine 执行
	// Go 语言新手有时会将一个带缓存的 channel 当做同一个 goroutine 中的队列使用，虽然语法看似简单，但实际上这是一个错误
	// channel 和 goroutine 的调度机制是紧密相连的，一个发送操作 —— 或许是整个程序 —— 可能会永远阻塞
	// 如果你只是需要一个简单的队列，使用 slice 就可以了

	// 下面的例子展示了一个使用了带缓存 channel 的应用
	// 它并发的向三个镜像站点发出了请求，三个镜像站点分散在不同的地理位置
	// 它们分别将收到的响应发送到带缓存的 channel，
	// 最后接收者只接收第一个收到的响应，也就是最快的那个响应
	// 因此 mirroredQuery 函数可能在另外两个响应慢的站点响应之前就返回了结果
	// （顺便说一下，多个 goroutine 并发的向同一个 channel 发送数据，或从同一个 channel 接收数据都是常见的用法）
	fmt.Println(mirroredQuery())

	// 如果我们使用了无缓存的 channel，那么两个慢的 goroutine 将会因为没有人接收而被永远卡住
	// 这种情况，称为 goroutine 泄露，这将是一个 bug
	// 和垃圾变量不同，泄露的 goroutine 并不会被自动回收，因此确保每个不再需要的 goroutine 能正常退出是重要的

	// 关于无缓存的或有缓存的 channels 之间的选择，或者是带缓存的 channels 的容量大小的选择，都可能影响程序的正确性
	// 无缓存 channel 更强的保证了每个发送操作与相应的同步接收操作；
	// 但是对于带缓存的 channel，这些操作是解耦的
	// 同样，即使我们知道将要发送到一个 channel 的信息的数量上限，创建一个对应容量大小的带缓存 channel 也是不现实的
	// 因为这要求在执行任何接收操作之前缓存所有已经发送的值，如果未能分配足够的缓冲将导致程序死锁

	// channel 的缓存也可能影响程序的性能
	// 想象一家蛋糕店有三个厨师，一个烘焙，一个上糖衣，还有一个将每个蛋糕传递给下一个在生产线上的厨师，
	// 在狭小的厨房环境，每个厨师在完成自己的工作后必须等待下一个厨师已经准备好接收它，
	// 这类似于在一个无缓存的 channel 上进行沟通

	// 如果在每个厨师之间有放置一个蛋糕的额外空间，那么每个厨师就可以将一个完成的蛋糕临时放在那里而马上进入下一个蛋糕的制作
	// 这类似于将 channel 的缓存队列的容量设置为 1
	// 只要每个厨师的平均工作效率相近，那么其中大部分的传输工作将是迅速的，个体之间细小的效率差异将在交接过程中弥补
	// 如果厨师之间有更大的额外空间 —— 也就是更大容量的缓存队列 —— 将可以在不停止生产线的前提下消除更大的效率波动，
	// 例如一个厨师可以短暂的休息，然后再加快赶上进度而不影响其它人

	// 另一方面，如果生产线的前期阶段一直快于后续阶段，那么它们之间的缓存在大部分时间都将是满的
	// 相反，如果后续阶段比前期阶段更快，那么它们之间的缓存在大部分时间都将是空的
	// 对于这类场景，额外的缓存并没有带来任何的好处

	// 生产线的隐喻对于理解 channels 和 goroutines 的工作机制是很有帮助的
	// 例如，如果第二阶段是需要精心制作的复杂操作，那么该厨师可能无法跟上第一阶段厨师的进度，或无法满足第三阶段厨师的需求
	// 要解决这个问题，我们可以雇佣另一个厨师来帮助完成第二阶段的工作，他执行相同的任务但是独立工作
	// 这类似于基于相同的 channels 创建另一个独立的 goroutine

	// 我们没有太多空间展示全部的细节，但是 gopl.io/ch8/cake 包模拟了这个蛋糕店，可以通过不同的参数调整
	// 它还对上面提到的几种场景提供对应的基准测试（见 11.4 基准测试章节）
}

func mirroredQuery() string {
	responses := make(chan string, 3)
	go func() { responses <- request("http://asia.gopl.io") }()
	go func() { responses <- request("http://europe.gopl.io") }()
	go func() { responses <- request("http://americas.gopl.io") }()
	return <-responses
}

func request(hostname string) string {
	res, err := http.Get(hostname)
	if err != nil {
		return fmt.Sprintln(err)
	}
	defer res.Body.Close()
	// content, err := ioutil.ReadAll(res.Body)
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Sprintln(err)
	}
	// return string(content)
	return hostname
}
