function findCenter(head) {
    let fast = head, slow = head
    while (fast !== null && fast.next !== null) {
        fast = fast.next.next
        slow = slow.next
    }
    return slow
}

/**
 * 思路
 * 快慢指针，快指针走两步，慢指针走一步
 */