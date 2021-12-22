// 请你找出所有和为0且不重复的三元组
function threeSum(nums) {
    if (nums.length < 3) return []
    const arr = nums.sort((a,b) => a-b)
    if (arr[0] > 0 || arr[arr.length-1] < 0) return []
    const ret = []
    // 先固定一个数
    for (let i = 0; i < arr.length; i++) {
        // 正数不需要考虑
        if (arr[i] > 0) return ret
        if (i > 0 && arr[i] === arr[i-1]) continue
        let left = i + 1
        let right = arr.length-1
        while (left < right) {
            const temp = arr[i] + arr[left] + arr[right]
            if (temp === 0) {
                ret.push([nums[i], nums[left], nums[right]])
                while (left < right && nums[left] === nums[left+1]) {
                    left++
                }
                while (left < right && nums[right] === nums[right-1]) {
                    right--
                }
                left++
                right--
            } else if (temp < 0) {
                left++
            } else {
                right--
            }
        }
    }
    return ret
}

/**
 * 思路 
 * 先将数组排序
 * 然后固定一个非正数（如果起始为正数，则不可能三个数加起来为0），使用双指针left和right指向两端
 * 如果加起来之和大于0，则right--，如果小于0，则left++
 */