package main

// 「快乐数」 定义为：
// 对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
// 然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
// 如果这个过程 结果为 1，那么这个数就是快乐数。
// 如果 n 是 快乐数 就返回 true ；不是，则返回 false

// func isHappy(n int) bool {
// 	m := map[int]int{}
// 	for n != 1 {
// 		if _, ok := m[n]; ok {
// 			return false
// 		}
// 		m[n]++
// 		numStr := strconv.Itoa(n)
// 		n = 0
// 		for _, v := range numStr {
// 			val, _ := strconv.Atoi(string(v))
// 			n += val * val
// 		}
// 	}
// 	return true
// }

func isHappy(n int) bool {
	slow, fast := n, step(n)
	for fast != 1 && slow != fast {
		// 慢指针走一步
		// 快指针走两步
		slow = step(slow)
		fast = step(step(fast))
	}
	return fast == 1
}

func step(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}
