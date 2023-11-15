package main

//在一个大小为 n 且 n 为 偶数 的链表中，对于 0 <= i <= (n / 2) - 1 的 i ，第 i 个节点（下标从 0 开始）的孪生节点为第 (n-1-i) 个节点 。
// 比方说，n = 4 那么节点 0 是节点 3 的孪生节点，节点 1 是节点 2 的孪生节点。这是长度为 n = 4 的链表中所有的孪生节点。
// 孪生和 定义为一个节点和它孪生节点两者值之和。
// 给你一个长度为偶数的链表的头节点 head ，请你返回链表的 最大孪生和

func pairSum(head *ListNode) int {
	val := []int{}
	for node := head; node != nil; node = node.Next {
		val = append(val, node.Val)
	}
	ret := 0
	for i, _ := range val {
		ret = max(ret, val[i]+val[len(val)-1-i])
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
