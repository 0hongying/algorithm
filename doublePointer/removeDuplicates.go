package main

// 给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使得出现次数超过两次的元素只出现两次，返回删除后数组的新长度。
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	// fast进行数组遍历
	// slow指向有效数组的最后一位
	slow := 0
	for _, v := range nums {
		// slow < 2 前k位直接保留
		// 与前面k个元素不一致，则保留
		if slow < 2 || v != nums[slow-2] {
			nums[slow] = v
			slow++
		}
	}
	return slow
}

// func removeDuplicates2(nums []int) int {
// 	var process func(k int) int
// 	process = func(k int) int {
// 		u := 0
// 		for _, v := range nums {
// 			if u < k || nums[u-k] != v {
// 				nums[u] = v
// 				u++
// 			}
// 		}
// 		return u
// 	}
// 	return process(2)
// }

// func main() {
// 	nums := []int{
// 		1, 2, 2, 2, 2, 2, 3,
// 	}
// 	ret := removeDuplicates(nums)
// 	fmt.Printf("%v", ret)
// }
