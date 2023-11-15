package main

// 给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积
func productExceptSelf(nums []int) []int {
	ret := make([]int, len(nums))

	k := 1
	// 除该元素左边的乘积
	for i := 0; i < len(nums); i++ {
		ret[i] = k
		k *= nums[i]
	}

	k = 1
	// 除该元素右边的乘积
	for i := len(nums) - 1; i >= 0; i-- {
		ret[i] *= k
		k *= nums[i]
	}

	return ret
}
