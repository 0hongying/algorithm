package main

// 给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的连续子数组的个数 。
// 子数组是数组中元素的连续非空序列

func subarraySum(nums []int, k int) int {
	sum := 0
	ans := 0
	for i, v := range nums {
		sum += v
		if k == sum {
			ans++
		}
		temp := sum
		for j := 0; j < i; j++ {
			temp -= nums[j]
			if temp == k {
				ans++
			}
		}
	}
	return ans
}

// func main() {
// 	k := 3
// 	nums := []int{1, 2, 3}
// 	subarraySum(nums, k)
// }
