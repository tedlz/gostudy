package main

import (
	"fmt"
	"sort"
)

// 011、匿名函数
// go run 011_toposort.go
// 输出：
// 1:      intro to programming
// 2:      discrete math
// 3:      data structures
// 4:      algorithms
// 5:      linear algebra
// 6:      calculus
// 7:      formal languages
// 8:      computer organization
// 9:      compilers
// 10:     databases
// 11:     operating systems
// 12:     networks
// 13:     programming languages
func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

// prereqs 记录了每个课程的前置课程
// 这类问题被称作拓补排序
var prereqs = map[string][]string{
	"algorithms": {"data structures"}, // 算法: {数据结构}
	"calculus":   {"linear algebra"},  // 微积分: {线性代数}
	"compilers": { // 编译器
		"data structures",       // 数据结构
		"formal languages",      // 形式语言
		"computer organization", // 计算机组成
	},
	"data structures":       {"discrete math"},                            // 数据结构: {离散数学}
	"databases":             {"data structures"},                          // 数据库: {数据结构}
	"discrete math":         {"intro to programming"},                     // 离散数学: {编程入门}
	"formal languages":      {"discrete math"},                            // 形式语言: {离散数学}
	"networks":              {"operating systems"},                        // 网络: {操作系统}
	"operating systems":     {"data structures", "computer organization"}, // 操作系统: {数据结构, 计算机组成}
	"programming languages": {"data structures", "computer organization"}, // 编程语言: {数据结构, 计算机组成}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}
