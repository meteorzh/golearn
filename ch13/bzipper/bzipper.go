// Bzipper 读取输入，使用 bzip2 压缩然后输出数据
package main

import (
	"golearn/ch13/bzip"
	"io"
	"log"
	"os"
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
