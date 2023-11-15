// 判断是否为回文数字
function isPalindrome(x) {
    if (x < 0) return false
    let ret = 0, y = x
    while (y) {
        ret = y % 10 + ret * 10 
        y = parseInt(y / 10)
    }
    return ret === x
}