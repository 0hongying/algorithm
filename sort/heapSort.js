// 大顶堆：arr[i] >= arr[2i+1] && arr[i] >= arr[2i+2] 即父节点大于等于左右子节点
// 小顶堆：arr[i] <= arr[2i+1] && arr[i] <= arr[2i+2] 即父节点小于等于左右子节点
// 步骤： 
// 最后一个非叶子结点（n/2-1)
// 


function heapSort(nums) {
    let len = nums.length - 1
    // 构建最大堆
    buildMaxHeap(nums, len)
    for (let i = len; i >= 1; --i) {
        // 堆顶元素和末尾元素交换
        swap(nums, i, 0)
        len -= 1
        // 调整堆
        maxHeapify(nums, 0, len)
    }
}

function swap(nums, i, j) {
    let temp = nums[i]
    nums[i] = nums[j]
    nums[j] = temp
}

function buildMaxHeap(nums, len) {
    for (let i = parseInt(len / 2) - 1; i >= 0; --i) {
        maxHeapify(nums, i, len)
    }
}

function maxHeapify(nums, i, len) {
    for (; (i << 1) + 1 <= len;) {
        // 左孩子
        let lson = (i << 1) + 1
        // 右孩子
        let rson = (i << 1) + 2
        let large
        if (lson <= len && nums[lson] > nums[i]) {
            large = lson
        } else {
            large = i
        }
        if (rson <= len && nums[rson] > nums[large]) {
            large = rson
        }
        if (large != i) {
            swap(nums, i, large)
            i = large
        } else {
            break
        }
    }
}
