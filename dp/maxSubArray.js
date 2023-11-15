// 最大子串和
// 给定数组为 A = [1， 3， -2， 4， -5]， 则最大连续子序列和为 6，即 1 + 3 +（-2）+ 4 = 6
function maxSubArray(nums) {
    const dp = [];
    let max = nums[ 0 ];
    dp[ 0 ] = nums[ 0 ];
    for (let i = 1; i < nums.length; i++) {
        dp[ i ] = Math.max(nums[ i ], dp[ i - 1] + nums[ i ]);
        max = Math.max(max, dp[i]);
    }
    return max;
}

// dp[i]表示以i结尾的最大和
// dp[i] = Max(nums[i], d[i-1] + nums[i])
const nums = [ 1, 3, -2, 4, -5 ];
const ret = maxSubArray(nums);
console.log(ret);

