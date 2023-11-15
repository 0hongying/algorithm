package main

// 给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)
	first, last := -1, -1

	// 找左边界[left, right)
	for left < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			first = mid
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	// 找右边界[left, right)
	left = 0
	right = len(nums)
	for left < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			last = mid
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return []int{first, last}
}

// func main() {
// 	nums := []int{
// 		5, 7, 7, 7, 7, 10,
// 	}
// 	target := 10
// 	ret := searchRange(nums, target)
// 	fmt.Printf("%v", ret)
// }
