package main

// 给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等
func canPartition(nums []int) bool {

	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}

	target := sum / 2

	dp := make([][]bool, len(nums))
	for i, _ := range dp {
		dp[i] = make([]bool, target+1)
		// j - num[i] = 0 成立， 说明 num[i] 这个数就恰好能够在被分割为单独的一组，其余的数分割成为另外一组
		dp[i][0] = true
	}

	// dp[i][j]前i个数字的和为j
	for i := 1; i < len(nums); i++ {
		for j := 1; j <= target; j++ {
			if j >= nums[i] {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[len(nums)-1][target]
}

// func main() {
// 	nums := []int{1, 5, 11, 5}
// 	ret := canPartition(nums)
// 	fmt.Printf("%v", ret)
// }
