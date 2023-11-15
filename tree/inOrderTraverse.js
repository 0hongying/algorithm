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

// 非递归
function inOrderTraverse2(root) {
    const ret = [];
    const stack = [];
    while (root || stack.length) {
        // 一直把左节点入栈
        while (root) {
            stack.push(root);
            root = root.left;
        }
        root = stack.pop();
        ret.push(root.val);
        root = root.right;
    }
    return ret;
};
