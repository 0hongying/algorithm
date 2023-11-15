function longestCommonPrefix(strs) {
    strs.sort()
    let ret = ''
    for (let i = 0; i < Math.min(strs[0].length, strs[strs.length-1].length); i++) {
        if (strs[0][i] === strs[strs.length-1][i]) {
            ret += strs[0][i]
        } else {
            break
        }
    }
    return ret
}


console.log(longestCommonPrefix(["daa","doga","do"]))