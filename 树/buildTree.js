// 根据前序和中序生成二叉树
let map = new Map()
function buildTree(preorder, inorder) {
    let n = preorder.length
    for (let i = 0; i < n; i++) {
        map.put(inorder[i], i)
    }
    return myBuildTree(preorder, inorder, 0, n-1, 0, n-1)
}

function myBuildTree(preorder, inorder, preorderLeft, preorderRight, inorderLeft, inorderRight) {
    if (preorderLeft > preorderRight) return null
    // 前序第一个节点就是根节点
    let preorderRoot = preorderLeft
    // 找到中序中的根节点
    let inorderRoot = map.get(preorder[preorderRoot])
    // 生成根节点
    const root = new TreeNode(preorder[preorderRoot])
    // 以根节点分割中序遍历
    let leftSubTree = inorderRoot - inorderLeft
    // 前序[左边界+1, 左边界+左孩子个数] === 中序[左边界, 根-1]
    root.left = myBuildTree(preorder, inorder, preorderLeft+1, preorderLeft+leftSubTree, inorderLeft, inorderRoot-1)
    // 前序[左边界+左孩子个数+1, 右边界] === 中序[根+1, 右边界] 
    root.right = myBuildTree(preorder, inorder, preorderLeft+leftSubTree+1, preorderRight, inorderRoot+1, inorderRight)
    return root
}

/**
 * 思路
 * 前序第一个元素即为根节点，根据根节点将中序分割为左孩子和右孩子
 * 前序[左边界+1, 左边界+左孩子个数] === 中序[左边界, 根-1]
 * 前序[左边界+左孩子个数+1, 右边界] === 中序[根+1, 右边界] 
 */

