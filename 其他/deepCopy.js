// 深拷贝
function deepCopy(source) {
    let traget = Array.isArray(source) ? [] : {}
    for (const k in source) {
        if (source.hasOwnProperty(k)) {
            if (typeof source[k] === 'object') {
                traget[k] = deepCopy(source[k])
            } else {
                traget[k] = source[k]
            }
        }
    }
    return traget
}
/**
 * 思路
 * 当属性为object时，递归调用
 */