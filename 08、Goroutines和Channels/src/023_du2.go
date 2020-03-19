package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

// 023、示例：并发的目录遍历
// go run 023_du2.go -v /bin /usr /data $HOME
func main() {
	// 下面这个 du 的变种会间歇打印内容，不过只有在调用时提供了 -v 的 flag 才会显示程序进度信息
	// 在 roots 目录上循环的后台 goroutine 在这里保持不变
	// 主 goroutine 现在使用了计时器来每 500ms 生成事件，然后用 select 语句来等待文件大小的消息来更新总大小数据
	// 如果 -v 的 flag 在运行时没有传入的话，tick 这个 channel 会保持为 nil，这样在 select 里的 case 也就相当于被禁用了

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
			walkDir2(root, fileSizes)
		}
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
				break loop // fileSizes 已关闭
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage2(nfiles, nbytes) // 循环打印结果
		}
	}
	printDiskUsage2(nfiles, nbytes) // 最终结果

	// 然而这个程序还是会花上很长时间才结束，无法对 walkDir 做并行化处理是因为磁盘系统并行限制
}

var verbose = flag.Bool("v", false, "显示详细进度消息")

func printDiskUsage2(nfiles, nbytes int64) {
	fmt.Printf("%d files	%.1f GB\n", nfiles, float64(nbytes)/1e9)
}

// walkDir 递归遍历以 dir 为根的文件树，并在 fileSizes 上发送每个找到的文件大小
func walkDir2(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents2(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir2(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// dirents 返回目录 dir 的条目
func dirents2(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du2: %v\n", err)
		return nil
	}
	return entries
}
