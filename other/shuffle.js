// 洗牌算法
function shuffle(nums) {
    for (let i = nums.length-1; i >= 0; i--) {
        // [0, i]与前面的数字包括自己在内交换位置
        let index = parseInt(Math.random()*(i+1))
        let temp = nums[i]
        nums[i] = nums[index]
        nums[index] = temp
    }
}

const nums = [ 1, 3, 9, 7, 8, 4, 2, 6, 5 ];
shuffle(nums);
console.log(nums);

