package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// 024、示例：并发的目录遍历
// go run 024_du3.go -v /bin /usr /data $HOME
func main() {
	// 下面这个第三个版本的 du，会对每一个 walkDir 的调用创建一个新的 goroutine
	// 它使用 sync.WaitGroup 来对仍旧活跃的 walkDir 调用进行计数，
	// 另一个 goroutine 会在计数器减为零的时候将 fileSizes 这个 channel 关闭

	// 确定初始目录
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// 并行遍历文件树的每个根
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir3(root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	// 定期打印结果
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64

loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage3(nfiles, nbytes)
		}
	}
	printDiskUsage3(nfiles, nbytes)
}

var verbose = flag.Bool("v", false, "显示详细进度消息")

func printDiskUsage3(nfiles, nbytes int64) {
	fmt.Printf("%d files	%.1f GB\n", nfiles, float64(nbytes)/1e9)
}

// walkDir 递归遍历以 dir 为根的文件树，并在 fileSizes 上发送每个找到的文件大小
func walkDir3(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents3(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir3(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// sema 用来限制 dirents3 中的并发数量
var sema = make(chan struct{}, 20)

// dirents 返回目录 dir 的条目
func dirents3(dir string) []os.FileInfo {
	sema <- struct{}{}        // 获取 token
	defer func() { <-sema }() // 释放 token

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du3: %v\n", err)
		return nil
	}
	return entries
}
