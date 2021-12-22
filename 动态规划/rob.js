// 打家劫舍（相邻只能偷一个）
function rob1(nums) {
    if (nums.length === 1 || nums.length === 2) return Math.max(...nums)
    let dp = []
    dp[0] = nums[0]
    dp[1] = Math.max(nums[0], nums[1])
    for (let i = 2; i < nums.length; i++) {
        // 隔一家偷
        dp[i] = Math.max(dp[i-2]+ nums[i], dp[i-1])
    }
    return dp[nums.length-1]
}

// 打家劫舍（相邻只能偷一个且头尾也算相邻）
function rob2(nums) {
    if (nums.length === 1 || nums.length === 2) return Math.max(...nums)
    let dp = []
    let ret1, ret2
    // 偷第一家
    dp[0] = nums[0]
    dp[1] = nums[0]
    // 因为偷了第一家，所有最后一家不能再偷了
    for (let i = 2; i < nums.length -1; i++) {
        dp[i] = Math.max(dp[i-2]+ nums[i], dp[i-1])
    }
    ret1 = dp[dp.length-1]
    // 不偷第一家
    dp[0] = 0
    dp[1] = nums[1]
    // 第一家没偷，所以最后一家可以偷
    for (let i = 2; i < nums.length; i++) {
        dp[i] = Math.max(dp[i-2]+ nums[i], dp[i-1])
    }
    ret2 = dp[dp.length-1]
    return Math.max(ret1, ret2)
}
