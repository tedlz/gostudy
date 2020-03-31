package memo_test

import (
	memo "gostudy/09、基于共享变量的并发/files/memo5"
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
https://golang.org, 435.792991ms, 11077 bytes
https://godoc.org, 413.531218ms, 7159 bytes
https://play.golang.org, 469.68865ms, 6013 bytes
https://gopl.io, 739.564664ms, 4154 bytes
https://golang.org, 8.996µs, 11077 bytes
https://godoc.org, 4.523µs, 7159 bytes
https://play.golang.org, 4.428µs, 6013 bytes
https://gopl.io, 3.226µs, 4154 bytes
--- PASS: Test (2.06s)
=== RUN   TestConcurrent
https://golang.org, 183.731796ms, 11077 bytes
https://golang.org, 183.010075ms, 11077 bytes
https://play.golang.org, 186.633117ms, 6013 bytes
https://play.golang.org, 186.725974ms, 6013 bytes
https://godoc.org, 219.483566ms, 7159 bytes
https://godoc.org, 218.690298ms, 7159 bytes
https://gopl.io, 363.915084ms, 4154 bytes
https://gopl.io, 363.045606ms, 4154 bytes
--- PASS: TestConcurrent (0.36s)
PASS
ok      gostudy/09、基于共享变量的并发/files/memo5      2.428s
*/

/*
命令：
go test -run=TestConcurrent -v -race
执行结果：
=== RUN   TestConcurrent
https://golang.org, 655.305788ms, 11077 bytes
https://golang.org, 654.391792ms, 11077 bytes
https://godoc.org, 840.871498ms, 7159 bytes
https://godoc.org, 857.190799ms, 7159 bytes
https://play.golang.org, 914.602491ms, 6013 bytes
https://play.golang.org, 895.276347ms, 6013 bytes
https://gopl.io, 1.180055181s, 4154 bytes
https://gopl.io, 1.200626796s, 4154 bytes
--- PASS: TestConcurrent (1.20s)
PASS
ok      gostudy/09、基于共享变量的并发/files/memo5      1.270s
*/
