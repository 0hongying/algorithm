// 0 1 1 2 3 5 8 13 21
function fibonacci(n) {
    let x = 0, y = 1, z = 0
    for (let i = 1; i <= n; i++) {
        z = x + y
        x = y
        y = z
    }
    return z
}

// 1 2 3 5 8 13 21
function fibonacci1(n) {
    let x = 1, y = 2, z = 0
    for (let i = 3; i <= n; i++) {
        z = x + y
        x = y
        y = z
    }
    return z
}

/**
 * 思路
 * 先把0 1 1 2 3 5 8 13 21完整列出来
 * x代表第一个数，y代表第二个数，z代表两者之和
 */