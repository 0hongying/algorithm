package main

// 求一个字符串s中不含有重复字符的最长子串的长度
func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	n := len(s)
	if n == 0 {
		return 0
	}
	var max = 0
	var left = 0
	for i := 0; i < len(s); i++ {
		if v, ok := m[s[i]]; ok {
			left = Max(left, v+1)
		}
		m[s[i]] = i
		max = Max(max, i-left+1)
	}
	return max
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// func main() {
// 	s := "abba"
// 	ret := lengthOfLongestSubstring(s)
// 	fmt.Printf("%v", ret)
// }
