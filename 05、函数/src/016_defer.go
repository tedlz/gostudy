package main

import (
	"fmt"
	"log"
	"time"
)

// 016、延迟函数调用
// go run 016_defer.go
// 输出：
// 2020/01/09 10:56:58 enter bigSlowOperation
// 2020/01/09 10:57:01 exit bigSlowOperation (3.000667366s)
// double(4) = 8
// double(4) = 8
// 12
func main() {
	bigSlowOperation()
	_ = double(4)
	fmt.Println(triple(4))
}

// 在处理其它资源时，也可以采用 defer 机制，比如对文件的操作
// func readFile(filename string) ([]byte, error) {
// 	   f, err := os.Open(filename)
// 	   if err != nil {
// 		   return nil, err
// 	   }
// 	   defer f.Close()
// 	   return readAll(f)
// }
// 或是处理互斥锁
// var mu sync.Mutex
// var m = make(map[string]int)
// func lookup(key int) string {
// 	   mu.Lock()
// 	   defer mu.Unlock()
// 	   return m[key]
// }

// 调试复杂程序时，defer 机制也常被用于记录何时进入和退出函数
// 不要忘记 defer 后面的圆括号 ()，否则本该在进入时执行的操作会在退出时执行；本该在退出时执行的，永远不会被执行
func bigSlowOperation() {
	defer trace("bigSlowOperation")()
	time.Sleep(3 * time.Second)
}
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

// defer 语句中的函数会在 return 语句更新返回值变量后再执行
// 又因为在函数中定义的匿名函数可以访问该函数包括返回值变量在内的所有变量
// 所以，对匿名函数采用 defer 机制，可以使其观察函数的返回值
func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	return x + x
}

// 被延迟执行的匿名函数甚至可以修改函数返回给调用者的返回值
func triple(x int) (result int) {
	defer func() { result += x }()
	return double(x)
}

// 在循环体中的 defer 语句需要特别注意，因为只有在函数执行完毕后，这些被延迟的函数才会执行
// 下面的代码会导致系统的文件描述符耗尽，因为在所有文件处理完之前，没有文件会被关闭
// func readFile(filenames []string) {
// 	   for _, filename := range filenames {
// 		   f, err := os.Open(filename)
// 		   if err != nil {
// 			   return err
// 		   }
// 		   defer f.Close() // NOTE: risky; could run out of file descriptors
// 	   }
// }

// 一种解决方法是将循环体中的 defer 语句移至另外一个函数，每次循环时调用这个函数
// func readFile2(filenames []string) {
// 	   for _, filename := range filenames {
// 		   err := doFile(filename)
// 	   }
// }
// func doFile(filename string) error {
// 	   f, err := os.Open(filename)
// 	   if err != nil {
// 		   return err
// 	   }
// 	   defer f.Close()
// }
