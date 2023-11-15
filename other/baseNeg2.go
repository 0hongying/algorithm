package main

// 给你一个整数 n ，以二进制字符串的形式返回该整数的 负二进制（base -2）表示。

func baseNeg2(n int) string {
	if n == 0 {
		return "0"
	}
	ans := []byte{}
	k := 1
	for n != 0 {
		if n%2 != 0 {
			ans = append(ans, '1')
			n -= k
		} else {
			ans = append(ans, '0')
		}
		k *= -1
		n /= 2
	}
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return string(ans)
}

// func main() {
// 	n := 10
// 	ret := baseNeg2(n)
// 	fmt.Printf("%v", ret)
// }

// 1 + (-2) + 4 + -8
