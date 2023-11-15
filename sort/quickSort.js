// O(nlogn)
// 不稳定

function quickSort(arr, left, right) {
    if (left > right) {
        return
    }
    let i = left, j = right, k = arr[left]
    while (i < j) {
        // 从右找第一个小于基准
        while (i < j && arr[j] >= k) {
            j--
        }
        // 从左找第一个大于基准
        while (i < j && arr[i] <= k) {
            i++
        }
        if (i < j) {
            let temp = arr[i]
            arr[i] = arr[j]
            arr[j] = temp
        }
    }
    // 交换基准
    arr[left] = arr[i]
    arr[i] = k
    quickSort(arr, left, i-1)
    quickSort(arr, i+1, right)
}
const arr = [3, 2, 3, 1, 2, 4, 5, 5, 6]
quickSort(arr, 0, 8)

console.log(arr)
