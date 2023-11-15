package main

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 给你一棵二叉树的根节点，返回所有重复的子树
// root = [1,2,3,4,null,2,4,null,null,4]
// [[2,4],[4]]
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	// 创建一个空map
	m := map[string]int{}
	// 创建一个长度为0，容量为0的切片
	ret := []*TreeNode{}
	var dfs func(node *TreeNode) string

	dfs = func(node *TreeNode) string {
		if node == nil {
			return "|"
		}
		key := strconv.Itoa(node.Val) + "_"
		key += dfs(node.Left)
		key += dfs(node.Right)
		if v, ok := m[key]; ok {
			m[key] = v + 1
			if m[key] == 2 {
				ret = append(ret, node)
			}
		} else {
			m[key] = 1
		}
		return key
	}
	dfs(root)
	return ret
}
