// 约瑟夫环
function josephus(n, m) {
    // f(1) = 0, 从后往前推
    let ret = 0
    for (let i = 2; i <= n; i++) {
        ret = (ret + m) % i
    }
    return ret
}
