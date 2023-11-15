// 千分位分隔符
function toLocaleString(num) {
     // 获取整数部分
    return num && num.toString().replace(/\d+/, function(str) {
        // x(?=y): 仅匹配x后面位y
        // $1: 将第一个捕获组后面加逗号
        return str.replace(/(\d)(?=(\d{3})+$)/g, '$1,')
    })
}

const num = 11111;
const ret = toLocaleString(num);
console.log(ret);