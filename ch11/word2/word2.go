// 包 word 提供了文字游戏相关的函数
package word2

import "unicode"

// IsPalindrome 判断一个字符串是否是回文字符串
// 忽略字母大小写，以及非字母字符
func IsPalindrome(s string) bool {
	letters := make([]rune, 0, len(s))
	// var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
