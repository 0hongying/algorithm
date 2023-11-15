function binarySearch(arr, target) {
    let left = 0
    let right = arr.length
    //左开右闭[left, right)
    while (left < right) {
        let mid = parseInt((left+right)/2)
        if (arr[mid] === target) {
            // 找到直接返回
            return mid
        } else if (arr[mid] < target) {
            left = mid + 1 
        } else {
            right = mid
        }
    }
    return -1
}

const arr = [ 1, 2, 3, 4, 5, 9, 9, 9, 9, 10, 11, 12, 13, 14, 15 ];
const ret = binarySearch(arr, 9);
console.log(ret);

// 左边界
function binarySearch1(arr, target) {
    let left = 0
    let right = arr.length
    //左开右闭[left, right)
    while (left < right) {
        let mid = parseInt((left+right)/2)
        if (arr[mid] === target) {
            // 找到后，继续缩小右范围
            right = mid
        } else if (arr[mid] < target) {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return arr[left] === target ? left : -1
}

const arr1 = [ 1, 2, 3, 4, 5, 9, 9, 9, 9, 10, 11, 12, 13, 14, 15 ];
const ret1 = binarySearch1(arr1, 9);
console.log(ret1);

// 右边界
function binarySearch2(arr, target) {
    let left = 0
    let right = arr.length
    // 左开右闭[left, right)
    while (left < right) {
        let mid = parseInt((left+right)/2)
        if (arr[mid] === target) {
            // 找到后，继续缩小左范围
            left = mid + 1
        } else if (arr[mid] < target) {
            left = mid + 1
        } else {
            right = mid
        }
    }
    // 这里是right-1
    return arr[right-1] === target ? right-1 : -1
}

const arr2 = [ 1, 2, 3, 4, 5, 9, 9, 9, 9, 10, 11, 12, 13, 14, 15 ];
const ret2 = binarySearch2(arr2, 9);
console.log(ret2);
