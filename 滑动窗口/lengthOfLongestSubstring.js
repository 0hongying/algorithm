// 求一个字符串s中不含有重复字符的最长子串的长度
function lengthOfLongestSubstring(s) {
    if (s.length === 0) return 0
    const map = new Map()
    let max = 0, left = 0
    // i为右边界
    for (let i = 0; i < s.length; i++) {
        // 修改子串的起始位置
        if (map.has(s.charAt(i))) left = Math.max(left, map.get(s.charAt(i))+1)
        // 更新下标
        map.set(s.charAt(i), i)
        max = Math.max(max, i-left+1)
    }
    return max
}

/**
 * 思路
 * 使用map记录字符的下标
 */

