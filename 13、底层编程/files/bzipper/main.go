// bzipper 读取输入，bzip2 对其进行压缩，然后将其写入
package main

import (
	"io"
	"log"
	"os"

	"gostudy/13、底层编程/files/bzip"
)

func main() {
	w := bzip.NewWriter(os.Stdout)
	if _, err := io.Copy(w, os.Stdin); err != nil {
		log.Fatalf("bzipper: %v\n", err)
	}
	if err := w.Close(); err != nil {
		log.Fatalf("bzipper: close: %v\n", err)
	}
}
