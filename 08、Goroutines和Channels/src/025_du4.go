package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// 025、并发的退出
// go run 025_du4.go -v /bin /usr /data $HOME
func main() {
	// 有时候我们需要通知 goroutine 停止它正在干的事情，
	// 比如一个正在执行计算的 web 服务，然而它的客户端已经断开了和服务端的连接

	// Go 语言并没有提供在一个 goroutine 中终止另一个 goroutine 的方法，
	// 因为这样会导致 goroutine 之间的共享变量落在未定义的状态上

	// 在 8.7 节中的 rocket launch 程序中，我们往名字叫 abort 的 channel 里发送了一个简单的值
	// 在 countdown 的 goroutine 中会把这个值理解为自己的退出信号，但如果我们想要退出两个或任意多个 goroutine 怎么办呢？

	// 一种可能的手段是向 abort 的 channel 里发送和 goroutine 数目一样多的事件来退出它们
	// 如果这些 goroutine 中已经有一些自己退出了，
	// 那么我们 channel 里的事件数比 goroutine 还多，这样会让我们的发送直接被阻塞
	// 另一方面，如果这些 goroutine 又生成了其它的 goroutine，
	// 我们 channel 里的事件又太少了，所以有些 goroutine 可能会无法接收到退出消息
	// 一般情况下我们很难知道在某一时刻具体有多少个 goroutine 在运行着的

	// 另外，当一个 goroutine 从 abort channel 中接收到一个值的时候，
	// 它会消费掉这个值，这样其它的 goroutine 就没法看到这条消息
	// 为了能够达到我们退出 goroutine 的目的，我们需要更靠谱的策略，来通过一个 channel 把消息广播出去，
	// 这样 goroutine 们能够看到这条事件消息，并且在事件完成之后，可以知道这件事已经发生过了

	// 回忆一下，我们关闭了一个 channel 并且被消费掉了所有已发送的值，
	// 操作 channel 之后的代码可以立即被执行，并且会产生零值
	// 我们可以将这个机制扩展一下，来作为我们的广播机制：
	// 不要向 channel 发送值，而是用关闭一个 channel 来进行广播

	// 只要一个小修改，我们就可以把退出逻辑加入到前一节的 du 程序
	// 首先，我们创建一个退出的 channel，这个 channel 不会向其中发送任何值，但其所在的闭包内要写明程序需要退出
	// 我们同时还定义了一个工具函数，cancelled，这个函数在被调用的时候会轮询退出状态

	// 确定初始目录
	roots := os.Args[1:]
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// 下面我们创建了一个标准的输入流读取内容的 goroutine，这是一个比较典型的连接到终端的程序
	// 每当有输入被读到（比如用户按了回车键），这个 goroutine 就会把取消消息通过关闭 done 的 channel 广播出去
	// 检测到输入时取消遍历
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	// 并行遍历文件树的每个根
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir4(root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	// 定期打印结果
	tick := time.Tick(500 * time.Millisecond)
	var nfiles, nbytes int64
loop:

	// 现在我们需要使我们的 goroutine 来对取消进行响应
	// 在 main goroutine 中，我们添加了 select 的第三个 case 语句，尝试从 done channel 中接收内容
	// 如果这个 case 被满足的话，在 select 到的时候即会返回，
	// 但在结束之前我们需要把 fileSizes channel 中的内容排空，在 channel 被关闭之前，舍弃掉所有值
	// 这样可以保证对 walkDir 的调用不要被向 fileSizes 发送信息阻塞住，可以正确地完成
	for {
		select {
		case <-done:
			// 清空 fileSizes 以允许现有的 goroutines 完成
			for range fileSizes {
				// 什么也不做
			}
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes 已关闭
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage4(nfiles, nbytes)
		}
	}
	printDiskUsage4(nfiles, nbytes)
}

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func printDiskUsage4(nfiles, nbytes int64) {
	fmt.Printf("%d files	%.1f GB\n", nfiles, float64(nbytes)/1e9)
}

// walkDir 这个 goroutine 一启动就会轮询取消状态，如果取消状态被设置的话会直接返回，并且不做额外的事情
// 这样我们将所有在取消事件之后创建的 goroutine 改变为无操作
// walkDir 递归遍历以 dir 为根的文件树，并在 fileSizes 上发送每个找到的文件大小
func walkDir4(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirents4(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir4(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// sema 用来限制 dirents4 中的并发数量
var sema = make(chan struct{}, 20)

// 在 walkDir 函数的循环中我们对取消状态进行轮询可以带来明显的益处，可以避免在取消事件发生时还去创建 goroutine
// 取消本身是有一些代价的，想要快速的响应，需要对程序逻辑进行侵入式的修改
// 确保在取消发生之后不要有代价太大的操作可能会需要修改你代码里的很多地方，
// 但是在一些重要的地方去检查取消事件也确实能带来很大的好处

// 对这个程序的一个简单的性能分析可以揭示瓶颈在 dirents 函数中获取一个信号量
// 下面的 select 可以让这种操作可以被取消，并且可以将取消时的延迟从几百毫秒降低到几十毫秒
// dirents 返回目录 dir 的条目
func dirents4(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}: // 获取 token
	case <-done:
		return nil // cancelled
	}
	defer func() { <-sema }() // 释放 token

	f, err := os.Open(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du4.1: %v\n", err)
	}
	defer f.Close()

	entries, err := f.Readdir(0) // 0 => 不限，读取所有条目
	if err != nil {
		fmt.Fprintf(os.Stderr, "du4.2: %v\n", err)
		// 不要 return，Readdir 可能只返回部分结果
	}
	return entries
}

// 现在当取消发生时，所有后台的 goroutine 都会迅速停止并且主函数会返回
// 当然，当主函数返回时，一个程序会退出，而我们又无法在主函数退出的时候确认其已经释放了所有的资源
// （译注：因为程序都退出了，你的代码都没法执行了）
// 这里有一个方便的窍门我们可以一用：
// 取代掉直接从主函数返回，我们调用一个 panic，然后 runtime 会把每一个 goroutine 的栈 dump 下来
// 如果 main goroutine 是唯一一个剩下来的 goroutine 的话，它会清理掉自己的一切资源
// 但是如果其它的 goroutine 没有退出，它们可能没办法被正确地取消掉，也有可能会被取消但是取消操作很花时间
// 所以这里的一个调研还是很有必要的
// 我们用 panic 来获取足够的信息来验证我们上面的判断，看看最终到底是什么样的情况
