package basename2

import (
	"strings"
)

func basename(s string) string {
	slash := strings.LastIndex(s, "/") // 如果没有找到/, 则slash取值-1
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
