package main

// 给你一个包含若干星号 * 的字符串 s 。
// 在一步操作中，你可以：
// 选中 s 中的一个星号。
// 移除星号 左侧 最近的那个 非星号 字符，并移除该星号自身。
// 返回移除 所有 星号之后的字符

func removeStars(s string) string {
	stack := []rune{}
	for _, c := range s {
		if c != '*' {
			stack = append(stack, c)
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return string(stack)
}
