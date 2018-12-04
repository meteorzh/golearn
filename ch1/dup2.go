// dup2 打印输入中多次出现的行的个数和文本
// 它从stdin或指定的文件列表读取
package main

import (
	"bufio"
	"fmt"
	"os"
)

// 原版
// func main() {
// 	counts := make(map[string]int)
// 	files := os.Args[1:]
// 	if len(files) == 0 {
// 		countLines(os.Stdin, counts)
// 	} else {
// 		for _, arg := range files {
// 			f, err := os.Open(arg)
// 			if err != nil {
// 				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
// 				continue
// 			}
// 			countLines(f, counts)
// 			f.Close()
// 		}
// 	}

// 	for line, n := range counts {
// 		if n > 1 {
// 			fmt.Printf("%d\t%s\n", n, line)
// 		}
// 	}
// }

// func countLines(f *os.File, counts map[string]int) {
// 	input := bufio.NewScanner(f)
// 	for input.Scan() {
// 		counts[input.Text()]++
// 	}
// 	// 注意： 忽略input.Err()中可能的错误
// }

// 练习输出文件名改版
func maindup2() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for filename, fileCount := range counts {
		for line, n := range fileCount {
			if n > 1 {
				fmt.Printf("%d\t%s\t%s\n", n, line, filename)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	counts[f.Name()] = make(map[string]int)
	for input.Scan() {
		counts[f.Name()][input.Text()]++
	}
	// 注意： 忽略input.Err()中可能的错误
}
