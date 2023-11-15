package main

// 生成有效括号
func generateParenthesis(n int) []string {
	s := ""
	ret := []string{}
	var dfs func(s string, left, right int)
	dfs = func(s string, left, right int) {
		// 左括号的数量小于右括号 或左括号的数量用完
		if left > right || left < 0 {
			return
		}
		if left == 0 && right == 0 {
			ret = append(ret, s)
			return
		}

		dfs(s+"(", left-1, right)
		dfs(s+")", left, right-1)

	}
	dfs(s, n, n)
	return ret
}

// func main() {
// 	ret := generateParenthesis(3)
// 	fmt.Printf("%v", ret)
// }
