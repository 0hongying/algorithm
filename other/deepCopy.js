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

const source = { name: 'lhy', age: 26, game:[ 'poekr', 'sparrow soul' ], target: { name: 'tiktok' }  };
const ret = deepCopy(source);
console.log(ret);