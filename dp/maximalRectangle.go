package main

import "math"

// 给定一个由 0 和 1 组成的矩阵 matrix ，找出只包含 1 的最大矩形，并返回其面积。
func maximalRectangle(matrix []string) int {
	if len(matrix) == 0 {
		return 0
	}
	row := len(matrix)
	col := len(matrix[0])
	left := make([][]int, row)
	// 计算每个元素的左边连续1的个数
	for i, m := range matrix {
		left[i] = make([]int, col)
		for j, v := range m {
			if v == '0' {
				continue
			}
			if j > 0 {
				left[i][j] = left[i][j-1] + 1
			} else {
				left[i][j] = 1
			}
		}
	}

	ret := 0
	for i, m := range matrix {
		for j, v := range m {
			if v == '0' {
				continue
			}
			// 计算每个柱形图的面积
			// 面积 = min(宽度) * 高度
			width := math.MaxInt32
			area := math.MinInt32
			for k := i; k >= 0; k-- {
				width = min(width, left[k][j])
				area = max(area, width*(i-k+1))
			}
			ret = max(ret, area)
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

// func main() {
// 	matrix := []string{
// 		"10100", "10111", "11111", "10010",
// 	}
// 	maximalRectangle(matrix)
// }
