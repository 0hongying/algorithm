package main

// 给你一个无重复元素的整数数组candidates和一个目标整数target，找出candidates中可以使数字和为目标数target的所有不同组合
// candidates 中的数字可以无限制重复被选取。如果至少一个所选数字数量不同，则两种组合是不同的

func combinationSum(candidates []int, target int) [][]int {
	ret := [][]int{}
	temp := []int{}
	var dfs func(target, index int)
	dfs = func(target, index int) {
		if target == 0 {
			ret = append(ret, append([]int(nil), temp...))
			return
		}
		for i := index; i < len(candidates); i++ {
			if target-candidates[i] >= 0 {
				temp = append(temp, candidates[i])
				// i不加1，因为可以重复使用数字
				dfs(target-candidates[i], i)
				temp = temp[:len(temp)-1]
			}
		}
	}
	dfs(target, 0)
	return ret
}

// func main() {
// 	candidates := []int{2, 3, 6, 7}
// 	target := 7
// 	ret := combinationSum(candidates, target)
// 	fmt.Printf("%v", ret)
// }
