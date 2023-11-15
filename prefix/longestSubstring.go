package main

// 给你一个二元数组 nums ，和一个整数 goal ，请你统计并返回有多少个和为 goal 的 非空 子数组。
// 子数组 是数组的一段连续部分

func numSubarraysWithSum(nums []int, goal int) int {
	m := map[int]int{}
	sum := 0
	ans := 0
	for _, v := range nums {
		m[sum]++
		sum += v
		ans += m[sum-goal]
	}
	return ans
}

// func main() {
// 	num := []int{1, 2, 3}
// 	ret := numSubarraysWithSum(num, 3)
// 	fmt.Printf("%v", ret)
// }
