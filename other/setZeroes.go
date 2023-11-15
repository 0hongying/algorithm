package main

// 给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0
func setZeroes(matrix [][]int) {
	row := len(matrix)
	col := len(matrix[0])

	rowBool := make([]bool, row)
	colBool := make([]bool, col)

	// 先把原数组的 0 作标记
	for i, m := range matrix {
		for j, v := range m {
			if v == 0 {
				rowBool[i] = true
				colBool[j] = true
			}
		}
	}

	for i, m := range matrix {
		for j, _ := range m {
			if rowBool[i] || colBool[j] {
				matrix[i][j] = 0
			}
		}
	}
}

func setZeroes2(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])
	row0, col0 := false, false
	// 判断原数组的行是否有 0
	for _, v := range matrix[0] {
		if v == 0 {
			row0 = true
			break
		}
	}
	// 判断原数组的列是否有 0
	for _, r := range matrix {
		if r[0] == 0 {
			col0 = true
			break
		}
	}
	// 修改原数组
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	// 如果原数组的第一行或第一列有 0
	if row0 {
		for j := 0; j < m; j++ {
			matrix[0][j] = 0
		}
	}
	if col0 {
		for _, r := range matrix {
			r[0] = 0
		}
	}
}
