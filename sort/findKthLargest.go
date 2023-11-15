package main

import "fmt"

// 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
func findKthLargest(nums []int, k int) int {
	TopKSplit(nums, len(nums)-k, 0, len(nums)-1)
	return nums[len(nums)-k]
}

func Parition(nums []int, start, stop int) int {
	if start >= stop {
		return -1
	}
	pivot := nums[start]
	left, right := start, stop
	for left < right {
		// 从右边找比基准小的
		for left < right && nums[right] >= pivot {
			right--
		}
		nums[left] = nums[right]
		// 从左边找比基准大的
		for left < right && nums[left] <= pivot {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = pivot
	return left
}

func TopKSplit(nums []int, k, start, stop int) {
	if start < stop {
		index := Parition(nums, start, stop)
		if index == k {
			return
		} else if index < k {
			// 继续在 index 右边寻找
			TopKSplit(nums, k, index+1, stop)
		} else {
			// 继续在 index 左边寻找
			TopKSplit(nums, k, start, index-1)
		}
	}
}

// func findKthLargest(nums []int, k int) int {
// 	for i := 0; i < k; i++ {
// 		bubbleSort(nums)
// 	}
// 	return nums[len(nums)-k]
// }

// func bubbleSort(nums []int) {
// 	for i, v := range nums {
// 		if i < len(nums)-1 && v > nums[i+1] {
// 			nums[i], nums[i+1] = nums[i+1], nums[i]
// 		}
// 	}
// }

func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	ret := findKthLargest(nums, 4)
	fmt.Printf("%v", ret)
}
