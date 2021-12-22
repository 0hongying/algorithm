// 后缀表达式
const operator = []
const priority = {
    '+': 1,
    '-': 1,
    '*': 2,
    '/': 2
}

function getTop() {
    return operator[operator.length-1]
}

function reversePoland(formula) {
    const ret = []
    for (const character of formula) {
        if (Number(character)) {
            ret.push(character)
        } else {
            if (!getTop()) {
                // 栈为空
                operator.push(character)
            } else if (character === ')') {
                // 运算符为右括号
                while (1) {
                    if (getTop() === '(') {
                        operator.pop()
                        break
                    } else {
                        ret.push(operator.pop())
                    }
                }
            } else if (character === '(' || getTop() === '(' || priority[character] > priority[getTop()] ) {
                // 1.运算符为左括号
                // 2.栈顶元素为左括号
                // 3.运算符优先级大于栈顶元素优先级
                operator.push(character)
            } else {
                ret.push(operator.pop())
                operator.push(character)
            }
        }
    }
    while(getTop()) {
        ret.push(operator.pop())
    }
    return ret
}

function calculateReversePoland() {
    const ret = []
    for (const character of expression) {
        if (Number(character)) {
            ret.push(character)
        } else {
            const num1 = ret.pop()
            const num2 = ret.pop()
            ret.push(calculateReversePoland(num1, num2, character))
        }
    }
    return ret.pop()
}

function calculate(num1, num2, op) {
    switch (op) {
        case '+': return num1 + num2
        case '-': return num1 - num2
        case '*': return num1 * num2
        case '/': return num1 / num2
        default:
            throw new Error('类型不支持')
    }
}