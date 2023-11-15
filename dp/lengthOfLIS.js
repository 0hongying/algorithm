// 求最长递增子序列（可以不连续）
function lengthOfLIS(nums) {
    const dp = []
    let maxans = 1
    dp[0] = 1
    for (let i = 1; i < nums.length; i++) {
        // 默认子序列中只有nums[i]
        dp[i] = 1
        for (let j = 0; j < i; j++) {
            // 只有小于nums[i]才是递增
            if (nums[j] < nums[i]) {
                dp[i] = Math.max(dp[i], dp[j]+1)
            }
        }
        maxans = Math.max(maxans, dp[i])
    }
    return maxans
}

// 0, 1, 2, 3
// dp[i]表示以nums[i]为结尾的最长递增子序列（默认为1，即只有nums[i]）
const nums = [ 0, 1, 0, 3, 2, 3 ];
const ret = lengthOfLIS(nums);
console.log(ret);