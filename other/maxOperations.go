package main

//给你一个整数数组 nums 和一个整数 k 。
// 每一步操作中，你需要从数组中选出和为 k 的两个整数，并将它们移出数组。
// 返回你可以对数组执行的最大操作数

func maxOperations(nums []int, k int) int {
	m := map[int]int{}
	ret := 0
	for _, v := range nums {
		if m[k-v] > 0 {
			m[k-v]--
			ret++
		} else {
			m[v]++
		}
	}
	return ret
}

// func main() {
// 	nums := []int{2, 5, 4, 4, 1, 3, 4, 4, 1, 4, 4, 1, 2, 1, 2, 2, 3, 2, 4, 2}
// 	k := 3
// 	maxOperations(nums, k)
// }
