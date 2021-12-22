 function minArray(numbers) {
    let left = 0, right = numbers.length-1
    while (left < right) {
        let mid = parseInt((left+right)/2)
        // [6,7,3,4,5] 旋转点一定在左排序数组
        if (numbers[mid] < numbers[right]) right = mid
        // [3,4,5,1,2] 旋转点一定在右排序数组 
        else if (numbers[mid] > numbers[right]) left = mid +1
        // 旋转点不确定, 但肯定不是最后一位
        else right--
    }
    return numbers[left]
}

/**
 * 思路
 * 左边界一定大于右边界，根据中心点与右边界的大小关系，判断旋转点在哪个半区
 */