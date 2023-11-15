package main

func candy(ratings []int) int {
	candy := make([]int, len(ratings))
	for i, _ := range ratings {
		candy[i] = 1
	}

	// 先满足右规则
	for i, _ := range ratings {
		if i < len(ratings)-1 && ratings[i] < ratings[i+1] && candy[i] >= candy[i+1] {
			candy[i+1] = candy[i] + 1
		}
	}

	// 先满足左规则
	for i := len(ratings) - 1; i >= 0; i-- {
		if i > 0 && ratings[i] < ratings[i-1] && candy[i] >= candy[i-1] {
			candy[i-1] = candy[i] + 1
		}
	}

	ret := 0
	for _, v := range candy {
		ret += v
	}
	return ret
}

// func main() {
// 	ratings := []int{1, 6, 10, 8, 7, 3, 2}
// 	candy(ratings)
// }
