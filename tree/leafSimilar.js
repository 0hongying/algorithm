// https://leetcode.cn/problems/leaf-similar-trees/description/

var leafSimilar = function(root1, root2) {
    const ret = []
    var dfs = (root) => {
        if (root == null) {
            return null
        }
        if (root.left == null && root.right == null) {
            ret.push(root.val)
        }
        dfs(root.left)
        dfs(root.right)
    }
    dfs(root1)
    dfs(root2)
    const l = ret.length / 2
    for (let index = 0; index < l; index++) {
        if (ret[index] !== ret[l+index]) {
            return false
        }
    }
    return true
};