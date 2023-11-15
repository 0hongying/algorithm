package main

// https://leetcode.cn/problems/koko-eating-bananas/description/?envType=study-plan-v2&envId=leetcode-75

func minEatingSpeed(piles []int, h int) int {
	left, right := 1, 1000000000

	helped := func(x int) int {
		sum := 0
		for _, v := range piles {
			sum += (v + x - 1) / x
		}
		return sum
	}

	for left < right {
		mid := (left + right) / 2
		if helped(mid) > h {
			left = mid + 1
		} else if helped(mid) == h {
			right = mid
		} else {
			right = mid
		}
	}

	return left
}

// func main() {
// 	piles := []int{312884470}
// 	h := 312884469
// 	minEatingSpeed(piles, h)
// }
