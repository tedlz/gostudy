package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// 020、Recover 捕获异常
// go run 020_recover.go https://golang.org
// 输出：
// multiple title elements
// exit status 1
//
// go run 020_recover.go https://google.cn
// 输出：
// Google
func main() {
	// 通常来说，不应该对 panic 异常做任何处理，但有时也许我们可以从异常中恢复，至少我们可以在程序崩溃前做些操作
	// 举个例子，当 web 服务器遇到不可预料的严重问题时，在崩溃前应将所有的连接关闭，如果不做任何处理会使客户端一直处于等待状态
	// 如果 web 服务器还在开发阶段，可将异常信息反馈到客户端，帮助调试

	// 如果在延迟函数调用中调用了内置函数 recover，并且定义了该 defer 语句的函数发生了 panic 异常，
	// recover 会使程序从 panic 中恢复，并返回 panic value。导致 panic 异常的函数不会继续运行，但能正常返回
	// 在未发生 panic 时调用 recover，recover 会返回 nil

	// 我们以语言解析器为例，说明 recover 的使用场景
	// 考虑到语言解析器的复杂性，即使某个语言解析器目前工作正常，也无法肯定它没有漏洞
	// 因此，当某个异常出现时，我们不会选择让解析器崩溃，而是会将 panic 异常当做普通的解析错误，并附加额外信息提醒用户报告此错误
	// func Parse(input string) (s *Syntax, err error) {
	// 	   defer func () {
	// 		   if p := recover(); p != nil {
	// 			   err = fmt.Errorf("internal error: %v", p)
	// 		   }
	// 	   }()
	//     ...parser
	// }
	// 延迟函数调用（deferred 函数）帮助 Parse 从 panic 中恢复
	// 在 deferred 函数内部，panic value 被附加到错误信息中，并用 err 变量接收错误信息，返回给调用者
	// 我们也可以通过调用 runtime.Stack 往错误信息中添加完整的堆栈调用信息

	// 不加区分的恢复所有的 panic 异常，不是可取的做法，因为在 panic 之后，无法保证包级变量的状态仍然和我们预期的一致
	// 比如，对数据结构的一次重要更新没有被完整完成、文件或者网络连接没有被关闭、获得的锁没有被释放
	// 此外，如果写日志时产生的 panic 被不加区分的恢复，可能会导致漏洞被忽略

	// 虽然把对 panic 的处理都集中在一个包下，有助于简化对复杂和不可预料问题的处理，
	// 但作为被广泛遵守的规范，你不应该试图去恢复其它包引起的 panic
	// 公有的 API 应该将函数的运行失败作为 error 返回，而不是 panic
	// 同样的，你也不应该恢复一个由他人开发的函数引起的 panic，比如说调用者传入的回调函数，因为你无法保证这样做是安全的

	// 有时我们很难完全遵循规范，举个例子，net/http 包中提供了一个 web 服务器，将收到的请求分发给用户提供的处理函数
	// 很显然，我们不能因为某个处理函数引发的 panic 异常而杀掉整个进程
	// web 服务器遇到处理函数导致的 panic 时会调用 recover，输出堆栈信息，继续运行
	// 这样的做法在实战中很便捷，但也会引起资源泄露，或是因为 recover 操作，导致其它问题

	// 基于以上原因，安全的做法是有选择性的 recover
	// 换句话说，只恢复应该被恢复的 panic 异常，此外，这些异常所占的比例应该尽可能的低
	// 为了标识某个 panic 是否应该被恢复，我们可以将 panic value 设置成特殊类型，
	// 在 recover 时对 panic value 进行检查，如果发现是特殊类型，就将这个 panic 作为 error 处理，
	// 如果不是，则按照正常的 panic 处理，下面的例子中我们会看到这种方式
	doc, err := parse(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	title, err := soleTitle(doc)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(title)

	// 有些情况下，我们无法恢复，某些致命错误会导致 Go 在运行时终止程序，如内存不足
}

func parse(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parse %s as HTML: %v", url, err)
	}
	return doc, nil
}

// deferred 函数调用 recover，并检查 panic value
// 当 panic value 是 bailout 类型时，deferred 函数生成一个 error 返回给调用者
// 当 panic value 是其它 non-nil 值时，表示发生了未知的 panic 异常，
// deferred 函数将调用 panic 函数并将当前的 panic value 作为参数传入，此时等同于 recover 没有做任何操作
// （请注意：在例子中对可预期的错误采用了 panic，这违反了之前的建议，我们在此只是想向读者演示这种机制）
func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}
	defer func() {
		switch p := recover(); p {
		case nil:
		case bailout{}:
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p)
		}
	}()
	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if title != "" {
				panic(bailout{})
			}
			title = n.FirstChild.Data
		}
	}, nil)
	if title == "" {
		return "", fmt.Errorf("no title element")
	}
	return title, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
