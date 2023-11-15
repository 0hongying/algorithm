package main

import (
	"sort"
)

// 给定一个可包含重复数字的序列nums，按任意顺序返回所有不重复的全排列
func permuteUnique(nums []int) [][]int {
	// 先排序
	sort.Ints(nums)
	path := make([]int, len(nums))
	visit := make([]bool, len(nums))
	ret := [][]int{}
	var dfs func(int)
	dfs = func(i int) {
		if i == len(nums) {
			ret = append(ret, append([]int(nil), path...))
			return
		}
		for j, v := range nums {
			// 数字重复 或者 被使用过
			if visit[j] || j > 0 && !visit[j-1] && v == nums[j-1] {
				continue
			}
			path[i] = nums[j]
			visit[j] = true
			dfs(i + 1)
			visit[j] = false
		}
	}
	dfs(0)
	return ret
}

// func main() {
// 	nums := []int{1, 1, 2}
// 	ret := permuteUnique(nums)
// 	fmt.Printf("%v", ret)
// }
