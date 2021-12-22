function mergeKLists(lists) {
    if (!lists.length) return null
    if (lists.length === 1) return lists[0]
    // 两两合并
    lists.splice(0, 2, merge2Lists(lists[0], lists[1]))
    return mergeKLists(lists)
}

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