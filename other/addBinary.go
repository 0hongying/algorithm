package main

import "strconv"

func addBinary(a string, b string) string {
	ans := ""
	temp := 0
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0 || temp != 0; {
		num1 := 0
		num2 := 0
		if i >= 0 {
			num1 = int(a[i] - '0')
		}
		if j >= 0 {
			num2 = int(b[j] - '0')
		}

		sum := temp + num1 + num2
		ans = strconv.Itoa(sum%2) + ans
		temp = sum / 2

		i--
		j--
	}
	return ans
}
