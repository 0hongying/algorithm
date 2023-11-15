package main

// 峰值元素是指其值严格大于左右相邻值的元素。
// 给你一个整数数组 nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可
func findPeakElement(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[mid+1] {
			// 峰值在左边
			right = mid
		} else {
			// 峰值在右边
			left = mid + 1
		}
	}
	return right
}
