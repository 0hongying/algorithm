// 删除有序数组中的重复项
function removeDuplicates(nums) {
    let t = 0
    for (let i = 0; i < nums.length; i ++ ) {
        if (i == 0 || nums[i] != nums[i - 1]) {
            nums[t++] = nums[i]
        }
    }
    return t
}