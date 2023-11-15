package main

import "strconv"

func divide(dividend int, divisor int) int {
	sign := true
	if dividend > 0 && divisor < 0 || dividend < 0 && divisor > 0 {
		sign = false
	}

	dividend = abs(dividend)
	divisor = abs(divisor)

	if dividend < divisor {
		return 0
	}

	dividendStr := string(dividend)
	left := 0
	ans := ""

	for i := 0; i < len(dividendStr); i++ {
		left += int(dividendStr[i] - '0')
		if left < divisor {
			ans += "0"
			continue
		} else {
			temp := 0
			for left >= divisor {
				left -= dividend
				temp++
			}
			ans += string(temp)
		}
	}

	ret, _ := strconv.Atoi(ans)
	if sign {
		return ret
	} else {
		return -ret
	}
}

// func main() {
// 	dividend := 10
// 	divisor := 3

// 	ret :=
// }
