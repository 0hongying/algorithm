package main

// 二叉树的 直径 是指树中任意两个节点之间最长路径的 长度 。这条路径可能经过也可能不经过根节点 root 。
// 两节点之间路径的 长度 由它们之间边数表示

func diameterOfBinaryTree(root *TreeNode) (ans int) {
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return -1
		}

		left := dfs(root.Left) + 1
		right := dfs(root.Right) + 1

		if ans < left+right {
			ans = left + right
		}
		if left < right {
			return right
		}
		return left
	}
	dfs(root)
	return
}
