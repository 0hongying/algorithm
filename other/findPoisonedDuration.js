function findPoisonedDuration(timeSeries, duration) {
    let last = -1, ret = 0
    for (let i = 0; i < timeSeries.length; i++) {
        let end = timeSeries[i] + duration - 1
        // 判断是否有交集
        ans += last < s ? duration : end - last
        last = end
    }
    return ret
}