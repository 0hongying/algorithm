package main

// 给你一个满足下述两条属性的 m x n 整数矩阵：
// 每行中的整数从左到右按非严格递增顺序排列。
// 每行的第一个整数大于前一行的最后一个整数。
// 给你一个整数 target ，如果 target 在矩阵中，返回 true ；否则，返回 false

// func searchMatrix(matrix [][]int, target int) bool {
// 	n := len(matrix[0])
// 	m := len(matrix)
// 	left, right := 0, m*n

// 	for left < right {
// 		mid := (left + right) / 2
// 		x := mid / n
// 		y := mid % n
// 		if matrix[x][y] < target {
// 			left = mid + 1
// 		} else if matrix[x][y] == target {
// 			return true
// 		} else {
// 			right = mid
// 		}
// 	}
// 	return false
// }

func searchMatrix(matrix [][]int, target int) bool {
	left, right := 0, len(matrix)
	n := len(matrix[0])
	k := 0

	for left < right {
		mid := (left + right) / 2
		if matrix[mid][n-1] < target {
			k = left
			left = mid + 1
		} else if matrix[mid][n-1] == target {
			return true
		} else {
			right = mid
		}
	}

	left, right = 0, len(matrix[0])
	for left < right {
		mid := (left + right) / 2
		if matrix[k][mid] < target {
			left = mid + 1
		} else if matrix[k][mid] == target {
			return true
		} else {
			right = mid
		}
	}
	return false
}

func main() {
	target := 10
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	searchMatrix(matrix, target)
}
