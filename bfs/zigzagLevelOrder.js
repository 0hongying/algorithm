function TreeNode(val, left, right) {
    this.val = (val===undefined ? 0 : val)
    this.left = (left===undefined ? null : left)
    this.right = (right===undefined ? null : right)
}

// 给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行
var zigzagLevelOrder = function(root) {
    const ret = []
    let index = 0
    if (!root) return ret
    const q = []
    q.push(root)
    while (q.length !== 0) {
        const size = q.length
        ret.push([])
        for (let i = 1; i <= size; i++) {
            const node = q.shift()
            ret[ret.length-1].push(node.val)
            if (node.left) q.push(node.left)
            if (node.right) q.push(node.right)
        }
        if (index & 1) ret[ret.length-1].reverse()
        index ++
    }
    return ret
};