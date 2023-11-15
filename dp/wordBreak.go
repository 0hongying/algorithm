package main

// 给你一个字符串 s 和一个字符串列表 wordDict 作为字典。请你判断是否可以利用字典中出现的单词拼接出 s 。
// 注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用
func wordBreak(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)
	for _, w := range wordDict {
		wordDictSet[w] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	// dp[i]表示以i结尾的字符串是否符合条件
	// dp[i]= dp[j] && check(s[j..i−1])
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

// func main() {
// 	s := "leetcode"
// 	wordDict := []string{"leet", "code"}
// 	ret := wordBreak(s, wordDict)
// 	fmt.Printf("%v", ret)
// }
