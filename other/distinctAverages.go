package main

import "sort"

func distinctAverages(nums []int) int {
	m := map[float32]int{}
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	for left < right {
		v := (float32(nums[left]) + float32(nums[right])) / 2
		m[v]++
		left++
		right--
	}
	return len(m)
}

// func main() {
// 	nums := []int{4, 1, 4, 0, 3, 5}
// 	distinctAverages(nums)
// }
