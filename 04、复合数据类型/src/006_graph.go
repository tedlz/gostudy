package main

import "fmt"

// 006、图表
// go run 006_graph.go
// 输出
// true
// false
func main() {
	addEdge("a", "b")
	fmt.Println(hasEdge("a", "b")) // true
	fmt.Println(hasEdge("x", "y")) // false
}

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}
