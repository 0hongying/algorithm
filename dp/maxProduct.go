package main

// 给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
// 子数组 是数组的连续子序列

func maxProduct(nums []int) int {
	ans := nums[0]
	minN, maxN := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		temp1 := minN * nums[i]
		temp2 := maxN * nums[i]
		// 维护最小值和最大值，有可能出现交换的情况
		maxN = max(max(temp1, temp2), nums[i])
		minN = min(min(temp1, temp2), nums[i])
		ans = max(ans, maxN)
	}
	return ans
}
