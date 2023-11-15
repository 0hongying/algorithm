package main

// 给你一个数组 candies 和一个整数 extraCandies ，其中 candies[i] 代表第 i 个孩子拥有的糖果数目。
// 对每一个孩子，检查是否存在一种方案，将额外的 extraCandies 个糖果分配给孩子们之后，此孩子有 最多 的糖果。注意，允许有多个孩子同时拥有 最多 的糖果数目

func kidsWithCandies(candies []int, extraCandies int) []bool {
	m := 0
	for _, v := range candies {
		m = max(m, v)
	}
	ret := make([]bool, len(candies))

	for i, v := range candies {
		if v == m || v+extraCandies >= m {
			ret[i] = true
		}
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
