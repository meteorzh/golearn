// Nonempty 演示了slice的就地修改算法
package main

import (
	"fmt"
)

// nonempty返回一个新的slice，slice中的元素都是非空字符串
// 在函数的调用过程中，底层数组的元素发生了改变
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

// 通过append方式实现
func nonempty2(strings []string) []string {
	out := strings[:0] // 引用原始slice的新的零长度的slice
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

// 移除slice中的元素，并保留剩余元素的顺序
func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

// 移除slice中的元素，不保留顺序
func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data))
	fmt.Printf("%q\n", data)
}
