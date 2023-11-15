package main

// 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）

func swapPairs(head *ListNode) *ListNode {
	h := &ListNode{Next: head}
	node0 := h
	node1 := head
	for node1 != nil && node1.Next != nil {
		node2 := node1.Next
		node3 := node2.Next

		node0.Next = node2
		node2.Next = node1
		node1.Next = node3

		node0 = node1
		node1 = node3
	}
	return h.Next
}
