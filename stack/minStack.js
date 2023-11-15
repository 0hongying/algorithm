// 时间O(1) + 空间O(n)
var MinStack = function() {
    this.stack = []
    this.min_stack = [Infinity]
}

MinStack.prototype.push = function(x){
    // 最小栈，保证栈顶为最小
    this.stack.push(x)
    this.min_stack.push(Math.min(this.min_stack[this.min_stack.length-1],x))
}

MinStack.prototype.pop = function() {
    // 一起出栈
    this.stack.pop()
    this.min_stack.pop()
}

MinStack.prototype.top = function() {
    return this.stack[this.stack.length-1]
}

MinStack.prototype.getMin = function() {
    return this.min_stack[this.min_stack.length-1]
}

// 时间O(1) + 空间O(1)

var MinStack = function() {
    this.stack = []
    this.min
}

MinStack.prototype.push = function(x) {
    if (this.stack.length === 0) {
        this.min = x
        this.stack.push(0)
    } else {
        let compare = x - this.min
        this.stack.push(compare)
        this.min = compare < 0 ? x : this.min
    }
}

MinStack.prototype.pop = function() {
    let top = this.stack[this.stack.length-1]
    // 栈顶元素若为负数，需要更新min
    this.min = top < 0 ? this.min - top : this.min
    this.stack.pop()
}

MinStack.prototype.top = function() {
    let top = this.stack[this.stack.length-1]
    // 栈顶元素若为负数，说明min就是x
    return top < 0 ? this.min : this.min + top
}

MinStack.prototype.getMin = function() {
    return this.min
}

/**
 * 第一种思路
 * 使用一个最小值栈，记录整个栈的最小值
 * 入栈：比较当前元素和栈顶元素哪个小
 * 出栈：两个栈同时执行出栈
 * 获取栈顶：返回栈顶元素
 * 获取最小值：返回最小值栈顶元素
 * 
 * 第二种思路
 * 使用一个min记录最小值，栈记录差值
 * 入栈：当栈为空时，x直接入栈，min为x；当栈不为空，x-min的值入栈，min为x或不变
 * 出栈：栈执行出栈，如果top为负数，则重新修改min的值等于min-top
 * 获取栈顶：如果top为负数，则等于min，否则等于min+top
 * 获取最小值：返回min
 */


const minStack = new MinStack();
minStack.push(1);
minStack.push(2);
minStack.push(3);
minStack.push(0);

// minStack.pop();
const top = minStack.top();
console.log(top);
const min = minStack.getMin();
console.log(min);

