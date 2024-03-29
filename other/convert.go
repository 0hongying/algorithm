package main

import (
	"strings"
)

// 将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 N 字形排列。
// 比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
// P   A   H   N
// A P L S I I G
// Y   I   R
// 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。

// 找规律：
// 周期 = 2 * r - 2
// x = 字符串下标 % 周期
// y = Min(x, 周期-x)

func convert(s string, numRows int) string {
	// 周期
	n := numRows*2 - 2
	rowStr := make([]string, numRows)
	for i, c := range s {
		x := i % n
		rowStr[min(x, n-x)] += string(c)

	}
	return strings.Join(rowStr, "")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// func main() {
// 	s := "PAYPALISHIRING"
// 	numRows := 3
// 	ret := convert(s, numRows)
// 	fmt.Printf("%v", ret)
// }
