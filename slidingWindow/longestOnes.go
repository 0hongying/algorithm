package main

// 给定一个二进制数组 nums 和一个整数 k，如果可以翻转最多 k 个 0 ，则返回 数组中连续 1 的最大个数
func longestOnes(nums []int, k int) int {
	left, right := 0, 0
	index := []int{}
	ret := 0
	for right < len(nums) {
		if nums[right] == 0 {
			index = append(index, right)
		}

		for len(index) > k {
			if nums[left] == 0 {
				index = index[1:]
			}
			left++
		}

		if ret < right-left+1 {
			ret = right - left + 1
		}
		right++

	}
	return ret
}

// func main() {
// 	nums := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
// 	longestOnes(nums, 2)
// }
