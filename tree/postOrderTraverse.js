// 后序
// 左 --> 右 --> 根
function postOrderTraverse(root) {
    const ret = []
    var traverse = (root => {
        if (root !== null) {
            traverse(root.left)
            traverse(root.right)
            ret.push(root.val)
        }
    })
    traverse(root)
    return ret
}

// 非递归
function preOrderTraverse2(root) {
    const stack = []
    const ret = []
    if (root) stack.push(root)
    while(stack.length) {
        const node = stack.pop()
        ret.push(node.val)
        // 先右再左
        if (node.right) stack.push(node.right)
        if (node.left) stack.push(node.left)
    }
    return ret.reverse(); // 左->右->根
}