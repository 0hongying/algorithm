// 链表反转
function reverseList(head) {
    let prev = null
    let curr = head
    while (curr) {
        const next = curr.next
        curr.next = prev
        prev = curr
        curr = next
    }
    return prev
}

// next = cur -> next
// cur -> next = prev
// prev = cur
// cur = next
