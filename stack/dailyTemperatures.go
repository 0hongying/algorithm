package main

// 给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替

func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	stack := []int{}

	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			// 保留没有找到
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[index] = i - index
		}
		stack = append(stack, i)
	}

	return ans
}
