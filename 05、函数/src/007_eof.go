package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 007、文件结尾错误 EOF
func main() {
	// 从文件中读取 n 个字节
	// 如果 n 等于文件长度，读取过程的任何错误都表示失败
	// 如果 n 小于文件长度，调用者会重复读取固定大小的数据直到文件结束
	// 这会导致调用者必须分别处理由文件结束引起的各种错误
	// 基于这样的原因，io 包保证任何由文件结束引起的读取失败都返回同一个错误，io.EOF

	// 下面的例子展示了如何从标准输入中读取字符，以及判断文件结束
	read()
	// （04、复合数据类型的 005_charcount.go 展示了更加复杂的代码）
	// 因为文件结束这种错误不需要更多描述，所以 io.EOF 有固定的错误信息 EOF
	// 对于其它错误，我们可能需要在错误信息中描述错误的类型和数量，这使得我们不能像 io.EOF 采用固定的错误信息
	// 在 7.11 节中，我们会提出更系统的方法区分某些固定的错误值
}

func read() error {
	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break // 结束读取
		}
		if err != nil {
			return fmt.Errorf("read failed: %v", err)
		}
		// ...
	}
	return nil
}
