package main

import "math"

// 给你一个整数数组 nums ，判断这个数组中是否存在长度为 3 的递增子序列
// 如果存在这样的三元组下标 (i, j, k) 且满足 i < j < k ，使得 nums[i] < nums[j] < nums[k] ，返回 true ；否则，返回 false

// func increasingTriplet(nums []int) bool {
// 	n := len(nums)
// 	// f[1]=x 代表长度为 1 的上升子序列最小结尾元素为 x
// 	// f[2]=y 代表长度为 2 的上升子序列的最小结尾元素为 y
// 	f := []int{math.MaxInt32, math.MaxInt32, math.MaxInt32}
// 	for i := 0; i < n; i++ {
// 		t := nums[i]
// 		if f[2] < t {
// 			// 最大数小于当前值
// 			return true
// 		} else if f[1] < t && t < f[2] {
// 			f[2] = t
// 		} else if f[1] > t {
// 			f[1] = t
// 		}
// 	}
// 	return false
// }

func increasingTriplet(nums []int) bool {
	n := len(nums)
	// f[1]=x 代表长度为 1 的上升子序列最小结尾元素为 x
	// f[2]=y 代表长度为 2 的上升子序列的最小结尾元素为 y
	f := []int{math.MaxInt32, math.MaxInt32, math.MaxInt32}
	for i := 0; i < n; i++ {
		t := nums[i]
		if f[2] < t {
			// 最大数小于当前值
			return true
		} else if f[1] < t && t < f[2] {
			f[2] = t
		} else if f[1] > t {
			f[1] = t
		}
	}
	return false
}

// func main() {
// 	nums := []int{2, 6, 3, 4, 5}
// 	increasingTriplet(nums)
// }
