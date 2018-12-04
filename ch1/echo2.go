package main

import (
	"fmt"
	"os"
)

func mainecho2() {
	// s, sep := "", ""
	// for _, arg := range os.Args[1:] {
	// 	s += sep + arg
	// 	sep = " "
	// }
	// fmt.Println(s)
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
}
