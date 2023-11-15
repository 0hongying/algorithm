package main

import (
	"math"
)

// 给定一个三角形 triangle ，找出自顶向下的最小路径和
// 每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([][]int, len(triangle))
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	// dp[i][j]表示移动到i,j的最小路径和
	dp[0][0] = triangle[0][0]

	for i := 1; i < n; i++ {
		// 第一列相邻结点只有上结点
		dp[i][0] = triangle[i][0] + dp[i-1][0]
		for j := 1; j < i; j++ {
			// 选上一步相邻结点中最小路径的一个 + 本结点
			dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
		}
		// 第n行第n列相邻结点只有第n-1行第n-1列
		dp[i][i] = triangle[i][i] + dp[i-1][i-1]
	}

	ret := math.MaxInt32
	for i := 0; i < n; i++ {
		// 从最后一行找最小
		ret = min(ret, dp[n-1][i])
	}
	return ret
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// func main() {
// 	triangle := [][]int{
// 		{2},
// 		{3, 4},
// 		{6, 5, 7},
// 		{4, 1, 8, 3},
// 	}
// 	ret := minimumTotal(triangle)
// 	fmt.Printf("%v", ret)
// }
