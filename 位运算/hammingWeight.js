// 返回其二进制表达式中数字位数为 '1' 的个数
var hammingWeight = function(n) {
    let ret = 0
    while (n) {
        // 去除最低位1
        n = n & (n-1)
        ret ++
    }
    return ret 
}

/**
 * 思路
 * 一个数&（该数-1）可以消去该数二进制最低位的1
 */