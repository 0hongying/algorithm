// 在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
// 判断数组中是否含有该整数
function findNumberIn2DArray(matrix, target) {
    if (matrix == null || matrix.length == 0 || matrix[0].length == 0) return false
    const m = matrix[0].length-1
    const n = matrix.length-1
    let row = n
    let column = 0
    // 从左下角开始找
    while (row >=0 && column <= m) {
        if (matrix[row][column] === target) {
            return true
        } else if (matrix[row][column] < target) {
            // 往右寻找
            column++
        } else {
            // 往上寻找
            row--
        } 
    }
    return false
}

/**
 * 思路
 * 从左下角开始找，比target小则向右找，反之向上找
 */