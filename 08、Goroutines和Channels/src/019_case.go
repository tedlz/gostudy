package main

import "fmt"

// 019、基于 select 的多路复用 - case
// go run 019_case.go
func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println("A", i, x)
		case ch <- i:
			fmt.Println("B", i, i)
		}
	}
}
