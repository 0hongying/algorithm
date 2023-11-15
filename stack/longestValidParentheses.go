package main

// 给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度

func longestValidParentheses(s string) int {
	stack := []int{-1}
	ans := 0
	for i, v := range s {
		if v == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				index := stack[len(stack)-1]
				ans = max(ans, i-index)
			} else {
				stack = append(stack, i)
			}
		}
	}
	return ans
}

// func main() {
// 	s := "(()"
// 	longestValidParentheses(s)
// }
