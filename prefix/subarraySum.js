var subarraySum = function(nums, k) {
    let ans = 0
    let prev = 0
    for (let i = 0; i < nums.length; i++) {
        prev += nums[i]
        if (prev == k) {
            ans++
        }

        var temp = prev
        for (let j = 0; j < i; j++) {
            temp -= nums[j]
            if (temp == k) {
                ans++
            }
        }
    }
    return ans
};

subarraySum([1,1,1], 2)