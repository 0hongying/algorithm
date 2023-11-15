// 杨辉三角
var triangle = function(numRows) {
    const ret = []
    for (let i = 0; i < numRows; i++) {
        ret.push([])
        for (let j = 0; j <= i; j++) {
            if (j === 0 || j === i) {
                // 第一位和最后一位为1
                ret[i].push(1)
            } else {
                ret[i].push(ret[i-1][j]+ret[i-1][j-1])
            }
        }   
    }
    return ret
}

const numRows = 6;
const ret = triangle(numRows);
console.log(ret);
