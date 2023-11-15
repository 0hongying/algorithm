// 删除二叉搜索树的某个节点
function deleteNode(root, key) {
    if (root === null) {
        return null
    }
    if (root.val < key) {
        root.right = deleteNode(root.right, key)
    } else if (root.val > key) {
        root.left = deleteNode(root.left, key)
    } else {
        if (root.left === null && root.right === null) {
            // 删除节点为叶子节点，可以直接删除
            root = null
        } else if (root.right !== null) {
            // 删除节点，包含两种情况
            // 1.只有右子树
            // 2.既有左子树又有右子树
            root.val = suffix(root)
            root.right = deleteNode(root.right, root.val)
        } else {
            // 删除节点只有左子树
            root.val = prefix(root)
            root.left = deleteNode(root.left, root.val)
        }
    }
    return root
}

// 获取右孩子的最左节点
function suffix(root) {
    root = root.right
    while(root.left !== null){
        root = root.left
    }
    return root.val
}

// 获取左孩子的最右节点
function prefix(root) {
    root = root.left
    while(root.right !==null){
        root = root.right
    }
    return root.val
}

/**
 * 思路
 * 分为4种情况
 * 1.删除节点为叶子节点
 * 2.删除节点既有左节点又有右节点
 * 3.删除节点只有左节点
 * 4.删除节点只有右节点
 * 
 * 2和4为同一种情况
 */