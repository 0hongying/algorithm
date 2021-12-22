// 层次遍历
function levelOrderTraverse(root) {
    const ret = []
    if (!root) return ret
    const q = []
    // 先将根节点压入队列
    q.push(root)
    while (q.length !== 0) {
        const size = q.length
        ret.push([])
        // 循环队列
        for (let i = 1; i <= size; i++) {
            const node = q.shift()
            ret[ret.length-1].push(node.val)
            if (node.left) q.push(node.left)
            if (node.right) q.push(node.right)
        }
    }
    return ret
}