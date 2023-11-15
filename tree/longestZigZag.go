package main

// func longestZigZag(root *TreeNode) int {
// 	var dfs func(root *TreeNode, flag, prev int) (ans int)
// 	dfs = func(root *TreeNode, flag, prev int) (ans int) {
// 		if root == nil {
// 			return
// 		}

// 		if flag != prev {
// 			ans++
// 		}

// 		return max(dfs(root.Left, flag*-1), dfs(root.Right, flag*-1))
// 	}
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
