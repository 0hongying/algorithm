package main

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := []int{}

	if len(tokens) == 0 {
		return 0
	}

	for _, v := range tokens {
		if v == "+" || v == "/" || v == "*" || v == "-" {
			num2 := stack[len(stack)-1]
			num1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			res := cal(num1, num2, v)
			stack = append(stack, res)
		} else {
			num, _ := strconv.Atoi(v)
			stack = append(stack, num)
		}
	}
	return stack[len(stack)-1]
}

func cal(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	}
	return 0
}

// func main() {
// 	token := []string{
// 		"4", "13", "5", "/", "+",
// 	}
// 	evalRPN(token)
// }
