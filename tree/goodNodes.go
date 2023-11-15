package main

import "math"

func goodNodes(root *TreeNode) int {
	ans := 0
	var dfs func(node *TreeNode, v int)
	dfs = func(node *TreeNode, v int) {
		if node == nil {
			return
		}

		if node.Val >= v {
			v = node.Val
			ans++
		}

		dfs(node.Left, v)
		dfs(node.Right, v)
	}

	dfs(root, math.MinInt32)

	return ans
}
