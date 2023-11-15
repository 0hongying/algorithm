package main

import (
	"math"
)

// 找出该数组中满足其和≥target的长度最小的连续子数组，并返回其长度。如果不存在符合条件的子数组，返回0
func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 滑动窗口的左边界和右边界
	left, right := 0, 0
	sum := 0
	ret := math.MaxInt
	for right < len(nums) {
		sum += nums[right]
		for sum >= target {
			ret = min(ret, right-left+1)
			sum -= nums[left]
			left++
		}
		right++
	}
	if ret == math.MaxInt {
		ret = 0
	}
	return ret
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// func main() {
// 	target := 14
// 	nums := []int{1, 1, 1, 1, 1, 1, 1, 1}
// 	ret := minSubArrayLen(target, nums)
// 	fmt.Printf("%v", ret)
// }
