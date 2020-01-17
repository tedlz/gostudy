package main

import "fmt"

// 001、接口是合约
// go run 001_bytecounter.go
// 输出：
// 5
// 12
func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // 5
	c = 0          // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // 12
}

// ByteCounter *
type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}
