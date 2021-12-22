// 判断它是否是高度平衡的二叉树

// 自底向上
function isBalanced1(root) {
    return height1(root) >= 0
};

function height1(root) {
    if (root == null) return 0
    // 
    let leftHeight = height1(root.left)
    let rightHeight = height1(root.right)
    if (leftHeight == -1 || rightHeight == -1 || Math.abs(leftHeight - rightHeight) > 1) {
        return -1
    } else {
        return Math.max(leftHeight, rightHeight) + 1
    }
}

//自顶向下
function isBalanced2(root) {
    if (root === null) return true
    else return Math.abs(height2(root.left)-height2(root.right)) <= 1 && isBalanced2(root.left) && isBalanced2(root.right)
}

function height2(root) {
    if (root === null) return 0
    else Math.max(height2(root.left), height2(root.left))+1
}