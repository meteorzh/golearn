package main

import "fmt"

// 无限打印自然数的平方
func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// counter
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	// squarer
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	// printer (在主 goroutine 中)
	for {
		fmt.Println(<-squares)
	}
}
