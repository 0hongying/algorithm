package main

// 给你一个整数数组 nums，返回 nums 中最长等差子序列的长度。

func longestArithSeqLength(nums []int) int {
	n := len(nums)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, 1001)
	}
	ans := 0
	for i := 1; i < n; i++ {
		for k := 0; k < i; k++ {
			j := nums[i] - nums[k] + 500
			f[i][j] = max(f[i][j], f[k][j]+1)
			ans = max(ans, f[i][j])
		}
	}
	return ans + 1
}
