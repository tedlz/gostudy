package thumbnail_test

import (
	"gostudy/08、Goroutines和Channels/files/thumbnail"
	"log"
	"os"
	"sync"
)

// 本节中，我们会探索一些用来在并行时循环迭代的常见并发模型
// 我们会探究从全尺寸图片生成一些缩略图的问题

// 下面的程序会循环迭代一些图片文件，并为每一张图片生成一个缩略图

// makeThumbnails 制作指定文件的缩略图
func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

// 显然我们处理文件的顺序无关紧要，因为每个文件的缩放操作和其它图片的处理操作都是彼此独立的
// 像这种子问题都是完全彼此独立的问题被叫做易并行问题
// 易并行问题是最容易被实现成并行的一类问题，并且最能够享受并发带来的好处，能够随着并行的规模线性的扩展
// 下面让我们并行的执行这些操作，从而将文件 I/O 的延迟隐藏掉，并用上多核 CPU 的计算能力来拉伸图像
// 我们的第一个并发程序只是使用了一个 go 关键字，这里我们先忽略掉错误，之后再进行处理

// 注意：不正确！
func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		go thumbnail.ImageFile(f) // 注意：忽略错误
	}
}

// 这个版本运行的实在有点太快，实际上，由于它比最早的版本使用的时间要短的多，即使当文件名的 slice 中只包含有一个元素
// 这就有点奇怪了，如果程序没有并发执行的话，那为什么一个并发的版本还是要快呢？
// 答案其实是 makeThumbnails 在它还没有完成工作之前就已经返回了
// 它启动了所有的 goroutine，每一个文件名对应一个，但没有等待它们一直到执行完毕

// 没有什么直接的办法能够等待 goroutine 完成，但是我们可以改变 goroutine 里的代码，
// 让其能够将完成情况报告给外部的 goroutine 知晓，使用的方式是向一个共享的 channel 中发送事件
// 因为我们已经确切的知道有 len(filenames) 个内部 goroutine，所以外部的 goroutine 只需要在返回之前对这些事件计数

// makeThumbnails3 并行创建指定文件的缩略图
func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) {
			thumbnail.ImageFile(f) // 注意：忽略错误
			ch <- struct{}{}
		}(f)
	}
	// 等待 goroutine 完成
	for range filenames {
		<-ch
	}
}

// 注意，我们将 f 的值作为一个显式变量传给了函数，而不是在循环的闭包中声明：
// for _, f := range filenames {
// 	go func() {
// 		thumbnail.ImageFile(f) // 注意：不正确！
// 		...
// 	}()
// }

// 回忆一下在之前 5.6.1 一节中，匿名函数中的循环变量快照问题
// 上面这个单独的变量 f 是被所有的匿名函数值所共享，且会被连续的循环迭代所更新的
// 当新的 goroutine 开始执行字面函数时，for 循环可能已经更新了 f 并且开始了另一轮的迭代或者（更有可能）已经结束了整个循环
// 所以当这些 goroutine 开始读取 f 的值时，它们所看到的值已经是 slice 的最后一个元素了
// 显式的添加这个参数，我们才能确保使用的 f 是当 go 语句执行时的 “当前” 那个 f

// 如果我们想要从每一个 worker goroutine 往主 goroutine 中返回值时该怎么办呢？
// 当我们调用 thumbnail.ImageFile 创建文件失败的时候，它会返回一个错误
// 下一个版本的 makeThumbnails 会返回其在做缩放操作时接收到的第一个错误

// makeThumbnails4 并行创建指定文件的缩略图
// 如果任何一步失败，它将返回错误
func makeThumbnails4(filenames []string) error {
	errors := make(chan error)

	for _, f := range filenames {
		go func(f string) {
			_, err := thumbnail.ImageFile(f)
			errors <- err
		}(f)
	}

	for range filenames {
		if err := <-errors; err != nil {
			return err // 注意：错误：goroutine 泄露
		}
	}

	return nil
}

// 这个程序有一个微妙的 bug
// 当它遇到第一个非 nil 的 error 时会直接将 error 返回到调用方，使得没有一个 goroutine 去排空 errors channel
// 这样剩下的 worker goroutine 在向这个 channel 中发送值时，都会永远的阻塞下去，并且永远都不会退出
// 这种情况叫做 goroutine 泄露（见 8.4.4），可能会导致整个程序卡住或者跑出 out of memory 的错误

// 最简单的解决办法就是用一个具有合适大小的 buffered channel，
// 这样这些 worker goroutine 向 channel 中发送错误时就不会被阻塞
// （一个可选的解决办法是创建一个另外的 goroutine，当 main goroutine 返回第一个错误的同时去排空 channel）

// 下一个版本的 makeThumbnails 使用了一个 buffered channel 来返回生成的图片文件的名字，附带生成时的错误

// makeThumbnails5 并行创建指定文件的缩略图
// 它以任意顺序返回生成的文件名，或错误（如果任何一步失败）
func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}

	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = thumbnail.ImageFile(f)
			ch <- it
		}(f)
	}

	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}

	return thumbfiles, nil
}

// 我们最后一个版本的 makeThumbnails 返回了新文件们的大小总计数
// 和前面版本都不同的一点是，我们在这个版本里没有把文件名放入 slice，而是通过一个 string 的 channel 传过来
// 所以我们无法对循环的次数进行预测

// 为了知道最后一个 goroutine 什么时候结束（最后一个结束并不一定是最后一个开始），我们需要一个递增的计数器
// 在每一个 goroutine 启动时加 1，在 goroutine 退出时减 1
// 这需要一种特殊的计数器，这个计数器需要在多个 goroutine 操作时做到安全，并且提供在其减为零之前一直等待的一种方法
// 这种计数类型被称为 sync.WaitGroup，下面的代码就用到了这种方法

// makeThumbnails6 为从 channel 接收的每个文件制作缩略图
// 它返回它创建的文件占用的字节数
func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup // 工作中 goroutine 的数量
	for f := range filenames {
		wg.Add(1)
		// worker
		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb) // 忽略错误
			sizes <- info.Size()
		}(f)
	}

	// closer
	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}
	return total
}

// 注意 Add 和 Done 的方法不对称
// Add 是为计数器加一，必须在 worker goroutine 开始之前调用，而不是在 goroutine 中；
// 否则的话我们没办法确定 Add 是在 closer goroutine 调用 Wait 之前被调用
// 并且 Add 还有一个参数，但 Done 却没有任何参数；其实它和 Add(-1) 是等价的
// 我们使用 defer 来确保计数器即使是在出错的情况下依然能够正确的被减掉
// 上面的程序代码结构是当我们使用并发循环，但又不知道迭代次数时很通常而且很地道的写法

// size channel 携带了每一个文件的大小到 main goroutine，在 main goroutine 中使用了 range loop 来计算总和
// 观察一下我们是怎样创建一个 closer goroutine，并让其在所有 worker goroutine 们结束之后再关闭 sizes channel 的
// 两步操作：wait 和 close，必须是基于 sizes 的循环的并发
// 考虑一下另一种方案：如果等待操作被放在了 main goroutine 中，在循环之前，这样的话就永远都不会结束了，
// 如果在循环之后，那么又变成了不可达的部分，因为没有任何东西去关闭这个 channel，这个循环就永远都不会终止
