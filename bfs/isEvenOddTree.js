// 二叉树根节点所在层下标为 0 ，根的子节点所在层下标为 1 ，根的孙节点所在层下标为 2 ，依此类推
// 偶数下标 层上的所有节点的值都是 奇 整数，从左到右按顺序 严格递增
// 奇数下标 层上的所有节点的值都是 偶 整数，从左到右按顺序 严格递减

var isEvenOddTree = function(root) {
    const queue = []
    if (root == null) {
         return true
    }
    let index = 0
    queue.push(root)
    while (queue.length != 0) {
        ret = []
        const size = queue.length
        for (let i = 0; i < size; i++) {
            const node = queue.shift()
            ret.push(node.val)
            if (node.left) queue.push(node.left)
            if (node.right) queue.push(node.right)
        }
        if (index % 2 == 0) {
            for (let i = 0; i < ret.length; i++) {
                if (ret[i] % 2 == 0) return false
                if (ret[i] >= ret[i+1]) return false
            }
        } else {
            for (let i = 0; i < ret.length; i++) {
                if (ret[i] % 2 == 1) return false
                if (ret[i] <= ret[i+1]) return false
            }
        }
        index++
    }
    return true
};