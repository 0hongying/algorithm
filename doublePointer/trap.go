package main

// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水

func trap(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	ans := 0

	for left < right {
		// 左边界最大
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])

		if leftMax < rightMax {
			ans += leftMax - height[left]
			left++
		} else {
			ans += rightMax - height[right]
			right--
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
