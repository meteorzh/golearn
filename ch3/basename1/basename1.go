package basename1

// basename 移除路径部分和.后缀
// 例如：a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// 保留最后一个.之前的所有内容
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
