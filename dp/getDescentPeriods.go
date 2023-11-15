package main

// 给你一个整数数组 prices ，表示一支股票的历史每日股价，其中 prices[i] 是这支股票第 i 天的价格。
// 一个 平滑下降的阶段 定义为：对于 连续一天或者多天 ，每日股价都比 前一日股价恰好少 1 ，这个阶段第一天的股价没有限制。
// 请你返回 平滑下降阶段 的数

func getDescentPeriods(prices []int) int64 {
	l := len(prices)
	if l == 0 || l == 1 {
		return int64(l)
	}
	dp := make([]int64, len(prices))
	var ans int64 = 0
	for i := 0; i < l; i++ {
		if i > 0 && prices[i-1]-prices[i] == 1 {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = 1
		}
		ans += dp[i]
	}
	return ans
}
