// 绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18
function cuttingRope(n) {
    if (n <= 3) return n-1
    let m = parseInt(n/3)
    if (n % 3 === 0) return Math.pow(3, m)
    else if (n % 3 === 1) return Math.pow(3, m-1) * 4
    else return Math.pow(3, m) * 2
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