package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	val := dfs(root)
	return max(val[0], val[1])
}

func dfs(node *TreeNode) []int {
	if node == nil {
		return []int{0, 0}
	}
	l, r := dfs(node.Left), dfs(node.Right)
	// 选根节点（只能在左右孩子中未选中的）
	selected := node.Val + l[1] + r[1]
	// 不选根节点（可以从左右孩子中最大的）
	notSelected := max(l[0], l[1]) + max(r[0], r[1])
	return []int{selected, notSelected}
}
