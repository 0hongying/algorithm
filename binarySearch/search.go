package main

// https://leetcode.cn/problems/search-in-rotated-sorted-array/?envType=study-plan-v2&envId=top-100-liked

func search(nums []int, target int) int {
	index := findMin(nums)
	if target > nums[len(nums)-1] {
		return binarySearch(nums, 0, index, target)
	}
	return binarySearch(nums, index, len(nums), target)
}

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func binarySearch(nums []int, left, right, target int) int {
	for left < right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			return mid
		} else {
			right = mid
		}
	}
	return -1
}
