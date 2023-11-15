// 合并两个链表
function merge2Lists(l1, l2) {
    let head = new ListNode()
    let tail = head
    while (l1 !== null && l2 !== null) {
        if (l1.val < l2.val) {
            tail.next = l1
            l1 = l1.next
        } else {
            tail.next = l2
            l2 = l2.next
        }
        tail = tail.next
    }
    tail.next = l1 === null ? l2 : l1
    return head.next
}