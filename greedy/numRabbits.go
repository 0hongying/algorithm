package main

import (
	"sort"
)

func numRabbits(answers []int) int {
	sort.Ints(answers)
	ret := 0
	for i := 0; i < len(answers); i++ {
		ret += answers[i] + 1
		j := answers[i]
		// 跳过 answer 个相同值
		for j > 0 && i+1 < len(answers) && answers[i] == answers[i+1] {
			i++
			j--
		}
	}
	return ret
}

// func main() {
// 	answers := []int{0, 0, 1, 1, 1}
// 	ret := numRabbits(answers)
// 	fmt.Printf("%v", ret)
// }
