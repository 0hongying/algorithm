package main

// 给定一个包含 n + 1 个整数的数组 nums ，其数字都在 [1, n] 范围内（包括 1 和 n），可知至少存在一个重复的整数。
// 假设 nums 只有 一个重复的整数 ，返回 这个重复的数

func findDuplicate(nums []int) int {
	left, right := 0, len(nums)
	ans := 0
	for left < right {
		mid := (left + right) / 2

		count := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] <= mid {
				count++
			}
		}

		if count <= mid {
			left = mid + 1
		} else {
			right = mid
			ans = mid
		}

	}
	return ans
}

// func main() {
// 	nums := []int{3, 1, 3, 4, 2}
// 	findDuplicate(nums)
// }
