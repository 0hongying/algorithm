package main

// 给你一个字符串 s ，找出其中最长的回文子序列，并返回该序列的长度。
// 子序列定义为：不改变剩余字符顺序的情况下，删除某些字符或者不删除任何字符形成的一个序列

func longestPalindromeSubseq(s string) int {
	// dp[i][j]表示从第i个字符到第j个字符中最大的字符串
	n := len(s)
	dp := make([][]int, n)
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}

	// 斜着遍历，保证 dp[i][j]的左下右方向都被计算出来了
	// https://pic.leetcode-cn.com/3e044efa067077b64cf08c393e29d4025f26aa46eba1727c8948f3c0035a03da.jpg
	for i := n - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i+1][j])
			}
		}
	}

	return dp[0][n-1]
}
