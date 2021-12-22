function gcd(n, m) {
    let temp
    while (m !== 0) {
        // 保存除数
        temp = m
        m = n % m
        n = temp
    }
    return n
}

/**
 * 思路
 * 32 % 24 = 1 ...... 8
 * 24 % 8 = 3 ...... 0
 */