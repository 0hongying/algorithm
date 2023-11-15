package main

// 给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目

func pathSum(root *TreeNode, targetSum int) int {
	ans := 0
	var dfs func(root *TreeNode, targetSum int, path []int)
	dfs = func(root *TreeNode, targetSum int, path []int) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		ans += getCnt(path, targetSum)
		dfs(root.Left, targetSum, path)
		dfs(root.Right, targetSum, path)
	}
	dfs(root, targetSum, []int{})
	return ans
}

func getCnt(path []int, targetSum int) int {
	cnt := 0
	sum := 0
	for i := len(path) - 1; i >= 0; i-- {
		sum += path[i]
		if sum == targetSum {
			cnt++
		}
	}
	return cnt
}
