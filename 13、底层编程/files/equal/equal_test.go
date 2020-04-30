package equal

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEqual(t *testing.T) {
	one, oneAgain, two := 1, 1, 2

	type CyclePtr *CyclePtr
	var cyclePtr1, cyclePtr2 CyclePtr
	cyclePtr1 = &cyclePtr1
	cyclePtr2 = &cyclePtr2

	type CycleSlice []CycleSlice
	var cycleSlice = make(CycleSlice, 1)
	cycleSlice[0] = cycleSlice

	ch1, ch2 := make(chan int), make(chan int)
	var ch1ro <-chan int = ch1

	type mystring string

	var iface1, iface1Again, iface2 interface{} = &one, &oneAgain, &two

	for _, test := range []struct {
		x, y interface{}
		want bool
	}{
		// 基本类型
		{1, 1, true},
		{1, 2, false},   // 不同值
		{1, 1.0, false}, // 不同类型
		{"foo", "foo", true},
		{"foo", "bar", false},
		{mystring("foo"), "foo", false}, // 不同类型
		// slice
		{[]string{"foo"}, []string{"foo"}, true},
		{[]string{"foo"}, []string{"bar"}, false},
		{[]string{}, []string(nil), true},
		// slice 循环
		{cycleSlice, cycleSlice, true},
		// map
		{
			map[string][]int{"foo": {1, 2, 3}},
			map[string][]int{"foo": {1, 2, 3}},
			true,
		},
		{
			map[string][]int{"foo": {1, 2, 3}},
			map[string][]int{"foo": {1, 2, 3, 4}},
			false,
		},
		{
			map[string][]int{},
			map[string][]int(nil),
			true,
		},
		// pointer
		{&one, &one, true},
		{&one, &two, false},
		{&one, &oneAgain, true},
		{new(bytes.Buffer), new(bytes.Buffer), true},
		// pointer 循环
		{cyclePtr1, cyclePtr1, true},
		{cyclePtr2, cyclePtr2, true},
		{cyclePtr1, cyclePtr2, true}, // 它们深度相等
		// function
		{(func())(nil), (func())(nil), true},
		{(func())(nil), func() {}, false},
		{func() {}, func() {}, false},
		// array
		{[...]int{1, 2, 3}, [...]int{1, 2, 3}, true},
		{[...]int{1, 2, 3}, [...]int{1, 2, 4}, false},
		// channel
		{ch1, ch1, true},
		{ch1, ch2, false},
		{ch1ro, ch1, false}, // 注意：不相等
		// interface
		{&iface1, &iface1, true},
		{&iface1, &iface2, false},
		{&iface1Again, &iface1, true},
	} {
		if Equal(test.x, test.y) != test.want {
			t.Errorf("Equal(%v, %v) = %t", test.x, test.y, !test.want)
		}
	}
}

func Example_equal() {
	fmt.Println(Equal([]int{1, 2, 3}, []int{1, 2, 3}))        // true
	fmt.Println(Equal([]string{"foo"}, []string{"bar"}))      // false
	fmt.Println(Equal([]string(nil), []string{}))             // true
	fmt.Println(Equal(map[string]int(nil), map[string]int{})) // true
}

func Example_equalCycle() {
	// 循环链接列表 a -> b -> a 和 c -> c
	type link struct {
		value string
		tail  *link
	}
	a, b, c := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}
	a.tail, b.tail, c.tail = b, a, c
	fmt.Println(Equal(a, a)) // true
	fmt.Println(Equal(b, b)) // true
	fmt.Println(Equal(c, c)) // true
	fmt.Println(Equal(a, b)) // false
	fmt.Println(Equal(a, c)) // false
}
