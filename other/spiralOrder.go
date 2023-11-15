package main

func spiralOrder(matrix [][]int) []int {
	ans := []int{}
	left, right := 0, len(matrix[0])
	up, down := 0, len(matrix)
	for true {
		for i := left; i < right; i++ {
			ans = append(ans, matrix[up][i])
		}
		up++
		if up > down {
			break
		}

		for i := up; i < down; i++ {
			ans = append(ans, matrix[i][right-1])
		}
		right--
		if left > right {
			break
		}

		for i := right - 1; i >= left; i-- {
			ans = append(ans, matrix[down-1][i])
		}
		down--
		if up > down {
			break
		}

		for i := down - 1; i >= up; i-- {
			ans = append(ans, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}
	return ans
}

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	spiralOrder(matrix)
}
