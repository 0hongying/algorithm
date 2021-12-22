// 盛最多水的容器
function maxArea(height) {
    let s = 0, e = height.length -1, max = Number.MIN_VALUE
    while (s < e) {
        // 获取较小的一侧
        let min = Math.min(height[s], height[e])
        max = Math.max(min * (e - s), max)
        if (height[s] < height[e]) s++
        else e-- 
    }
    return max
}

/**
 * 思路
 * 体积 = 两侧中的最小高度 * 下标差
 * 如果体积想要更大，则高度较小的一侧需要移动
 */