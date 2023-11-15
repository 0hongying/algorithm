package main

import "fmt"

// 给定一个不含重复数字的数组 nums ，返回其所有可能的全排列
// func permute(nums []int) [][]int {
// 	visit := make([]bool, len(nums))
// 	ret := [][]int{}
// 	path := make([]int, len(nums))
// 	var dfs func(int)
// 	dfs = func(i int) {
// 		if i == len(nums) {
// 			ret = append(ret, append([]int{}, path...))
// 			return
// 		}
// 		for j, ok := range visit {
// 			if !ok {
// 				path[i] = nums[j]
// 				visit[j] = true
// 				dfs(i + 1)
// 				// 回溯
// 				visit[j] = false
// 			}
// 		}
// 	}
// 	dfs(0)
// 	return ret
// }

func permute(nums []int) [][]int {
	ans := [][]int{}

	vis := make([]bool, len(nums))
	path := make([]int, len(nums))

	var dfs func(index int)
	dfs = func(index int) {
		if index == len(nums) {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if !vis[i] {
				vis[i] = true
				path[index] = nums[i]
				dfs(index + 1)
				vis[i] = false
			}
		}
	}
	dfs(0)
	return ans
}

func main() {
	nums := []int{5, 4, 6, 2}
	ret := permute(nums)
	fmt.Printf("%v", ret)
}
