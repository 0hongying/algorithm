package main

// 给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
func isSubsequence(s string, t string) bool {
	n, m := len(s), len(t)
	i, j := 0, 0
	// s 或 t 其中一个达到尾部
	for i < n && j < m {
		if s[i] == t[j] {
			// 相等则移动 i 指针
			i++
		}
		j++
	}
	return i == n
}
