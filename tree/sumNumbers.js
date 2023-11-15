// 给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
// 每条从根节点到叶节点的路径都代表一个数字
// 计算从根节点到叶节点生成的 所有数字之和

function sumNumbers(root) {
    var dfs = ((root, i) => {
        if (root === null) return 0
        const temp = i * 10 + root.val
        // 叶子节点直接计算value
        if (root.left === null && root.right === null) return temp
        return dfs(root.left, temp)+dfs(root.right, temp) 
    })
    return dfs(root, 0)
}

/**
 * 思路
 * num = num * 10 + value
 * 如果为叶子节点，则直接返回数值
 */

