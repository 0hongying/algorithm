package main

import "strings"

// 最后一个单词的长度
func lengthOfLastWord(s string) int {
	l := len(s)
	s = strings.Trim(s, " ")
	ret := 0
	for i := l - 1; i >= 0; i-- {
		if string(s[i]) != " " {
			ret++
		} else if string(s[i]) == " " && ret != 0 {
			break
		}
	}
	return ret
}

// func main() {
// 	s := "   fly me   to   the moon  "
// 	ret := lengthOfLastWord(s)
// 	fmt.Printf("%v", ret)
// }
