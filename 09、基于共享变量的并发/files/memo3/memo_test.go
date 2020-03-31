package memo_test

import (
	memo "gostudy/09、基于共享变量的并发/files/memo3"
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
https://golang.org, 430.254996ms, 11077 bytes
https://godoc.org, 979.962716ms, 7159 bytes
https://play.golang.org, 438.281281ms, 6013 bytes
https://gopl.io, 877.217864ms, 4154 bytes
https://golang.org, 906ns, 11077 bytes
https://godoc.org, 914ns, 7159 bytes
https://play.golang.org, 762ns, 6013 bytes
https://gopl.io, 495ns, 4154 bytes
--- PASS: Test (2.73s)
=== RUN   TestConcurrent
https://golang.org, 179.157301ms, 11077 bytes
https://golang.org, 185.268285ms, 11077 bytes
https://play.golang.org, 277.226289ms, 6013 bytes
https://play.golang.org, 281.954389ms, 6013 bytes
https://godoc.org, 319.011594ms, 7159 bytes
https://gopl.io, 340.736148ms, 4154 bytes
https://gopl.io, 341.761833ms, 4154 bytes
https://godoc.org, 348.250865ms, 7159 bytes
--- PASS: TestConcurrent (0.35s)
PASS
ok      gostudy/09、基于共享变量的并发/files/memo3      3.079s
*/

/*
命令：
go test -run=TestConcurrent -v -race
执行结果：
=== RUN   TestConcurrent
https://golang.org, 631.326487ms, 11077 bytes
https://play.golang.org, 620.708695ms, 6013 bytes
https://golang.org, 628.35805ms, 11077 bytes
https://godoc.org, 662.243534ms, 7159 bytes
https://godoc.org, 659.104454ms, 7159 bytes
https://play.golang.org, 747.629454ms, 6013 bytes
https://gopl.io, 961.063956ms, 4154 bytes
https://gopl.io, 953.452516ms, 4154 bytes
--- PASS: TestConcurrent (0.97s)
PASS
ok      gostudy/09、基于共享变量的并发/files/memo3      0.989s
*/
