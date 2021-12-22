// 给定一个非负整数数组 nums ，你最初位于数组的第一个下标 
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。判断你是否能够到达最后一个下标
function canJump1(nums) {
    let n = nums.length
    // 最远位置
    let rightmost = 0
    for (let i = 0; i < n; ++i) {
        // 判断当前最远步骤是否超过此刻下标
        if (i <= rightmost) {
            rightmost = Math.max(rightmost, i + nums[i]);
            if (rightmost >= n - 1) {
                return true
            }
        }
    }
    return false
}


// 给定一个非负整数数组 nums ，你最初位于数组的第一个下标 
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。（使用最少步骤）

function canJump2(nums) {
    let length = nums.length
    let end = 0
    let maxPosition = 0 
    let steps = 0
    for (let i = 0; i < length - 1; i++) {
        maxPosition = Math.max(maxPosition, i + nums[i]) 
        if (i === end) {
            end = maxPosition
            steps++
        }
    }
    return steps
}

/**
 * 思路
 * 遍历数组中的每一个位置，实时维护最远可以到达的位置（x+nums[x]）
 */