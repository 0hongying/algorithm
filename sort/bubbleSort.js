// O(n^2)
// 稳定

// 两两交换，将最大值换到尾部
function bubbleSort(arr) {
    for (let i = 0; i < arr.length-1; i++) {
        let swap = false
        for (let j = 0; j < arr.length-1-i; j++) {
            if (arr[j] > arr[j+1]) {
                swap = true
                let temp = arr[j]
                arr[j] = arr[j+1]
                arr[j+1] = temp
            }
        }
        if (!swap) return
    }
}


