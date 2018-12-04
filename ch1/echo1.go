// echo1 输出命令行参数
package main

import (
	"fmt"
	"os"
)

func mainecho1() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
