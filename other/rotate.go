package main

func rotate(nums []int, k int) {
	ret := []int{}
	ret = append(ret, nums[len(nums)-k:]...)
	ret = append(ret, nums[:len(nums)-k]...)
	nums = ret
}

// func main() {
// 	nums := []int{-1, -100, 3, 99}
// 	k := 2
// 	rotate(nums, k)
// }
