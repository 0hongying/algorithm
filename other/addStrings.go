package main

import (
	"strconv"
)

// 给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和并同样以字符串形式返回
func addStrings(num1 string, num2 string) string {
	add := 0
	ans := ""
	// add != 0 是保证最后的进位
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0')
		}

		if j >= 0 {
			y = int(num2[j] - '0')
		}

		ret := x + y + add
		// ans要放后面
		ans = strconv.Itoa(ret%10) + ans
		// 进位
		add = ret / 10
	}
	return ans
}

// func main() {
// 	num1 := "83"
// 	num2 := "1123"
// 	ret := addStrings(num1, num2)
// 	fmt.Printf("%v", ret)
// }
