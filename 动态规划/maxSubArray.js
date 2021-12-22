// 最大子串和
function maxSubArray(nums) {
    let pre = 0, maxAns = -Infinity
    nums.forEach((x) => {
        pre = Math.max(pre + x, x)
        maxAns = Math.max(maxAns, pre)
    })
    return maxAns
}