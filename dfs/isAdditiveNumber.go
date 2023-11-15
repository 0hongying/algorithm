package main

import "strconv"

func addString(a, b string) string {
	add := 0
	sum := 0
	ret := ""
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
		n, m := 0, 0
		if i >= 0 {
			n = int(a[i] - '0')
		}
		if j >= 0 {
			m = int(b[j] - '0')
		}
		sum = n + m + add
		ret = strconv.Itoa(sum%10) + ret
		add = sum / 10
	}
	return ret
}

func dfs(num string, first, second int) bool {
	n, last := len(num), 0
	for second < n {
		if (num[last] == '0' && first > last+1) || (num[first] == '0' && second > first+1) {
			return false
		}
		s := addString(num[last:first], num[first:second])
		if second+len(s) > n || num[second:second+len(s)] != s {
			return false
		}
		last, first, second = first, second, second+len(s)
	}
	return true
}

// 一个有效的 累加序列 必须 至少 包含 3 个数。除了最开始的两个数以外，序列中的每个后续数字必须是它之前两个数字之和。
// 给你一个只包含数字 '0'-'9' 的字符串，编写一个算法来判断给定输入是否是 累加数 。如果是，返回 true ；否则，返回 false

func isAdditiveNumber(num string) bool {
	for i := 1; i < len(num)-1; i++ {
		for j := i + 1; j < len(num); j++ {
			if dfs(num, i, j) {
				return true
			}
		}
	}
	return false
}
