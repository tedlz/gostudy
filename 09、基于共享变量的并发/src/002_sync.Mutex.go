package main

// 002、sync.Mutex 互斥锁
func main() {
	// 在 8.6 节中，我们使用了一个 buffered channel 作为一个计数信号量，
	// 来保证最多只有 20 个的 goroutine 会同时执行 HTTP 请求
	// 同理，我们可以用一个容量只有 1 的 channel 来保证最多只有一个 goroutine 在同一时刻访问一个共享变量
	// 一个只能为 0 和 1 的信号量叫做二元信号量（binary semaphore）
	//
	// （此处代码见 files/bank2）
	//
	// 这种互斥很实用，而且被 sync 包里的 Mutex 类型直接支持
	// 它的 Lock 方法能够获取到 token（锁），并且 Unlock 方法会释放这个 token
	//
	// （此处代码见 files/bank3）
	//
	// 每次一个 goroutine 访问 bank 的 balance 变量时，它都会调用 mutex 的 Lock 方法来获取一个互斥锁
	// 如果其它的 goroutine 已经获得了这个锁的话，这个操作会被阻塞直到其它 goroutine 调用了 Unlock 使该锁变为可用状态
	// mutex 会保护共享变量
	// 惯例来说，被 mutex 所保护的变量是在 mutex 变量声明之后立刻声明的
	// 如果你的做法和惯例不符，确保在文档里对你的做法进行说明

	// 在 Lock 和 Unlock 之间的代码段中的内容，goroutine 可以随便读取或者修改，这个代码段叫做临界区
	// 锁的持有者在其它 goroutine 获取该锁之前需要调用 Unlock
	// goroutine 在结束后释放锁是必要的，无论以哪条路径通过函数都需要释放，即使是在错误路径中，也要记得释放

	// 上面的 bank 程序例证了一种通用的并发模式
	// 一系列的导出函数封装了一个或多个变量，那么访问这些变量唯一的方式就是通过这些函数来做（或者方法，对于一个对象的变量来说）
	// 每一个函数在一开始就获取互斥锁并在最后释放锁，从而保证共享变量不会被并发访问
	// 这种函数、互斥锁和变量的编排叫做监控 monitor
	// (这种老式单词的 monitor 是受 monitor goroutine 的术语启发而来，两种用法都是一个代理人保证变量被顺序访问)

	// 由于在存款和查询余额函数中的临界区代码只有一行，没有分支调用 —— 在代码最后去调用 Unlock 就显得更为直截了当
	// 在更复杂的临界区的应用中，尤其是必须要尽早处理错误并返回的情况下，
	// 就很难去（靠人）判断对 Lock 和 Unlock 的调用是在所有路径中都能够严格配对的了
	// Go 语言里的 defer 简直就是这种情况下的救星：我们用 defer 来调用 Unlock，临界区会隐式的延伸到函数作用域的最后，
	// 这样我们就从 “总要记得函数在返回之后或者发生错误返回时要记得调用一次 Unlock” 这种状态中获得了解放
	// Go 会自动帮我们完成这些事情
	//
	// func Balance() int {
	// 	   mu.Lock()
	// 	   defer mu.Unlock()
	// 	   return balance
	// }
	//
	// 上面的例子里，Unlock 会在 return 语句读取完 balance 的值之后执行，所以 Balance 函数是并发安全的
	// 这带来的另一点好处是，我们再也不需要一个本地变量 b 了

	// 此外，一个 deferred Unlock 即使在临界区发生 panic 时依然会执行
	// 这对于用 recover（5.10 节）来恢复的程序来说是很重要的
	// defer 调用只会比显式地调用 Unlock 的成本高那么一点点，不过却在很大程度上保证了代码的整洁性
	// 大多数情况下对于并发程序来说，代码的整洁性比过度的优化更重要
	// 如果可能的话尽量使用 defer 来将临界区扩展到函数的结束

	// 考虑以下下面的 Withdraw 函数
	// 成功的时候，它会正确地减掉余额并返回 true
	// 但如果银行现有资金对交易来说不足，那么取款就会恢复余额，并返回 false
	//
	// -- 注意：非原子性
	// func Withdraw(amount int) bool {
	// 	   Deposit(-amount)
	// 	   if Balance() < 0 {
	// 	   	   Deposit(amount)
	// 	   	   return false // 资金不足
	// 	   }
	// 	   return true
	// }
	//
	// 函数终于给出了正确的结果，但是还有一点讨厌的副作用
	// 当过多的取款操作同时执行时，balance 可能会瞬时被减到零以下，这可能会引起一个并发的取款被不合逻辑地拒绝
	// 所以当 Bob 尝试买一辆 sports car 时，Alice 可能就没办法为她的早咖啡付款了
	// 这里的问题是，取款不是一个原子操作：
	// 它包含了三个步骤，每一步都需要去获取并释放互斥锁，但任何一次锁都不会锁上整个取款流程

	// 理想情况下，取款应该只在整个操作中获取一次互斥锁。下面这样的尝试是错误的：
	//
	// -- 注意：错误的！
	// func Withdraw(amount int) bool {
	// 	   mu.Lock()
	// 	   defer mu.Unlock()
	// 	   Deposit(-amount)
	// 	   if Balance() < 0 {
	// 	   	   Deposit(amount)
	// 	   	   return false // 资金不足
	// 	   }
	// 	   return true
	// }
	//
	// 上面这个例子中，Deposit 会调用 mu.Lock 去获取互斥锁，但因为 mutex 已经锁上了，而无法被重入
	// （译注：Go 里没有重入锁，关于重入锁的概念，请参考 Java）
	// 也就是说没法对已经锁上的 mutex 再次上锁，这会导致程序死锁，没法继续执行下去，Withdraw 会永远阻塞下去

	// 关于 Go 的互斥量不能重入这一点我们有很充分的理由
	// 互斥量的目的是为了确保共享变量在程序执行时的关键点上能够保证不变性
	// 不变性的其中之一是 “没有 goroutine 访问共享变量”
	// 但实际上对于 mutex 保护的变量来说，不变性还包括其它方面
	// 当一个 goroutine 获得了一个互斥锁时，它会断定这种不变性能够被保持
	// 其获取并保持锁期间，可能会去更新共享变量，这样不变性只是短暂地被破坏
	// 然而当其释放锁之后，它必须保证不变性已经恢复原样
	// 尽管一个可以重入的 mutex 也可以保证没有其它的 goroutine 在访问共享变量，但这种方式没法保证这些变量其它的不变性

	// 一个通用的解决方案是将一个函数分离为多个函数，比如我们把 Deposit 分成了两个：
	// 一个不导出的函数 deposit，这个函数假设锁总是会被保持并去做实际的操作；
	// 另一个是导出的函数 Deposit，这个函数会调用 deposit，但在调用前会先去获取锁
	// 同理我们可以将 Withdraw 也表示成这种形式
	//
	// func Withdraw(amount int) bool {
	// 	   mu.Lock()
	// 	   defer mu.Unlock()
	// 	   deposit(-amount)
	// 	   if balance < 0 {
	// 	   	   deposit(amount)
	// 	   	   return false
	// 	   }
	// 	   return true
	// }
	//
	// func Deposit(amount int) {
	// 	   mu.Lock()
	// 	   defer mu.Unlock()
	// 	   deposit(amount)
	// }
	//
	// func Balance() int {
	// 	   mu.Lock()
	// 	   defer mu.Unlock()
	// 	   return balance
	// }
	//
	// -- 此方法要求持有锁
	// func deposit(amount int) int {
	// 	   return balance += amount
	// }
	//
	// 当然，这里的存款 deposit 函数很小，实际上取款 Withdraw 函数不需要理会对它的调用，尽管如此，这里的表达还是表明了规则

	// 封装（6.6 节），用限制一个程序中的意外交互的方式，可以使我们获得数据结构的不变性
	// 因为某种原因，封装还帮我们获得了并发的不变性
	// 当你使用 mutex 时，确保 mutex 和其保护的变量没有被导出（即公开 public，在 Go 中，首字母大写的函数是 public 的）
	// 无论这些变量是包级变量还是 struct 的一个字段
}
