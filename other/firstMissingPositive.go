package main

// 给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。

func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); {
		v := nums[i]
		if v > 0 && v < len(nums) && v != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		} else {
			i++
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return len(nums) + 1
}

// func main() {
// 	nums := []int{2, 1}
// 	ret := firstMissingPositive(nums)
// 	fmt.Printf("%v", ret)
// }
