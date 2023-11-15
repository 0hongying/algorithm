package main

import "strings"

func maxVowels(s string, k int) int {
	left, right := 0, 0
	ret := 0
	mete := 0
	for right < len(s) {

		if strings.Contains("aeiou", string(s[right])) {
			mete++
		}

		for right-left >= k {
			if strings.Contains("aeiou", string(s[left])) {
				mete--
			}
			left++
		}

		if ret < mete {
			ret = mete
		}
		right++
	}
	return ret
}

// func main() {
// 	s := "aeiou"
// 	k := 3
// 	maxVowels(s, k)
// }
