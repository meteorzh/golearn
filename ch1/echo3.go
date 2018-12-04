package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func mainecho3() {
	// 1.使用strings.Join方法
	// fmt.Println(strings.Join(os.Args[1:], " "))
	// 2.直接打印slice
	// fmt.Println(os.Args[1:])

	// 3.比较循环拼接字符串和使用strings.Join两种方式的耗时
	s, sep := "", ""
	start := time.Now()
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(time.Since(start).Nanoseconds())
	fmt.Println(s)

	start = time.Now()
	s = strings.Join(os.Args[1:], " ")
	fmt.Println(time.Since(start).Nanoseconds())
	fmt.Println(s)

}
