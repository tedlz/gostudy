package main

// 001、竞争条件
func main() {
	// 在一个线性（就是说只有一个 goroutine）的程序中，程序的执行顺序只由程序的逻辑来决定
	// 在有两个或更多 goroutine 的程序中，每一个 goroutine 内的语句也是按照既定的顺序去执行的，
	// 但是一般情况下我们没法去知道分别位于两个 goroutine 的事件 x 和 y 的执行顺序，
	// 当我们没有办法确认一个事件是在另一个事件的前面或后面发生的话，就说明 x 和 y 这两个事件是并发的

	// 考虑一下，一个函数在线性程序中可以正确地工作
	// 如果在并发的情况下，这个函数依然可以正确工作的话，那么我们就说这个函数是并发安全的，并发安全的函数不需要额外的同步工作
	// 我们可以把这个概念概括为一个特定类型的一些方法和操作函数，
	// 对于某个类型来说，如果其所有的可访问的方法和操作都是并发安全的话，那么类型便是并发安全的

	// 在一个程序中，有非并发安全的类型的情况下，我们依然可以使这个程序并发安全
	// 并发安全的类型是例外，而不是规则，所以只有当文档中明确说明了其是并发安全的情况下，你才可以并发地去访问它
	// 我们会避免并发访问大多数的类型，
	// 无论是将变量局限在单一的一个 goroutine 内，还是用互斥条件维持更高级别的不变性都是为了这个目的

	// 导出包级别的函数一般情况下都是并发安全的
	// 由于 package 级的变量没法被限制在一个单一的 goroutine，所以修改这些变量必须使用互斥条件

	// 一个函数在并发调用时没法工作的原因太多了，比如死锁（deadlock）、活锁（livelock）和饿死（resource starvation）

	// 我们没空去讨论所有的问题，这里我们只聚焦在竞争条件上
	// 竞争条件指的是程序在多个 goroutine 执行交叉操作时，没有给出正确的结果
	// 竞争条件是一种很恶劣的场景，因为这种问题会一直潜伏在你的程序里，然后在非常少见的时候蹦出来：
	// 或许只会在很大的负载时发生，又或许是会在使用了某一个编译器、某一种平台或某一种架构的时候才会出现
	// 这些使得竞争条件带来的问题非常难以复现而且难以分析诊断

	// 传统上经常使用经济损失来为竞争条件做比喻，所以我们来看一个简单的银行账户程序
	//
	// 开设只有一个账户的银行
	// package bank
	// var balance int
	// func Deposit(amount int) { balance = balance + amount }
	// func Balance() int { return balance }
	//
	// 对于这个简单的程序而言，我们一眼就能看出，以任意顺序调用 Deposit 和 Balance 都会得到正确的结果
	// 也就是说，Balance 函数会给出之前所有存入额度之和
	// 然而，当我们并发地而不是顺序地调用这些函数的话，Balance 就再也没办法保证结果正确了
	// 考虑一下下面的两个 goroutine，其代表了一个银行联合账户的两笔交易：
	//
	// -- Alice
	// go func() {
	// 	   bank.Deposit(200)                // A1
	// 	   fmt.Println("=", bank.Balance()) // A2
	// }()
	//  -- Bob
	// go bank.Deposit(100) // B
	//
	// Alice 存了 200，同时检查她的余额，同时 Bob 存了 100
	// 因为 A1、A2 是和 B 并发执行的，我们没法预测它们发生的先后顺序
	// 直观来看，我们会认为其执行顺序有三种可能性：Alice 先、Bob 先，以及 Alice/Bob/Alice 交错执行
	// 下面的表格会展示经过每一步骤后 balance 变量的值，引号里的字符串表示余额单：
	//
	// | Alice 先   | Bob 先    | Alice/Bob/Alice |
	// | --------- | --------- | --------------- |
	// | 0         | 0         | 0               |
	// | A1 200    | B  100    | A1 200          |
	// | A2 ="200" | A1 300    | B  300          |
	// | B  300    | A2 "=300" | A2 "=300"       |
	//
	// 所有情况下的最终余额都是 300，唯一的变数是 Alice 的余额单是否包含了 Bob 交易，不过无论怎么着客户端都不会在意
	// 但事实是上面的直觉推断是错误的。第四种可能的结果是事实存在的，这种情况下 Bob 的存款会在 Alice 存款操作中间，
	// 在余额被读到（balance + amount）之后，在余额被更新之前（balance = ...），这样会导致 Bob 的交易丢失
	// 而这是因为 Alice 的存款操作 A1 实际上是两个操作的一个序列，读取然后写，可以称之为 A1r 和 A1w
	//
	// 下面是交叉时产生的问题：
	// | Data Race                        |
	// | -------------------------------- |
	// | 0                                |
	// | A1r    0    ... balance + amount |
	// | B    100                         |
	// | A1w  200     balance = ...       |
	// | A2   ="200"                      |
	//
	// 在 A1r 之后，balance + amount 会被计算为 200，所以这是 A1w 会写入的值，并不受其它存款操作的干预
	// 最终的余额是 200，此时 Bob 往银行里存的钱就丢了

	// 这个程序包含了一个特定的竞争条件，叫做数据竞争
	// 无论任何时候，只要有两个 goroutine 并发访问同一个变量，且至少其中一个是写操作的时候就会发生数据竞争

	// 如果数据竞争的对象是一个比一个机器字（译注：32 位机器上的一个字 = 4 字节）更大的类型时，事情就变得更麻烦了
	// 比如 interface、string、slice 类型都是如此
	// 下面的代码会并发地更新两个不同程度的 slice：
	//
	// var x []int
	// go func() { x = make([]int, 10) }()
	// go func() { x = make([]int, 1000000) }()
	// x[999999] = 1 // 注意：未定义的行为，可能发生内存损坏
	//
	// 最后一个语句中，x 的值是未定义的；其可能是 nil，也可能是长度为 10 的 slice，或者长度为 1000000 的 slice
	// 但是回忆一下 slice 的三个组成部分：指针（pointer）、长度（length）和容量（capacity）
	// 如果指针是从第一个 make 调用来，而长度从第二个 make 来，x 就变成了一个混合体，
	// 一个自称长度为 1000000 但实际上内部只有 10 个元素的 slice
	// 这样导致的结果是存储 999999 元素的位置会碰撞一个遥远的内存位置，这种情况下难以对值进行预测，而且 debug 也会变成噩梦
	// 这种语义雷区被称为未定义行为，对 C 程序员来说应该很熟悉；幸运的是在 Go 里面造成的麻烦要比在 C 里面小的多

	// 尽管并发程序的概念让我们知道并发并不是简单的语句交叉执行。我们将会在 9.4 节中看到，数据竞争可能会有奇怪的结果
	// 有些程序员也会提出理由来允许数据竞争，比如 “互斥条件代价太高”、“这个逻辑只用来做 logging”、“我不介意丢失一些消息” 等
	// 因为在他们的编译器或者平台上很少遇到问题，可能给了他们错误的信心
	// 一个好的经验法则是，根本就没有什么所谓的良性数据竞争，所以我们一定要避免数据竞争

	// 有三种方式可以避免数据竞争：

	// 第一种方法是不要写变量。
	// 考虑一下下面的 map，会被懒填充，也就是说每个 key 在被第一次请求到的时候才会去填值
	// 如果 Icon 是被顺序调用的话，这个程序会工作很正常；但如果 Icon 被并发调用，那么对于这个 map 来说就会存在数据竞争
	//
	// var icons = make(map[string]image.Image)
	// func loadIcon(name string) image.Image
	// -- 注意：非并发安全
	// func Icon(name string) image.Image {
	// 	   icon, ok := icons[name]
	// 	   if !ok {
	// 		   icon = loadIcon(name)
	// 		   icons[name] = icon
	// 	   }
	// 	   return icon
	// }
	//
	// 反之，如果我们在创建 goroutine 之前的初始化阶段，就初始化了 map 中的所有条目并且再也不去修改它们，
	// 那么任意数量的 goroutine 并发访问 Icon 都是安全的，因为每一个 goroutine 都只是去读取而已
	//
	// var icons = map[string]image.Image{
	// 	"spades.png":   loadIcon("spades.png"),
	// 	"hearts.png":   loadIcon("hearts.png"),
	// 	"diamonds.png": loadIcon("diamonds.png"),
	// 	"clubs.png":    loadIcon("clubs.png"),
	// }
	// -- 线程安全
	// func Icon(name string) image.Image { return icons[name] }
	//
	// 上面的例子里 icons 变量在包初始化阶段就已经被赋值了，包的初始化是在程序 main 函数开始执行之前就完成了的
	// 只要初始化完成了，icons 就再也不会被修改。永不修改或不可变的数据结构本质上就是并发安全的，不需要做同步
	// 不过显然我们没法用这种方法，因为 update 操作是必要的操作，尤其对于银行账户来说

	// 第二种避免数据竞争的方法是，避免从多个 goroutine 访问变量。
	// 这也是前一章中大多数程序所采用的方法
	// 例如前面的《并发的 Web 爬虫（8.6 节）》的 main goroutine 是唯一一个能够访问 seen map 的 goroutine，
	// 而《聊天服务器（8.10 节）》中的 broadcaster goroutine 是唯一一个能够访问 clients map 的 goroutine
	// 这些变量都被限定在了一个单独的 goroutine 中

	// 由于其它的 goroutine 不能够直接访问变量，它们只能使用一个 channel 来发送给指定的 goroutine 请求来查询更新变量
	// 这也就是 Go 的口头禅 “不要使用共享数据来通信；使用通信来共享数据”
	// 一个提供对一个指定的变量通过 channel 来请求的 goroutine 叫做这个变量的监控（monitor）goroutine
	// 例如 broadcaster goroutine 会监控 clients map 的全部访问

	// 下面是一个重写的银行例子（见 files/bank1），例子中 balance 变量被限制在了监控 goroutine 中，名为 teller

	// 下面的例子中，Cakes 会被严格地顺序访问，先是 baker goroutine，然后是 icer goroutine
	//
	// type Cake struct{ state string }
	// func baker(cooked chan<- *Cake) {
	// 	   for {
	// 		   cake := new(Cake)
	// 		   cake.state = "cooked"
	// 		   cooked <- cake // baker 再也不碰这个 cake
	// 	   }
	// }

	// func icer(iced chan<- *Cake, cooked <-chan *Cake) {
	// 	   for cake := range cooked {
	// 		   cake.state = "iced"
	// 		   iced <- cake // icer 再也不碰这个 cake
	// 	   }
	// }
	//

	// 第三种避免数据竞争的方法是允许很多 goroutine 去访问变量，但是在同一个时刻最多只有一个 goroutine 在访问
	// 这种方式被称为互斥，在下一节来讨论这个问题
}
