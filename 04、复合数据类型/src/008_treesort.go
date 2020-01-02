package main

import "fmt"

// 008、二叉树插入排序
// go run 008_treesort.go
// 输出：
// [1 3 5 7 9 2 4 6 8 0]
// [0 1 2 3 4 5 6 7 8 9]
func main() {
	arr := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 0}
	fmt.Println(arr) // [1 3 5 7 9 2 4 6 8 0]
	Sort(arr)
	fmt.Println(arr) // [0 1 2 3 4 5 6 7 8 9]
}

// 一个命名为 S 的结构体类型，将不能再包含 S 类型的成员（因为一个聚合的值不能包含它自身，数组同理）
// 但 S 类型的结构体允许包含 *s 指针类型的成员，这可以让我们创建递归的数据结构，比如链表和树结构等
type tree struct {
	value       int
	left, right *tree
}

// Sort 对值进行排序
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add2(root, v)
	}
	appendValues(values[:0], root)
}

func add2(t *tree, value int) *tree {
	if t == nil {
		// 相当于 return &tree{value: value}
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add2(t.left, value)
	} else {
		t.right = add2(t.right, value)
	}
	return t
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}
