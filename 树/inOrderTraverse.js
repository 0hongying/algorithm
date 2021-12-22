// 左 --> 根 --> 右
function inOrderTraverse(root) {
    const ret = []
    var traverse = (root => {
        if (root !== null) {
            traverse(root.left)
            ret.push(root.val)
            traverse(root.right)
        }
    })
    traverse(root)
    return ret
}