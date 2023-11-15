package main

import (
	"sort"
)

// 给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
// 字母异位词 是由重新排列源单词的所有字母得到的一个新单词。

func groupAnagrams(strs []string) [][]string {
	// 将字符串排序后当作 key
	m := map[string][]string{}
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] > s[j]
		})

		sortedStr := string(s)
		m[sortedStr] = append(m[sortedStr], str)
	}
	ans := make([][]string, 0, len(m))
	for _, v := range m {
		ans = append(ans, v)
	}
	return ans
}

// func main() {
// 	strs := []string{
// 		"eat", "tea", "tan", "ate", "nat", "bat",
// 	}
// 	ret := groupAnagrams(strs)
// 	fmt.Printf("%v", ret)
// }
