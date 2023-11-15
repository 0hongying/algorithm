package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 快慢指针
	fast := head
	ret := &ListNode{Next: head}
	slow := head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	// 删除第 n 个结点
	slow.Next = slow.Next.Next
	return ret.Next
}
