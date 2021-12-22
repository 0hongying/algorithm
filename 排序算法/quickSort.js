// O(nlogn)
// 不稳定
function quickSort(arr, left, right) {
    if (left > right) {
        return
    }
    let i = left, j = right, k = arr[left]
    while (i < j) {
        while (i < j && arr[j] >= k) {
            j--
        }
        while (i < j && arr[i] <= k) {
            i++
        }
        if (i < j) {
            let temp = arr[i]
            arr[i] = arr[j]
            arr[j] = temp
        }
    }
    arr[left] = arr[i]
    arr[i] = k
    quickSort(arr, left, i-1)
    quickSort(arr, i+1, right)
}