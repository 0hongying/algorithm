// 丑数就是只包含质因数 2、3 、5 的正整数
function isUgly(n) {
    if (n <= 0) return false
    for (const num of [2, 3, 5]) {
        while (n % num === 0) {
            n /= num
        }
    }
    return n === 1
}