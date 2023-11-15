package main

import "sort"

// 给你一个整数数组 nums，请你找出并返回能被三整除的元素最大和。

func maxSumDivThree(nums []int) int {
	v := make([][]int, 3)
	for _, num := range nums {
		v[num%3] = append(v[num%3], num)
	}

	// 升序
	sort.Slice(v[1], func(i, j int) bool {
		return v[1][i] > v[1][j]
	})
	sort.Slice(v[2], func(i, j int) bool {
		return v[2][i] > v[2][j]
	})

	ans, lb, lc := 0, len(v[1]), len(v[2])
	// lb-2 或 lc-2 是从后面开始枚举
	// cntb的取值范围为[n-2, n]
	// cntc的取值范围为[m-2, m]
	for cntb := max(lb-2, 0); cntb <= lb; cntb++ {
		for cntc := max(lc-2, 0); cntc <= lc; cntc++ {
			// (n + m * 2) % 3 (n：余数为1的个数，m：余数为2的个数)
			// (n + m * 2) % 3 = (n + m * 3 - m) % 3 = (n - m) % 3
			if (cntb-cntc)%3 == 0 {
				ans = max(ans, sum(v[1][:cntb])+sum(v[2][:cntc]))
			}
		}
	}

	return sum(v[0]) + ans
}

func sum(v []int) int {
	ans := 0
	for _, x := range v {
		ans += x
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func main() {
// 	nums := []int{3, 6, 7, 1, 10, 13}
// 	maxSumDivThree(nums)
// }
