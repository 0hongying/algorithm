// 寻找节点p和节点q最近的父节点
function lowestCommonAncestor(root, p, q) {
    if(root === null) return null
    // root为p,q中的一个
    if(root === p || root === q) return root
    //分解为求左子树的子问题和右子树的子问题,注意是后序遍历
    let left_have = lowestCommonAncestor(root.left,p,q)    
    let right_have = lowestCommonAncestor(root.right,p,q)
    // p，q分别为root的左右子树
    if(left_have && right_have) return root
    else return left_have ? left_have : right_have
}