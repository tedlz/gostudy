// Package methods 提供了一个函数，用于打印任意值的方法
package methods

import (
	"fmt"
	"reflect"
	"strings"
)

// Print 打印值 x 的方法集
func Print(x interface{}) {
	v := reflect.ValueOf(x)
	t := v.Type()
	fmt.Printf("type %s\n", t)

	for i := 0; i < v.NumMethod(); i++ {
		methType := v.Method(i).Type()
		fmt.Printf("func(%s) %s%s\n", t, t.Method(i).Name,
			strings.TrimPrefix(methType.String(), "func"))
	}
}
