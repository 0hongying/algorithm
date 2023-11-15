package main

// func jump(nums []int) int {
// 	left, right := 0, len(nums)
// 	ans := 0
// 	for left < right {
// 		max := math.MinInt32
// 		for i := 1; i < nums[left]; i++ {
// 			if max < nums[i+left] {
// 				max = nums[i+left]
// 				left = left + i
// 				ans++
// 			}
// 		}
// 	}
// 	return ans
// }

func jump(nums []int) int {
	ans := 0
	maxStep := 0
	end := 0
	for i := 0; i < len(nums)-1; i++ {
		v := nums[i]
		maxStep = max(maxStep, i+v)
		if i == end {
			end = maxStep
			ans++
		}
	}
	return ans
}

func main() {
	num := []int{2, 3, 1, 1, 4}
	jump(num)
}
