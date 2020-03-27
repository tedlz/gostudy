package main

// 005、sync.Once 初始化
func main() {

}

// 如果初始化成本比较大的话，那么将初始化延迟到需要的时候再去做就是一个比较好的选择
// 如果在程序启动的时候就去做这类的初始化的话，会增加程序的启动时间，
// 并且因为执行的时候可能也并不需要这些变量所以实际上有一些浪费
// 比如我们在本章早些时候看到的变量
//
// var icons map[string]image.Image

// 这个版本的 Icon 用到了懒初始化（lazy initialization）
//
// func loadIcons() {
// 	   icons = map[string]image.Image{
// 	   	   "spades.png":   loadIcon("spades.png"),
// 	   	   "hearts.png":   loadIcon("hearts.png"),
// 	   	   "diamonds.png": loadIcon("diamonds.png"),
// 	   	   "clubs.png":    loadIcon("clubs.png"),
// 	   }
// }
//
// -- 注意：非并发安全！
// func Icon(name string) image.Image {
// 	   if icons == nil {
// 	   	   loadIcons() // 一次性初始化
// 	   }
// 	   return icons[name]
// }

// 如果一个变量只被一个单独的 goroutine 所访问的话，我们可以使用上面的这种模板，但这种模板在 Icon 被并发调用时并不安全
// 就像前面银行的那个 Deposit 存款函数一样，Icon 函数也是由多个步骤组成的：
// 首先测试 icons 是否为空，然后 load 这些 icons，之后将 icons 更新为一个非空的值
// 直觉告诉我们最差的情况是 loadIcons 函数被多次访问会带来数据竞争
// 当第一个 goroutine 在忙着 loading 这些 icons 的时候，
// 另一个 goroutine 进入了 Icon 函数，发现变量是 nil，然后也会调用 loadIcons 函数

// 不过这种直觉是错误的（我们希望你从现在开始能够构建自己对并发的直觉，也就是说对并发的直觉总是不能被信任的）
// 回忆一下 9.4 节，因为缺少显式的同步，编译器和 CPU 是可以随意地去更改访问内存的指令顺序，
// 以任意方式，只要保证每一个 goroutine 自己的执行顺序一致
// 其中一种可能 loadIcons 的语句重排是下面这样，它会在填写 icons 变量的值之前先用一个空 map 来初始化 icons 变量
//
// func loadIcons() {
// 	   icons = make(map[string]image.Image)
// 	   icons["spades.png"] = loadIcon("spades.png")
// 	   icons["hearts.png"] = loadIcon("hearts.png")
// 	   icons["diamonds.png"] = loadIcon("diamonds.png")
// 	   icons["clubs"] = loadIcon("clubs")
// }

// 因此，一个 goroutine 在检查 icons 是非空时，也并不能就假设这个变量的初始化流程已经走完了

// 最简单且正确的保证所有 goroutine 能够观察到 loadIcons 效果的方式，是用一个 mutex 来同步检查
//
// var mu sync.Mutex // 守护 icons
// var icons map[string]image.Image
// -- 并发安全
// func Icon(name string) image.Image {
// 	   mu.Lock()
// 	   defer mu.Unlock()
// 	   if icons == nil {
// 	   	   loadIcons()
// 	   }
// 	   return icons[name]
// }

// 然而使用互斥访问 icons 的代价就是没有办法对该变量进行并发访问，即使变量已经被初始化完毕且再也不会变动
// 这里我们可以引入一个允许多读的锁：
//
// var mu sync.RWMutex
// var icons map[string]image.Image
// -- 线程安全
// func Icon(name string) image.Image {
// 	   mu.RLock()
// 	   if icons != nil {
// 	   	   icon := icons[name]
// 	   	   mu.RUnlock()
// 	   	   return icon
// 	   }
// 	   mu.RUnlock()
//     -- 获取排他锁
// 	   mu.Lock()
// 	   if icons == nil { // 注意：必须重检查 nil
// 	   	   loadIcons()
// 	   }
// 	   icon := icons[name]
// 	   mu.Unlock()
// 	   return icon
// }

// 上面的代码有两个临界区
// goroutine 首先会获取一个读锁，查询 map，然后释放锁。如果条目被找到了（一般情况下），那么会直接返回
// 如果没有找到，那 goroutine 会获取一个写锁
// 不释放共享锁的话，也没有任何办法来将一个共享锁升级为一个互斥锁，
// 所以我们必须重新检查 icons 变量是否为 nil，以防止在执行这一段代码的时候，icons 变量已经被其它 goroutine 初始化过了

// 上面的模板使我们的程序能够更好的并发，但是有一点复杂且容易出错
// 幸运的是，sync 包为我们提供了一个专门的方案来解决这种一次性初始化的问题：sync.Once
// 概念上来讲，一次性的初始化需要一个互斥量 mutex 和一个 boolean 变量来记录初始化是不是已经完成了，
// 互斥量用来保护 boolean 变量和客户端数据结构
// Do 这个唯一的方法需要接收初始化函数作为其参数
// 让我们用 sync.Once 来简化前面的 Icon 函数吧：
//
// var loadIconsOnce sync.Once
// var icons map[string]image.Image
// -- 线程安全
// func Icon(name string) image.Image {
// 	   loadIconsOnce.Do(loadIcons)
// 	   return icons[name]
// }

// 每一次对 Do（loadIcons）的调用都会锁定 mutex，并会检查 boolean 变量
// 在第一次调用时，boolean 变量的值是 false，Do 会调用 loadIcons 并会将 boolean 变量设置为 true
// 随后的调用什么都不会做，但是 mutex 同步会保证 loadIcons 对内存（这里指 icons 变量）产生的效果能够对所有 goroutine 可见
// 用这种方式来使用 sync.Once 的话，我们能够避免在变量被构建完成之前和其它 goroutine 共享该变量
