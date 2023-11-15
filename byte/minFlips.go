package main

// 给你三个正整数 a、b 和 c。
// 你可以对 a 和 b 的二进制表示进行位翻转操作，返回能够使按位或运算   a OR b == c  成立的最小翻转次数。
// 「位翻转操作」是指将一个数的二进制表示任何单个位上的 1 变成 0 或者 0 变成 1

func minFlips(a int, b int, c int) int {
	x, y, z := 0, 0, 0
	ret := 0
	for c > 0 || a > 0 || b > 0 {
		x = a & 1
		y = b & 1
		z = c & 1
		if z == 0 {
			if x == 1 {
				ret++
			}
			if y == 1 {
				ret++
			}
		} else {
			if x == 0 && y == 0 {
				ret++
			}
		}

		a = a >> 1
		b = b >> 1
		c = c >> 1
	}
	return ret
}
