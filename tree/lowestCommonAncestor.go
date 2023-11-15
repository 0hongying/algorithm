package main

// 找出二叉树中某两个节点的第一个共同祖先
func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		// 找到p，q结点
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// x 的左孩子和右孩子包含 p , q
	if left != nil && right != nil {
		return root
	}
	// x恰好为 p 或 q，右子树包含了另一个节点
	if left != nil {
		return left
	}
	// x恰好为 p 或 q，左子树包含了另一个节点
	return right
}
