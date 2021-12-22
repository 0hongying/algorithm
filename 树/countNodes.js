// 计算二叉树的节点个数
function countNodes(root) {
    return root ? countNodes(root.left)+countNodes(root.right)+1 : 0
}