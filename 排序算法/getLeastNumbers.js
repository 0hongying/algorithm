// 获取前k个数
var getLeastNumbers = function(arr, k) {
    if (k === 0 || arr.length === 0) return []
    return quickSearch(arr, 0, arr.length-1, k-1)
}


function quickSearch(arr, left, right, k) {
    const index = partition(arr, left, right)
    if (index === k) return arr.slice(0, k+1)
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
 * 思路
 * 变形快排求前k个数
 * 
 */