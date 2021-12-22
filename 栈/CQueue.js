// 使用两个栈实现队列
var CQueue = function() {
    // 入栈
    this.pushStack = []
    // 出栈
    this.popStack = []
}

CQueue.prototype.appendTail = function(value) {
    this.pushStack.push(value)
}

CQueue.prototype.deleteHead = function() {
    if (this.popStack.length === 0 && this.pushStack.length === 0) return -1
    else if (this.popStack.length !== 0) return this.popStack.pop()
    else {
        while(this.pushStack.length !== 0) this.popStack.push(this.pushStack.pop())
        return this.popStack.pop()
    }
}

/**
 * 思路
 * 一个栈（pushStack）用作入栈，另一个栈（popStack）用作出栈
 * 当popStack为空时，将pushStack所有元素出栈再入栈popStack
 */