package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct{}

func Constructor() Codec {
	return *new(Codec)
}

// 序列化和反序列化二叉树
func (Codec) serialize(root *TreeNode) string {
	sb := &strings.Builder{}
	var dfs func(node *TreeNode)
	// 根 -> 左 -> 右
	dfs = func(node *TreeNode) {
		if node == nil {
			sb.WriteString("null,")
			return
		}
		val := strconv.Itoa(node.Val)
		sb.WriteString(val)
		sb.WriteByte(',')
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return sb.String()
}

func (Codec) deserialize(data string) *TreeNode {
	dataArr := strings.Split(data, ",")
	var build func() *TreeNode
	build = func() *TreeNode {
		if dataArr[0] == "null" {
			dataArr = dataArr[1:]
			return nil
		}
		// 根 -> 左 -> 右
		val, _ := strconv.Atoi(dataArr[0])
		dataArr = dataArr[1:]
		return &TreeNode{val, build(), build()}
	}
	root := build()
	return root
}
