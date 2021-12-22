// 换硬币
function coinChange(coins, amount) {
    let dp = new Array(amount + 1).fill(Infinity)
    dp[0] = 0
    for (let i = 1; i <= amount; i++) {
        for (let coin of coins) {
            if (i - coin >= 0) {
                // 金额必须大于硬币额度
                dp[i] = Math.min(dp[i], dp[i - coin] + 1);
            }
        }
    }
    return dp[amount] === Infinity ? -1 : dp[amount]
}

/**
 * 思路
 * min(F(amount-coins[0])+1, F(amount-coins[n])+1)
 */