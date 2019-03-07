package equal

import (
	"fmt"
	"testing"
)

func TestEqual(t *testing.T) {
	fmt.Println(Equal(true, true))
	//fmt.Println(Equal([]int{1, 2, 3}, []int{1, 2, 3}))
	fmt.Println(Equal([]string{"foo"}, []string{"bar"}))
	fmt.Println(Equal([]string(nil), []string{}))
	// fmt.Println(Equal(map[string]int(nil), map[string]int{}))
}
