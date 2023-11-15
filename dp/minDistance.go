package main

// 给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数
// 你可以对一个单词进行如下三种操作：
// 插入一个字符
// 删除一个字符
// 替换一个字符

// dp[i][j]表示 work1 从 i 位置转变成 word2 到 j 位置需要最少步数
// dp[i-1][j-1] 是替换操作
// dp[i-1][j] 是删除操作
// dp[i][j-1] 是插入操作

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i, _ := range dp {
		dp[i] = make([]int, len(word2)+1)
	}

	for i := 1; i <= len(word1); i++ {
		dp[i][0] = i
	}

	for i := 1; i <= len(word2); i++ {
		dp[0][i] = i
	}

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(min(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1]) + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

// func main() {
// 	word1 := "horse"
// 	word2 := "ros"
// 	ret := minDistance(word1, word2)
// 	fmt.Printf("%v", ret)
// }
