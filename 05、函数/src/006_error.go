package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// 006、错误及常用处理方式
// go run 006_error.go bad.gopl.io
func main() {
	// 在 Go 中一部分函数总是能成功的运行
	// 比如 strings.Contains 和 strconv.FormatBool 函数，对各种可能的输入都做了良好的处理，使运行时几乎不会失败
	// 除非遇到灾难性和不可预料的情况，比如运行时的内存溢出，导致这种错误的原因很复杂，难以处理，从错误中恢复的可能性也很低
	// 还有一部分函数只要输入的参数满足一定条件，也能保证运行成功
	// 比如 time.Date 函数，该函数将年月日等参数构造成 time.Time 对象，除非最后一个参数（时区）是 nil
	// 这种情况下会引发 panic 异常，panic 是来自被调函数的信号，表示发生了某个已知的 bug
	// 一个良好的程序永远不该发生 panic 异常

	// 对于大部分函数而言，永远无法确保能否成功运行，因为错误的原因超出了程序员的控制
	// 举个例子，任何 I/O 操作都会面临出现错误的可能，即使是简单的读写
	// 因此，当本该可信的操作出乎意料的失败后，我们必须弄清楚导致失败的原因

	// 在 Go 的错误处理中，错误是软件包 API 和应用程序用户界面的一个重要组成部分，程序运行失败仅被认为是几个预期的结果之一
	// 对于那些将运行失败看做是预期结果的函数，它们会返回一个额外的返回值，通常是最后一个，来传递错误信息
	// 如果导致失败的原因只有一个，额外的返回值可以是一个布尔值，通常被命名为 ok
	// 比如 cache.Lookup 失败的唯一原因是 key 不存在，那么代码可以按照下面的方式组织：
	// value, ok := cache.Lookup(key)
	// if !ok {
	//     ... cache[key] does not exist
	// }

	// 通常导致失败的原因不止一种，尤其是对 I/O 操作而言，用户需要了解更多的错误信息
	// 因此，额外的返回值不再是简单的 bool 类型，而是 error 类型
	// 内置的 error 是接口类型，error 类型可能是 nil 和 non-nil
	// nil 意味着函数运行成功，non-nil 表示失败
	// 对于 non-nil 的 error 类型，我们可以通过调用 error 的 Error 函数或者输出函数获得字符串类型的错误信息：
	// fmt.Println(err)
	// fmt.Printf("%v", err)
	// 通常，当函数返回 non-nil 的 error 时，其它的返回值是未定义的（undefined），这些未定义的返回值应该被忽略
	// 然而，有少部分函数在发生错误时，仍然会返回一些有用的返回值
	// 比如，当读取文件发生错误时，Read 函数会返回可以读取的字节数以及错误信息
	// 对于这种情况，正确的处理方式应该是先处理这些不完整的数据，再处理错误
	// 因此对于函数的返回值要有清晰的说明，便于他人使用

	// 在 Go 中，函数运行失败时会返回错误信息，这些错误信息被认为是一种预期的值而非异常（exception）
	// 这使得 Go 有别于那些将函数运行失败看做是异常的语言
	// 虽然 Go 有各种异常机制，但这些机制仅被使用在处理那些未被预料到的错误，即 bug，而不是那些在健壮程序中应该被避免的程序错误
	// Go 这样设计的原因是由于对于某个应该在控制流程中处理的错误而言，将这个错误以异常的形式抛出会混乱对错误的描述
	// 当某个程序错误被当做异常处理后，这些错误会将堆栈根据信息返回给终端用户，这些信息复杂且无用，无法帮助定位错误
	// 因此，Go 使用控制流机制（例如 if 和 return）处理异常，这使得编码人员能更多的关注错误处理

	// 错误处理策略
	// 常用的五种方式：
	// ==========
	// 1、传播错误（这意味着函数中某个程序的失败，就会变成该函数的失败）
	// 以 004_findlinks2.go 的 findLinks 函数作例，如果函数中的 http.Get 调用失败，会直接将错误返回给调用者
	// resp, err := http.Get(url)
	// if err != nil{
	// 	   return nil, err
	// }
	// 当对 html.Parse 调用失败时，findLinks 不会直接返回 html.Parse 的错误，因为缺少两条重要信息
	// doc, err := html.Parse(resp.Body)
	// resp.Body.Close()
	// if err != nil {
	// 	   return nil, fmt.Errorf("parsing %s as HTML: %v", url,err)
	// }
	// (1) 错误发生在解析器
	// (2) url 已经被解析
	// 这些信息有助于错误的处理，findLinks 会构造新的错误信息返回给调用者

	// ==========
	// 2、重新尝试失败的操作（如果错误的发生是偶然的，或由不可预知的问题导致）
	// 重试时，我们需要限制重试的时间间隔或重试的次数，防止无限制的重试，代码见 WaitForServer

	// ==========
	// 3、输出错误信息并结束程序（如果错误发生后，程序无法继续运行）
	// 需要注意的是，这种策略只应在 main 中执行
	// 对库函数而言，应仅向上传播错误，除非该错误意味着程序内部包含不一致性，即遇到了 bug，才能在库函数中结束程序
	url := os.Args[1]
	// if err := WaitForServer(url); err != nil {
	// 	   fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
	// 	   os.Exit(1)
	// }
	// 输出：
	// 	2020/01/07 10:54:58 server not responding (Head bad.gopl.io: unsupported protocol scheme ""); retrying...
	// 2020/01/07 10:54:59 server not responding (Head bad.gopl.io: unsupported protocol scheme ""); retrying...
	// 2020/01/07 10:55:01 server not responding (Head bad.gopl.io: unsupported protocol scheme ""); retrying...
	// 2020/01/07 10:55:05 server not responding (Head bad.gopl.io: unsupported protocol scheme ""); retrying...
	// 2020/01/07 10:55:13 server not responding (Head bad.gopl.io: unsupported protocol scheme ""); retrying...
	// 2020/01/07 10:55:29 server not responding (Head bad.gopl.io: unsupported protocol scheme ""); retrying...
	// Site is down: server bad.gopl.io failed to respond after 1m0s
	// exit status 1
	// log.Fatalf 可以更简洁的代码达到与上文相同的效果，log 中的所有函数，都会在错误信息之前输出时间信息
	if err := WaitForServer(url); err != nil {
		log.Fatalf("Site is down: %v\n", err)
		os.Exit(1)
	}
	// 输出：
	// 	2020/01/07 10:56:45 server not responding (Head bad.gopl.io: unsupported protocol scheme ""); retrying...
	// 2020/01/07 10:56:46 server not responding (Head bad.gopl.io: unsupported protocol scheme ""); retrying...
	// 2020/01/07 10:56:48 server not responding (Head bad.gopl.io: unsupported protocol scheme ""); retrying...
	// 2020/01/07 10:56:52 server not responding (Head bad.gopl.io: unsupported protocol scheme ""); retrying...
	// 2020/01/07 10:57:00 server not responding (Head bad.gopl.io: unsupported protocol scheme ""); retrying...
	// 2020/01/07 10:57:16 server not responding (Head bad.gopl.io: unsupported protocol scheme ""); retrying...
	// 2020/01/07 10:57:48 Site is down: server bad.gopl.io failed to respond after 1m0s
	// exit status 1

	// 我们可以设置 log 的前缀信息屏蔽时间信息，一般而言，前缀信息会被设置成命令名
	// log.SetPrefix("wait: ")
	// log.SetFlags(0)

	// ==========
	// 4、输出错误信息，不中断程序运行
	if err := Ping(); err != nil {
		log.Printf("ping failed: %v; networking disabled", err)
	}
	// 或者标准错误流输出错误信息
	if err := Ping(); err != nil {
		fmt.Fprintf(os.Stderr, "ping failed: %v; networking disabled", err)
	}
	// log 包中的所有函数会为没有换行符的字符串增加换行符

	// ==========
	// 5、忽略错误
	// dir, err := ioutil.TempDir("", "scratch")
	// if err != nil {
	// 	return fmt.Errorf("failed to create temp dir: %v", err)
	// }
	// ...use temp dir…
	// os.RemoveAll(dir) // ignore errors; $TMPDIR is cleaned periodically
	// 尽管 os.RemoveAll 会失败，但上面的例子并没有做错误处理，因为操作系统会定期清理临时目录
	// 正因如此，虽然程序没有处理错误，但程序的逻辑不会因此受到影响
	// 我们应该在每次函数调用后，都养成考虑错误处理的习惯
	// 当你决定忽略某个错误时，你应该清晰的记录下你的意图
}

// WaitForServer *
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

// Ping *
func Ping() error {
	return fmt.Errorf("ping error")
}
