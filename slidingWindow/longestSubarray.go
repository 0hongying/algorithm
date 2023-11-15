package main

// 给你一个二进制数组 nums ，你需要从中删掉一个元素。
// 请你在删掉元素的结果数组中，返回最长的且只包含 1 的非空子数组的长度

func longestSubarray(nums []int) int {
	left, right := 0, 0
	ret := 0
	zero := 0
	for right < len(nums) {
		// 维护窗口内只有一个零
		if nums[right] == 0 {
			zero++
		}

		for zero > 1 {
			if nums[left] == 0 {
				zero--
			}
			left++
		}

		if ret < right-left {
			ret = right - left
		}
		right++
	}

	return ret
}

// func main() {
// 	nums := []int{0, 1, 1, 1, 0, 1, 1, 0, 1}
// 	longestSubarray(nums)
// }
