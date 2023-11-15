package main

// https://leetcode.cn/problems/binary-tree-right-side-view/description/?envType=study-plan-v2&envId=top-100-liked
func rightSideView(root *TreeNode) []int {
	ans := []int{}
	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if level == len(ans)+1 {
			ans = append(ans, node.Val)
		}

		dfs(node.Right, level+1)
		dfs(node.Left, level+1)

	}
	dfs(root, 1)
	return ans
}
