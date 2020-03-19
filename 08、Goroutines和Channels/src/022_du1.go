package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// 022、示例：并发的目录遍历
// go run 022_du1.go $HOME
func main() {
	// 确定初始目录
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// 遍历文件树
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	// 打印结果
	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles++
		nbytes += size
	}
	printDiskUsage(nfiles, nbytes)

	// 这个程序会在打印其结果之前卡住很长时间，如果运行时能让我们知道处理进度的话想必更好
	// 但是，如果简单地把 printDiskUsage 函数调用移动到循环里会导致其打印大量的输出
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files	%.1f GB\n", nfiles, float64(nbytes)/1e9)
}

// walkDir 递归遍历以 dir 为根的文件树，并在 fileSizes 上发送每个找到的文件大小
func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// dirents 返回目录 dir 的条目
func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}
