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