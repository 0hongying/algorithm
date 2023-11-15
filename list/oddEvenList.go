package main

//给定单链表的头节点 head ，将所有索引为奇数的节点和索引为偶数的节点分别组合在一起，然后返回重新排序的列表

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	odd := head
	even := head.Next
	evenHead := even

	for even != nil && even.Next != nil {
		// 奇数节点的next为偶数节点的next
		odd.Next = even.Next
		odd = odd.Next
		// 偶数节点的next为奇数节点的next
		even.Next = odd.Next
		even = even.Next
	}

	odd.Next = evenHead
	return head
}
