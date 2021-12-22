// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素
function singleNumber(arr) {
    let ret = 0
    for (const num of arr) {
        ret ^= num
    }
    return ret
}

/**
 * 思路
 * A ^ A ^ B = 0 ^ B = B 
 */