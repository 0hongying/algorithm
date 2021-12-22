// 寻找最大正方形面积
function maximalSquare(matrix) {
    let maxSide = 0
    if (matrix == null || matrix.length == 0 || matrix[0].length == 0) {
        return 0
    }
    let rows = matrix.length, columns = matrix[0].length
    // 二维数组dp
    let dp = new Array(rows)
    for (let i = 0; i < dp.length; i++) dp[i] = new Array(columns).fill(0)
    for (let i = 0; i < rows; i++) {
        for (let j = 0; j < columns; j++) {
            if (matrix[i][j] == '1') {
                if (i == 0 || j == 0) {
                    // 第一行和第一列最大面积为1
                    dp[i][j] = 1
                } else {
                    // 比较左边，上边，左上三者之间最小值+1 === 正方形的最大边长
                    dp[i][j] = Math.min(Math.min(dp[i - 1][j], dp[i][j - 1]), dp[i - 1][j - 1]) + 1
                }
                maxSide = Math.max(maxSide, dp[i][j]);
            }
        }
    }
    return maxSide * maxSide
}

/**
 * 思路
 * dp[i][j]表示第i行，第j列为右下角起点的最大正方形
 */