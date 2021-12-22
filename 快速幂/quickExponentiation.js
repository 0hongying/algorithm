function quickExponentiation(a, n) {
    if (n === 0) {
        return 1
    } else if (n & 1) {
        // 奇数
        return quickExponentiation(a, n-1) * a
    } else {
        // 偶数
        const temp = quickExponentiation(a, n/2)
        return temp * temp
    }
}

// 非递归
function quickExponentiation1(a, n) {
    let ret = 1
    while (n) {
        if (n & 1) {
            // 奇数
            ret *= a
        }
        a *= a
        n >>= 1
    }
    return ret
}

/**
 * 思路
 * 2 * 2 * 2 * 2 * 2 = 4 * 4 * 2
 * 偶数次幂等于(底数*底数)^n/2
 * 奇数次幂等于底数^n-1*底数
 */