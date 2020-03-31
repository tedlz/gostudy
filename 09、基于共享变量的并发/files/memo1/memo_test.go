package memo_test

import (
	memo "gostudy/09、基于共享变量的并发/files/memo1"
	"gostudy/09、基于共享变量的并发/files/memotest"
	"testing"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Sequential(t, m)
}

// 注意：不是并发安全的，测试失败！
func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Concurrent(t, m)
}

/*
命令：
go test -v
执行结果：
=== RUN   Test
https://golang.org, 456.236419ms, 11077 bytes
https://godoc.org, 391.116987ms, 7159 bytes
https://play.golang.org, 472.716751ms, 6013 bytes
https://gopl.io, 728.935826ms, 4154 bytes
https://golang.org, 1.729µs, 11077 bytes
https://godoc.org, 506ns, 7159 bytes
https://play.golang.org, 476ns, 6013 bytes
https://gopl.io, 376ns, 4154 bytes
--- PASS: Test (2.05s)
=== RUN   TestConcurrent
https://play.golang.org, 183.455094ms, 6013 bytes
https://golang.org, 190.565405ms, 11077 bytes
https://golang.org, 190.615849ms, 11077 bytes
https://godoc.org, 211.151938ms, 7159 bytes
https://godoc.org, 224.137141ms, 7159 bytes
https://play.golang.org, 280.114071ms, 6013 bytes
https://gopl.io, 352.995857ms, 4154 bytes
https://gopl.io, 356.540338ms, 4154 bytes
--- PASS: TestConcurrent (0.36s)
PASS
ok      gostudy/09、基于共享变量的并发/files/memo1      2.411s
*/

/*
命令：
go test -run=TestConcurrent -v -race
执行结果：
=== RUN   TestConcurrent
https://play.golang.org, 671.492084ms, 6013 bytes
==================
WARNING: DATA RACE
Write at 0x00c00009ed80 by goroutine 9:
  runtime.mapassign_faststr()
      /usr/local/go/src/runtime/map_faststr.go:202 +0x0
  gostudy/09%e3%80%81%e5%9f%ba%e4%ba%8e%e5%85%b1%e4%ba%ab%e5%8f%98%e9%87%8f%e7%9a%84%e5%b9%b6%e5%8f%91/files/memo1.(*Memo).Get()
      /data/go/src/gostudy/09、基于共享变量的并发/files/memo1/memo.go:27 +0x1ce
  gostudy/09%e3%80%81%e5%9f%ba%e4%ba%8e%e5%85%b1%e4%ba%ab%e5%8f%98%e9%87%8f%e7%9a%84%e5%b9%b6%e5%8f%91/files/memotest.Concurrent.func1()
      /data/go/src/gostudy/09、基于共享变量的并发/files/memotest/memotest.go:72 +0xde

Previous write at 0x00c00009ed80 by goroutine 14:
  runtime.mapassign_faststr()
      /usr/local/go/src/runtime/map_faststr.go:202 +0x0
  gostudy/09%e3%80%81%e5%9f%ba%e4%ba%8e%e5%85%b1%e4%ba%ab%e5%8f%98%e9%87%8f%e7%9a%84%e5%b9%b6%e5%8f%91/files/memo1.(*Memo).Get()
      /data/go/src/gostudy/09、基于共享变量的并发/files/memo1/memo.go:27 +0x1ce
  gostudy/09%e3%80%81%e5%9f%ba%e4%ba%8e%e5%85%b1%e4%ba%ab%e5%8f%98%e9%87%8f%e7%9a%84%e5%b9%b6%e5%8f%91/files/memotest.Concurrent.func1()
      /data/go/src/gostudy/09、基于共享变量的并发/files/memotest/memotest.go:72 +0xde

Goroutine 9 (running) created at:
  gostudy/09%e3%80%81%e5%9f%ba%e4%ba%8e%e5%85%b1%e4%ba%ab%e5%8f%98%e9%87%8f%e7%9a%84%e5%b9%b6%e5%8f%91/files/memotest.Concurrent()
      /data/go/src/gostudy/09、基于共享变量的并发/files/memotest/memotest.go:69 +0x10c
  gostudy/09%e3%80%81%e5%9f%ba%e4%ba%8e%e5%85%b1%e4%ba%ab%e5%8f%98%e9%87%8f%e7%9a%84%e5%b9%b6%e5%8f%91/files/memo1_test.TestConcurrent()
      /data/go/src/gostudy/09、基于共享变量的并发/files/memo1/memo_test.go:19 +0xd9
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:992 +0x1eb

Goroutine 14 (finished) created at:
  gostudy/09%e3%80%81%e5%9f%ba%e4%ba%8e%e5%85%b1%e4%ba%ab%e5%8f%98%e9%87%8f%e7%9a%84%e5%b9%b6%e5%8f%91/files/memotest.Concurrent()
      /data/go/src/gostudy/09、基于共享变量的并发/files/memotest/memotest.go:69 +0x10c
  gostudy/09%e3%80%81%e5%9f%ba%e4%ba%8e%e5%85%b1%e4%ba%ab%e5%8f%98%e9%87%8f%e7%9a%84%e5%b9%b6%e5%8f%91/files/memo1_test.TestConcurrent()
      /data/go/src/gostudy/09、基于共享变量的并发/files/memo1/memo_test.go:19 +0xd9
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:992 +0x1eb
==================
https://golang.org, 683.501458ms, 11077 bytes
==================
WARNING: DATA RACE
Write at 0x00c0000ae0a8 by goroutine 16:
  gostudy/09%e3%80%81%e5%9f%ba%e4%ba%8e%e5%85%b1%e4%ba%ab%e5%8f%98%e9%87%8f%e7%9a%84%e5%b9%b6%e5%8f%91/files/memo1.(*Memo).Get()
      /data/go/src/gostudy/09、基于共享变量的并发/files/memo1/memo.go:27 +0x1ec
  gostudy/09%e3%80%81%e5%9f%ba%e4%ba%8e%e5%85%b1%e4%ba%ab%e5%8f%98%e9%87%8f%e7%9a%84%e5%b9%b6%e5%8f%91/files/memotest.Concurrent.func1()
      /data/go/src/gostudy/09、基于共享变量的并发/files/memotest/memotest.go:72 +0xde

Previous write at 0x00c0000ae0a8 by goroutine 9:
  gostudy/09%e3%80%81%e5%9f%ba%e4%ba%8e%e5%85%b1%e4%ba%ab%e5%8f%98%e9%87%8f%e7%9a%84%e5%b9%b6%e5%8f%91/files/memo1.(*Memo).Get()
      /data/go/src/gostudy/09、基于共享变量的并发/files/memo1/memo.go:27 +0x1ec
  gostudy/09%e3%80%81%e5%9f%ba%e4%ba%8e%e5%85%b1%e4%ba%ab%e5%8f%98%e9%87%8f%e7%9a%84%e5%b9%b6%e5%8f%91/files/memotest.Concurrent.func1()
      /data/go/src/gostudy/09、基于共享变量的并发/files/memotest/memotest.go:72 +0xde

Goroutine 16 (running) created at:
  gostudy/09%e3%80%81%e5%9f%ba%e4%ba%8e%e5%85%b1%e4%ba%ab%e5%8f%98%e9%87%8f%e7%9a%84%e5%b9%b6%e5%8f%91/files/memotest.Concurrent()
      /data/go/src/gostudy/09、基于共享变量的并发/files/memotest/memotest.go:69 +0x10c
  gostudy/09%e3%80%81%e5%9f%ba%e4%ba%8e%e5%85%b1%e4%ba%ab%e5%8f%98%e9%87%8f%e7%9a%84%e5%b9%b6%e5%8f%91/files/memo1_test.TestConcurrent()
      /data/go/src/gostudy/09、基于共享变量的并发/files/memo1/memo_test.go:19 +0xd9
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:992 +0x1eb

Goroutine 9 (finished) created at:
  gostudy/09%e3%80%81%e5%9f%ba%e4%ba%8e%e5%85%b1%e4%ba%ab%e5%8f%98%e9%87%8f%e7%9a%84%e5%b9%b6%e5%8f%91/files/memotest.Concurrent()
      /data/go/src/gostudy/09、基于共享变量的并发/files/memotest/memotest.go:69 +0x10c
  gostudy/09%e3%80%81%e5%9f%ba%e4%ba%8e%e5%85%b1%e4%ba%ab%e5%8f%98%e9%87%8f%e7%9a%84%e5%b9%b6%e5%8f%91/files/memo1_test.TestConcurrent()
      /data/go/src/gostudy/09、基于共享变量的并发/files/memo1/memo_test.go:19 +0xd9
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:992 +0x1eb
==================
https://golang.org, 684.946146ms, 11077 bytes
https://godoc.org, 683.452239ms, 7159 bytes
https://godoc.org, 689.321533ms, 7159 bytes
https://play.golang.org, 741.77646ms, 6013 bytes
https://gopl.io, 1.013736586s, 4154 bytes
https://gopl.io, 1.023209596s, 4154 bytes
    TestConcurrent: testing.go:906: race detected during execution of test
--- FAIL: TestConcurrent (1.03s)
    : testing.go:906: race detected during execution of test
FAIL
exit status 1
FAIL    gostudy/09、基于共享变量的并发/files/memo1      1.038s
*/
