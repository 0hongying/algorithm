// 链表中倒数第k个节点
function getKthFromEnd(head, k) {
    let fast = head, slow = head
    // fast比slow快k步
    while(k > 0) {
        fast = fast.next
        k --
    }
    while(fast !== null) {
        fast = fast.next
        slow = slow.next
    }
    return slow
}

/**
 * 思路
 * 快慢指针，快指针比慢指针起始位置多k步，快指针到达尾部时，慢指针就是倒数第k个
 */