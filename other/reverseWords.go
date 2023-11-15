package main

import "strings"

func reverseWords(s string) string {
	s = strings.Trim(s, " ")
	left, right := len(s)-1, len(s)
	ret := ""
	for left >= 0 {
		for left >= 0 && s[left] != ' ' {
			left--
		}
		// 只保留单词，空格自己补充
		ret += string(s[left+1:right]) + " "
		for left >= 0 && s[left] == ' ' {
			left--
		}
		right = left + 1
	}
	ret = strings.Trim(ret, " ")
	return ret
}

// func main() {
// 	s := "a good   example"
// 	reverseWords(s)
// }
