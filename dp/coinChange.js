// 换硬币(返回可以凑成总金额所需的, 最少的硬币个数)
function coinChange(coins, amount) {
    let dp = new Array(amount + 1).fill(Infinity)
    dp[0] = 0
    for (let i = 1; i <= amount; i++) {
        for (let coin of coins) {
            if (i >= coin) {
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

// dp[ i ] 表示最小的兑换硬币个数
// dp[ i ] = Min(dp[ amount - coins[n] ] + 1, dp[ i ])
const coins = [ 1, 2, 4 ];
const amount = 11;
const ret = coinChange(coins, amount);
console.log(ret);
