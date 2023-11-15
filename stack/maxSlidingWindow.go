package main

func maxSlidingWindow(nums []int, k int) []int {
	ans := []int{}
	queue := []int{}
	for i, v := range nums {
		for len(queue) > 0 && nums[queue[len(queue)-1]] < v {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)

		if i-queue[0] >= k {
			queue = queue[1:]
		}

		if i >= k-1 {
			ans = append(ans, nums[queue[0]])
		}
	}
	return ans
}
