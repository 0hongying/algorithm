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

const ret = quickExponentiation(3, 3);
console.log(ret);

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

const ret = quickExponentiation1(3, 3);
console.log(ret);