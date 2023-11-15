package main

type pair struct {
	x, y int
}

var directions = []pair{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
} // 上下左右

// 给定一个m x n二维字符网格board和一个字符串单词word。如果 word 存在于网格中，返回 true；否则，返回 false
// ['A','B','C','E']
// ['S','F','C','S']
// ['A','D','E','E']
// word = 'ABCCED'
// true
func exist(board [][]byte, word string) bool {
	row, col := len(board), len(board[0])
	vis := make([][]bool, row)
	for i := range vis {
		vis[i] = make([]bool, col)
	}

	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		if board[i][j] != word[k] {
			return false
		}

		// 终止条件
		if k == len(word)-1 {
			return true
		}

		// 标记
		vis[i][j] = true
		// 回溯
		defer func() {
			vis[i][j] = false
		}()

		for index := 0; index < len(directions); index++ {
			// 寻找上，下，左，右 四个方向中符合条件的
			if newI, newJ := i+directions[index].x, j+directions[index].y; newI >= 0 && newI < row && newJ >= 0 && newJ < col && !vis[newI][newJ] {
				if dfs(newI, newJ, k+1) {
					return true
				}
			}
		}
		return false
	}

	for i := range board {
		for j := range board[i] {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}

// func main() {
// 	board := [][]byte{
// 		{'C', 'A', 'A'},
// 		{'A', 'A', 'A'},
// 		{'B', 'C', 'D'},
// 	}
// 	word := "AAB"
// 	ret := exist(board, word)
// 	fmt.Printf("%v", ret)
// }
