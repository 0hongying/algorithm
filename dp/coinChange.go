package main

import "math"

// 给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
// 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i, _ := range dp {
		dp[i] = math.MaxInt32
	}

	dp[0] = 0
	for i := 1; i < amount; i++ {
		for _, v := range coins {
			if i >= v {
				dp[i] = min(dp[i], dp[i-v]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
