// 判断是否对称二叉树
function isSymmetric(root) {
    return check(root, root)
}

function check(p, q) {
    if (p == null && q == null) {
        return true
    }
    if (p == null || q == null) {
        return false
    }
    // 左孩子的左孩子 === 右孩子的右孩子  && 左孩子的右孩子 === 右孩子的左孩子
    return p.val == q.val && check(p.left, q.right) && check(p.right, q.left)
} 