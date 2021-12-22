function binarySearch(arr, target) {
    let left = 0
    let right = arr.length
    while (left < right) {
        let mid = parseInt((left+right)/2)
        if (arr[mid] === target) {
            return mid
        } else if (arr[mid] < target) {
            left = mid + 1 
        } else {
            right = mid
        }
    }
    return -1
}

// 左边界
function binarySearch1(arr, target) {
    let left = 0
    let right = arr.length
    //左开右闭[left, right)
    while (letf < right) {
        let mid = parseInt((left+right)/2)
        if (arr[mid] === target) {
            right = mid
        } else if (arr[mid] < target) {
            left = mid + 1
        } else {
            // 不是mid-1
            right = mid
        }
    }
    return arr[left] === target ? left : -1
}
// 右边界
function binarySearch2(arr, target) {
    let left = 0
    let right = arr.length
    // 左开右闭[left, right)
    while (left < right) {
        let mid = parseInt((left+right)/2)
        if (arr[mid] === target) {
            left = mid + 1
        } else if (arr[mid] < target) {
            left = mid + 1
        } else {
            // 不是mid-1
            right = mid
        }
    }
    // 这里是right-1
    return arr[right-1] === target ? right-1 : -1
}


/**
 * 思路
 * 三种变形（1.二分法 2.左边界二分法 3.右边界二分法）
 */