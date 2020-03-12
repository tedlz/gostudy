// Package thumbnail 可以从较大尺寸的图像生成缩略图像
// 当前仅支持 JPEG
package thumbnail

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Image 返回 src 的缩略图
func Image(src image.Image) image.Image {
	// 计算缩略图大小，保留长宽比
	xs := src.Bounds().Size().X
	ys := src.Bounds().Size().Y
	width, height := 128, 128
	if aspect := float64(xs) / float64(ys); aspect < 1.0 {
		width = int(128 * aspect) // 竖向
	} else {
		height = int(128 / aspect) // 横向
	}
	xscale := float64(xs) / float64(width)
	yscale := float64(ys) / float64(height)

	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	// 一个非常粗糙的缩放算法
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			srcx := int(float64(x) * xscale)
			srcy := int(float64(y) * yscale)
			dst.Set(x, y, src.At(srcx, srcy))
		}
	}
	return dst
}

// ImageStream 从 r 读取图像，并将缩略图写入 w
func ImageStream(w io.Writer, r io.Reader) error {
	src, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	dst := Image(src)
	return jpeg.Encode(w, dst, nil)
}

// ImageFile2 从 infile 读取图像，并将缩略图写入 outfile
func ImageFile2(outfile, infile string) (err error) {
	in, err := os.Open(infile)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(outfile)
	if err != nil {
		return err
	}

	if err := ImageStream(out, in); err != nil {
		out.Close()
		return fmt.Errorf("缩放 %s 到 %s: %s", infile, outfile, err)
	}
	return out.Close()
}

// ImageFile 从 infile 读取图像，并将缩略图写入同一目录
// 它返回生成的文件名，如 “foo.thumb.jpg”
func ImageFile(infile string) (string, error) {
	ext := filepath.Ext(infile) // 例如 “.jpg”、“.JPEG”
	outfile := strings.TrimSuffix(infile, ext) + ".thumb" + ext
	return outfile, ImageFile2(outfile, infile)
}
