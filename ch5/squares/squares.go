package main

import (
	"fmt"
)

// squares 函数返回一个函数，后者包括下一次要用到的平方数
// the next square number each time it is called
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
