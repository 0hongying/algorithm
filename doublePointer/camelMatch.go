package main

// https://leetcode.cn/problems/camelcase-matching/description/

func camelMatch(queries []string, pattern string) (ret []bool) {
	for i := 0; i < len(queries); i++ {
		ret = append(ret, check(queries[i], pattern))
	}
	return
}

func check(s, t string) bool {
	m, n := len(s), len(t)
	i, j := 0, 0
	for ; j < n; i, j = i+1, j+1 {
		// 不匹配小写字母
		for i < m && s[i] != t[j] && (s[i] >= 'a' && s[i] <= 'z') {
			i++
		}
		// 不匹配大写字母
		if i == m || s[i] != t[j] {
			return false
		}
	}
	for i < m && s[i] >= 'a' && s[i] <= 'z' {
		i++
	}
	return i == m
}

// func main() {
// 	queries := []string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}
// 	pattern := "FoBaT"

// 	ret := camelMatch(queries, pattern)
// 	fmt.Printf("%v", ret)
// }
