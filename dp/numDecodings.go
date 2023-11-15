package main

// https://leetcode.cn/problems/decode-ways/description/
func numDecodings(s string) int {
	n := len(s)
	// 跳楼梯
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}

		if i > 1 && s[i-2] != '0' && int(s[i-2]-'0')*10+int(s[i-1]-'0') <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}

// func main() {
// 	s := "2213"
// 	ret := numDecodings(s)
// 	fmt.Printf("%v", ret)
// }
