package main

// 给你两个字符串 start 和 target ，长度均为 n 。每个字符串 仅 由字符 'L'、'R' 和 '_' 组成，其中：
// 字符 'L' 和 'R' 表示片段，其中片段 'L' 只有在其左侧直接存在一个 空位 时才能向 左 移动，而片段 'R' 只有在其右侧直接存在一个 空位 时才能向 右 移动。
// 字符 '_' 表示可以被 任意 'L' 或 'R' 片段占据的空位。
// 如果在移动字符串 start 中的片段任意次之后可以得到字符串 target ，返回 true ；否则，返回 false

func canChange(start string, target string) bool {
	n := len(start)
	a := []int{}
	b := []int{}
	for i := 0; i < n; i++ {
		if start[i] != '_' {
			a = append(a, i)
		}
		if target[i] != '_' {
			b = append(b, i)
		}
	}
	// LR的数量不一致
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if start[a[i]] != target[b[i]] {
			// 字符不一致
			return false
		}
		if start[a[i]] == 'L' && a[i] < b[i] {
			// L只能向左移动
			return false
		}
		if start[a[i]] == 'R' && a[i] > b[i] {
			// R只能向右移动
			return false
		}
	}
	return true

}
