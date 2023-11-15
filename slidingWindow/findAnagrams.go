package main

// 给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
// 异位词 指由相同字母重排列形成的字符串（包括相同的字符串)

func findAnagrams(s string, p string) []int {
	ans := []int{}
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return ans
	}

	left := 0
	right := sLen - pLen
	pArr, sArr := [26]int{}, [26]int{}
	// 初始化窗口
	for i, v := range p {
		pArr[v-'a']++
		sArr[s[i]-'a']++
	}

	for left <= right {
		// 数组可以直接比较
		if pArr == sArr {
			ans = append(ans, left)
		}
		sArr[s[left]-'a']--
		if left+pLen < sLen {
			sArr[s[left+pLen]-'a']++
		}
		left++
	}

	return ans
}

// func main() {
// 	s := "abab"
// 	p := "ab"
// 	findAnagrams(s, p)
// }
