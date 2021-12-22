// 汉诺塔
function hanoi(ret, n, x, y, z) {
    if (n === 1) {
        ret.push(`${x}移动${z}`)
    } else {
        // 把x上的圆盘借助z移动到y
        hanoi(ret, n-1, x, z, y)
        ret.push(`${x}移动${z}`)
        // 把y上的圆盘借助x移动到z
        hanoi(ret, n-1, y, x, z)
    }
}
