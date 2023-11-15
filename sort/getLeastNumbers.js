// 获取前k个数

/**
 * 快速排序求前k个数
 */
var getLeastNumbers = function(arr, k) {
    if (k === 0 || arr.length === 0) return []
    return quickSearch(arr, 0, arr.length-1, k - 1)
}

function quickSearch(arr, left, right, k) {
    const index = partition(arr, left, right)
    if (index === k) return arr.slice(arr.length - k)
    return index > k? quickSearch(arr, left, index - 1, k): quickSearch(arr, index + 1, right, k)
}

function partition(arr, left, right) {
    let i = left, j = right
    let k = arr[left]
    while (i < j) {
        while(i < j && arr[j] >= k) j--
        while(i < j && arr[i] <= k) i++
        if(i < j) {
            const temp = arr[j]
            arr[j] = arr[i]
            arr[i] = temp
        }
    }
    arr[left] = arr[i]
    arr[i] = k
    return i
}

/**
 * k次冒泡排序
 */
 function kBubbleSort(arr, k) {
    for (let i = 0; i < k; i++) {
        let swap = false
        for (let j = 0; j < arr.length-1-i; j++) {
            if (arr[j] > arr[j+1]) {
                swap = true
                let temp = arr[j]
                arr[j] = arr[j+1]
                arr[j+1] = temp
            }
        }
        if (!swap) {
            return arr.slice(arr.length - k);
        }
    }
    return arr.slice(arr.length - k);
}


/**
 * 大顶堆
 */


/**
 * 随机选择法
 */