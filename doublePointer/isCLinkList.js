// 判断是否成环
function isCycle(head) {
    let fast = head
    let slow = head
    while (fast !== null && fast.next !== null) {
        slow = slow.next
        fast = fast.next.next
        if (slow === fast) {
            return true
        }
    }
    return false
}

/**
 * 思路
 * 快慢指针，快指针走两步，慢指针走一步
 * 如果成环，那么快指针一定会追上慢指针
 * 如果快指针走3步，4步,会导致快指针多走几圈才相遇
 */