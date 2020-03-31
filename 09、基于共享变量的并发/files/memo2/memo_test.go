package memo_test

import (
	memo "gostudy/09、基于共享变量的并发/files/memo2"
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
https://golang.org, 412.616932ms, 11077 bytes
https://godoc.org, 509.605011ms, 7159 bytes
https://play.golang.org, 421.174278ms, 6013 bytes
https://gopl.io, 718.672394ms, 4154 bytes
https://golang.org, 1.212µs, 11077 bytes
https://godoc.org, 260ns, 7159 bytes
https://play.golang.org, 187ns, 6013 bytes
https://gopl.io, 233ns, 4154 bytes
--- PASS: Test (2.06s)
=== RUN   TestConcurrent
https://godoc.org, 314.417641ms, 7159 bytes
https://golang.org, 507.608668ms, 11077 bytes
https://gopl.io, 861.129179ms, 4154 bytes
https://play.golang.org, 1.159111376s, 6013 bytes
https://godoc.org, 1.159148373s, 7159 bytes
https://golang.org, 1.159223255s, 11077 bytes
https://gopl.io, 1.159252253s, 4154 bytes
https://play.golang.org, 1.15927606s, 6013 bytes
--- PASS: TestConcurrent (1.16s)
PASS
ok      gostudy/09、基于共享变量的并发/files/memo2      3.228s
*/

/*
命令：
go test -run=TestConcurrent -v -race
执行结果：
=== RUN   TestConcurrent
https://golang.org, 686.527752ms, 11077 bytes
https://godoc.org, 1.204628759s, 7159 bytes
https://play.golang.org, 1.72460784s, 6013 bytes
https://golang.org, 1.724617968s, 11077 bytes
https://gopl.io, 2.519667974s, 4154 bytes
https://godoc.org, 2.519450911s, 7159 bytes
https://play.golang.org, 2.519577108s, 6013 bytes
https://gopl.io, 2.519700733s, 4154 bytes
--- PASS: TestConcurrent (2.52s)
PASS
ok      gostudy/09、基于共享变量的并发/files/memo2      2.556s
*/
