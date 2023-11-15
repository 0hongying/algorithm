package main

// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量
func numIslands(grid [][]byte) int {
	ret := 0
	row := len(grid)
	col := len(grid[0])

	var dfs func(i, j int)
	dfs = func(i, j int) {
		// 越界
		if i < 0 || i >= row || j < 0 || j >= col {
			return
		}
		if grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		dfs(i+1, j)
		dfs(i, j+1)
		dfs(i-1, j)
		dfs(i, j-1)
	}

	for i, g := range grid {
		for j, v := range g {
			if v == '1' {
				ret++
				dfs(i, j)
			}
		}
	}
	return ret
}
