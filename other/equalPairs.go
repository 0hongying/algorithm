package main

import "fmt"

func equalPairs(grid [][]int) int {
	ans := 0
	m := map[string]int{}

	h := len(grid)

	for _, g := range grid {
		m[fmt.Sprint(g)]++
	}

	for i := 0; i < h; i++ {
		arr := []int{}
		for j := 0; j < h; j++ {
			arr = append(arr, grid[j][i])
		}
		if val, ok := m[fmt.Sprint(arr)]; ok {
			ans += val
		}
	}

	return ans
}

// func main() {
// 	grid := [][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}
// 	equalPairs(grid)
// }
