// 洗牌算法
function shuffle(nums) {
    for (let i = nums.length-1; i >= 0; i--) {
        let index = parseInt(Math.random()*(i+1))
        let temp = nums[i]
        nums[i] = nums[index]
        nums[index] = temp
    }
}