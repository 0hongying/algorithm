package main

// 给定一个整数数组 prices，其中 prices[i]表示第 i 天的股票价格 ；整数 fee 代表了交易股票的手续费用。
// 你可以无限次地完成交易，但是你每笔交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。
// 返回获得利润的最大值。
// 注意：这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费

func maxProfit(prices []int, fee int) int {
	// dp[i][0]表示在第i天不持有
	// dp[i][1]表示第i天持有
	n := len(prices)
	dp := make([][]int, n)
	for i, _ := range prices {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < n; i++ {
		// 昨天不持有 和 昨天持有，今天不持有
		dp[i][0] = max(prices[i]+dp[i-1][1]-fee, dp[i-1][0])
		// 今天持有 和 昨天持有
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return dp[n-1][0]

}

// 给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
// 设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

func maxProfit1(prices []int) int {
	n := len(prices)
	dp := make([][]int, n)
	for i, _ := range prices {
		dp[i] = make([]int, 4)
	}

	// 第i天进行一次买操作
	dp[0][0] = -prices[0]
	// 第i天完成一笔交易
	dp[0][1] = 0
	// 第i天完成一笔交易且进行第二次买操作
	dp[0][2] = -prices[0]
	// 第i天完成第二笔交易
	dp[0][3] = 0

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], prices[i]+dp[i][0])
		dp[i][2] = max(dp[i-1][2], dp[i][1]-prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i][2]+prices[i])
	}

	return dp[n-1][3]
}
