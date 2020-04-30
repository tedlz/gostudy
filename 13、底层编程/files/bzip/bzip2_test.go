package bzip_test

import (
	"bytes"
	"compress/bzip2"
	"io"
	"testing"

	"gostudy/13、底层编程/files/bzip"
)

// !
// ! 注意：本人在 debian 10 下运行 bzip2_test.go 时遇到了以下错误：
// ! fatal error: bzlib.h: No such file or directory
// ! 解决方法：
// ! sudo apt install libbz2-dev
// !

func TestBzip2(t *testing.T) {
	var compressed, uncompressed bytes.Buffer
	w := bzip.NewWriter(&compressed)

	// 写一百万条重复的信息，压缩一个副本但不压缩另一个副本
	tee := io.MultiWriter(w, &uncompressed)
	for i := 0; i < 1000000; i++ {
		io.WriteString(tee, "hello")
	}
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}

	// 检查压缩流的大小
	if got, want := compressed.Len(), 255; got != want {
		t.Errorf("1 million hellos compressed to %d bytes, want %d", got, want)
	}

	// 解压缩并与原件比较
	var decompressed bytes.Buffer
	io.Copy(&decompressed, bzip2.NewReader(&compressed))
	if !bytes.Equal(uncompressed.Bytes(), decompressed.Bytes()) {
		t.Error("decompression yielded a different message")
	}
}
