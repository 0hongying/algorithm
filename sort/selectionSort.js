// O(n^2)
// 不稳定

// https://oi-wiki.org/basic/images/selection-sort-1-animate-example.svg
function selectionSort(arr) {
    for (let i = 0; i < arr.length; i++) {
        let index = i
        for (let j = i+1; j < arr.length; j++) {
            if (arr[j] < arr[index]) {
                index = j
            }
        }
        let temp = arr[i]
        arr[i] = arr[index]
        arr[index] = temp
    }
}
