package main

import (
	"sort"
)

func canPartitionKSubsets(nums []int, k int) bool {
	var backTrace func(buckets []int, start, target int) bool
	backTrace = func(buckets []int, start, target int) bool {
		if start < 0 {
			return true
		}
		num := nums[start]
		for i := 0; i < len(buckets); i++ {
			if num+buckets[i] <= target {
				buckets[i] += num
				if backTrace(buckets, start-1, target) {
					return true
				}
				buckets[i] -= num
			}
			if buckets[i] == 0 {
				break
			}
		}
		return false
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}

	sort.Ints(nums)
	if sum%k != 0 || sum/k < nums[0] {
		return false
	}
	if nums[len(nums)-1] > sum/k {
		return false
	}
	return backTrace(make([]int, k), len(nums)-1, sum/k)
}

// func main() {
// 	nums := []int{1, 1, 1, 1, 2, 2, 2, 2}
// 	k := 4
// 	ret := canPartitionKSubsets(nums, k)
// 	fmt.Printf("%v", ret)
// }
