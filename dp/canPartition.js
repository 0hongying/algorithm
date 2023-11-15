var canPartition = function(nums) {
    let sum = 0
    for (const n of nums) {
        sum += n
    }

    if (sum % 2 != 0) {
        return false
    }

    const target = sum / 2
    const dp = new Array(target+1).fill(false)
    dp[0] = true
    for (let i = 1; i <= target; i++){
        for (const n of nums) {
            if (i >= n) {
                dp[i] = dp[i] || dp[i-n]
            }
        }
    }

    return dp[target]
};

canPartition([3,3,3,4,5])