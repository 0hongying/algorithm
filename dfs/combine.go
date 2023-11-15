package main

// 给定两个整数n和k，返回范围[1, n]中所有可能的k个数的组合
func combine(n int, k int) [][]int {
	ret := [][]int{}
	comb := []int{}

	var dfs func(index int)
	dfs = func(index int) {
		if len(comb) == k {
			ret = append(ret, append([]int(nil), comb...))
			return
		}

		for i := index; i <= n; i++ {
			comb = append(comb, i)
			dfs(i + 1)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(1)
	return ret
}

// func main() {
// 	n, k := 4, 2
// 	ret := combine(n, k)
// 	fmt.Printf("%v", ret)
// }
