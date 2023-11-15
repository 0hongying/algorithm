package main

func generate(numRows int) [][]int {
	ans := [][]int{}
	for i := 0; i < numRows; i++ {
		arr := []int{}
		for j := 0; j <= i; j++ {
			num := 1
			if j != 0 && j != i && i > 0 && len(ans[i-1]) >= 2 {
				num = ans[i-1][j-1] + ans[i-1][j]
			}
			arr = append(arr, num)
		}
		ans = append(ans, arr)
	}
	return ans
}

// func main() {
// 	generate(6)
// }
