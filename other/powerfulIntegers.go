package main

// 给定三个整数 x 、 y 和 bound ，返回 值小于或等于 bound 的所有 强整数 组成的列
// 如果某一整数可以表示为 xi + yj ，其中整数 i >= 0 且 j >= 0，那么我们认为该整数是一个 强整数

func powerfulIntegers(x int, y int, bound int) []int {
	ret := []int{}
	m := map[int]struct{}{}
	for i := 1; i <= bound; i *= x {
		for j := 1; i+j <= bound; j *= y {
			m[i+j] = struct{}{}
			if y == 1 {
				break
			}
		}
		if x == 1 {
			break
		}
	}

	for k, _ := range m {
		ret = append(ret, k)
	}
	return ret
}
