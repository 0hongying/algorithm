package main

import (
	"sort"
)

// 给你一个整数数组 nums ，找出 nums 的下一个排列。
// 整数数组的 下一个排列 是指其整数的下一个字典序更大的排列
// 123 => 132
// 321 => 123
// 115 => 151
// 123465 => 123546

// func nextPermutation(nums []int) {
// 	if len(nums) <= 1 {
// 		return
// 	}

// 	i, j, k := len(nums)-2, len(nums)-1, len(nums)-1
// 	// 从后往前找，找到升序部分（A[i] < A[j]）
// 	for i >= 0 && nums[i] >= nums[j] {
// 		i--
// 		j--
// 	}

// 	// 不是最后一个排列
// 	if i >= 0 {
// 		// 从后往前找，找到升序部分（A[i] < A[k]）
// 		for nums[i] >= nums[k] {
// 			k--
// 		}
// 		// swap A[i], A[k]
// 		nums[i], nums[k] = nums[k], nums[i]
// 	}

// 	// reverse A[j:end]
// 	for i, j := j, len(nums)-1; i < j; i, j = i+1, j-1 {
// 		nums[i], nums[j] = nums[j], nums[i]
// 	}
// }

func nextPermutation(nums []int) {
	i := len(nums) - 1
	for ; i > 0 && nums[i] <= nums[i-1]; i-- {
	}

	if i > 0 {
		k := 0
		for j := i; j < len(nums); j++ {
			if nums[j] > nums[i-1] {
				k = j
			}
		}
		nums[k], nums[i-1] = nums[i-1], nums[k]
	}

	sort.Ints(nums[i:])
}

// func main() {
// 	nums := []int{
// 		5, 1, 1,
// 	}
// 	nextPermutation(nums)
// 	fmt.Printf("%v", nums)
// }

// 1 2 3 6 1 4 5
// 1 2 3 4 5 1 6
