// jpeg 命令从标准输入读取 png 图像
// 并将其作为 jpeg 图像写入标准输出

package main

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png" // 注册 png 解码器
	"io"
	"os"
)

// go run ../../../03、基础数据类型/src/005_mandelbrot.go | go run main.go >mandelbrot.jpg
func main() {
	if err := toJPEG(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}
