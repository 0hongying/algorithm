package main

func largestRectangleArea(heights []int) int {
	ans := 0
	for i, v := range heights {
		ans = max(v, ans)
		num := v
		for j := i + 1; j < len(heights); j++ {
			num = min(heights[j], num)
			ans = max(ans, num*(j-i+1))
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

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// func main() {
// 	heights := []int{2, 1, 2}
// 	largestRectangleArea(heights)
// }
