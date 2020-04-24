package format_test

import (
	"fmt"
	"gostudy/12、反射/files/format"
	"testing"
	"time"
)

func Test(t *testing.T) {
	// 指针值只是示例，可能会因运行情况不同而有所差异
	var x int64 = 1
	var d time.Duration = 1 * time.Nanosecond
	fmt.Println(format.Any(x))                  // 1
	fmt.Println(format.Any(d))                  // 1
	fmt.Println(format.Any([]int64{x}))         // []int64 0xc000016210
	fmt.Println(format.Any([]time.Duration{d})) // []time.Duration 0xc000016218
}
