package memo_test

import (
	memo "gostudy/09、基于共享变量的并发/files/memo4"
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
https://golang.org, 457.849642ms, 11077 bytes
https://godoc.org, 427.238612ms, 7159 bytes
https://play.golang.org, 572.867886ms, 6013 bytes
https://gopl.io, 749.680727ms, 4154 bytes
https://golang.org, 1.328µs, 11077 bytes
https://godoc.org, 832ns, 7159 bytes
https://play.golang.org, 915ns, 6013 bytes
https://gopl.io, 1.025µs, 4154 bytes
--- PASS: Test (2.21s)
=== RUN   TestConcurrent
https://play.golang.org, 194.043713ms, 6013 bytes
https://play.golang.org, 194.184387ms, 6013 bytes
https://godoc.org, 309.036603ms, 7159 bytes
https://godoc.org, 308.939111ms, 7159 bytes
https://gopl.io, 336.438403ms, 4154 bytes
https://gopl.io, 336.535976ms, 4154 bytes
https://golang.org, 479.377886ms, 11077 bytes
https://golang.org, 479.369195ms, 11077 bytes
--- PASS: TestConcurrent (0.48s)
PASS
ok      gostudy/09、基于共享变量的并发/files/memo4      2.695s
*/

/*
命令：
go test -run=TestConcurrent -v -race
执行结果：
=== RUN   TestConcurrent
https://play.golang.org, 611.979811ms, 6013 bytes
https://play.golang.org, 604.972733ms, 6013 bytes
https://golang.org, 623.433327ms, 11077 bytes
https://golang.org, 624.253517ms, 11077 bytes
https://godoc.org, 752.495566ms, 7159 bytes
https://godoc.org, 745.014887ms, 7159 bytes
https://gopl.io, 1.051700496s, 4154 bytes
https://gopl.io, 1.045024294s, 4154 bytes
--- PASS: TestConcurrent (1.05s)
PASS
ok      gostudy/09、基于共享变量的并发/files/memo4      1.084s
*/
