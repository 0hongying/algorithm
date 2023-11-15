package main

// 给你一个整数数组nums，数组中的元素互不相同 。返回该数组所有可能的子集
// nums = [1,2,3]
// [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
func subsets(nums []int) [][]int {
	ret := [][]int{}
	comb := []int{}

	var dfs func(index int)
	dfs = func(index int) {
		ret = append(ret, append([]int(nil), comb...))
		// 终止条件
		if len(nums) == index {
			return
		}

		for i := index; i < len(nums); i++ {
			comb = append(comb, nums[i])
			dfs(i + 1)
			// 回溯
			comb = comb[:len(comb)-1]
		}
	}
	dfs(0)
	return ret
}

// func main() {
// 	nums := []int{1, 2, 3}
// 	ret := subsets(nums)
// 	fmt.Printf("%v", ret)
// }
