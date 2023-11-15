package main

import "math"

// 给你一个整数数组 nums，请你找出并返回能被三整除的元素最大和
// dp[i][j]表示前i个余数为j的最大和
// dp[i][j] = max(dp[i-1][j], dp[i-1][(j-x%3+3)%3]+x)

func maxSumDivThree(nums []int) int {
	dp := make([][3]int, len(nums)+1)
	dp[0] = [3]int{0, math.MinInt64, math.MinInt64}

	for i, num := range nums {
		i++
		for j := 0; j < 3; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i-1][(j-num%3+3)%3]+num)
		}
	}
	return dp[len(nums)][0]
}

// func main() {
// 	nums := []int{3, 6, 7, 1, 10, 13}
// 	maxSumDivThree(nums)
// }
