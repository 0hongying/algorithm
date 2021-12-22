// 循环节
function loopFestival(str) {
    if (/cdots$/.test(str)) {
        return /^\d+\.\d*?(\d+?)\1+cdots$/.test(str) ? RegExp.$1 : undefined
    }
}