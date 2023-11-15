 // 绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18
 function cuttingRope(n) {
    if (n <= 3) return n - 1
    let a = 1
    while (n > 4) {
        a = a * 3 % 1000000007
        n -= 3
    }
    return a * n % 1000000007
}

/**
 * 思路
 * 2 = 1 * 1
 * 3 = 1 * 2
 * 4 = 2 * 2
 * 5 = 2 * 3
 * 6 = 3 * 3
 * 
 * 由此可得：尽可能得将每段的长度切分为3时，乘积最大
 * 有三种情况
 * %3 === 0     ---> 3^n
 * %3 === 1     ---> 3^n-1*4
 * %3 === 2     ---> 3^n*2
 */


/**
 * 使用动态规划来解决 状态转移方程 = Max(dp[i], dp[i - j] * dp[j])
 */
function cuttingRopeV2(target) {
    const dp = [];
    dp[ 0 ] = 0;
    dp[ 1 ] = 1;

    for (let i = 2; i <= target; i++) {
        // 默认非 target 可以不切，即最大可能为本身
        dp[ i ] = i === target ? 1 : i;
        for (let j = 1; j < i; j++) {
            dp[ i ] = Math.max(dp[ i ], dp[ i - j] * dp[ j ]);
        }
        
    }
    return dp[ target ];
}

console.log(cuttingRopeV2(3));