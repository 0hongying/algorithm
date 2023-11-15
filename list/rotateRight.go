package main

// 给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}

	// 先统计链表长度
	count := 1
	node := head
	for node.Next != nil {
		node = node.Next
		count++
	}
	// 新增节点数
	k = count - k%count

	// 链接原头节点
	node.Next = head
	for k > 0 {
		node = node.Next
		k--
	}

	newHead := node.Next
	node.Next = nil

	return newHead
}
