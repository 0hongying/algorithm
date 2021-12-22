// 给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，计算其二进制数中的 1 的数目并将它们作为数组返回
var countBits = function(nums) {
    const ret = []
    ret.push(0)
    for (let i = 1; i <= nums; i++) {
        if (i % 2 === 0) {
            ret[i] = ret[i/2]
        } else {
            ret[i] = ret[i-1]+1
        }
    }
    return ret
}

/**
 * 思路
 * 3（101）和6（1010）二进制所含1的数目是一致的
 * 3（101）和2（10）二进制所含1的数目相差1
 */