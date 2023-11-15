// 请你找出所有和为0且不重复的三元组
var threeSum = function(nums) {
    nums.sort((a, b) => a - b)
    const ans = []
    let first = 0
    for (;first < nums.length;first++) {
        if (first > 0 && nums[first] == nums[first-1]) continue

        let second = first+1
        let right = nums.length - 1
        while (second < right) {
            const sum = nums[first] + nums[second] + nums[right]
            if (sum == 0) {
                ans.push([nums[first], nums[second], nums[right]])
                second++
                right--
            } else if (sum > 0) {
                right--
            } else {
                second++
            }
        }

    }
    return ans
};

/**
 * 思路 
 * 先将数组排序
 * 然后固定一个非正数（如果起始为正数，则不可能三个数加起来为0），使用双指针left和right指向两端
 * 如果加起来之和大于0，则right--，如果小于0，则left++
 */

const nums = [ 1, -2, -3, 4, 5, 1, 4, 9, -5, -6 ];
const ret = threeSum(nums);
console.log(ret);