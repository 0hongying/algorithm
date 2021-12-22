// O(n^2)
// 稳定
function insertSort(arr) {
    for (let i = 1; i < arr.length; i++) {
        let temp = arr[i]
        let j 
        for (j = i -1; j >=0 && arr[j] > temp; j--) {
            arr[j+1] = arr[j]
        }
        arr[j+1] = temp
    }
}
