package main

// 给你一个整数数组 nums ，返回全部为 0 的 子数组 数目
// [1,3,0,0,2,0,0,4]
// 子数组 [0] 出现了 4 次。
// 子数组 [0,0] 出现了 2 次
// 不存在长度大于 2 的全 0 子数组，所以我们返回 6

func zeroFilledSubarray(nums []int) int64 {
	count := 0
	ret := 0
	for _, num := range nums {
		if num == 0 {
			count++
			ret += count
		} else {
			count = 0
		}
	}
	return int64(ret)
}
