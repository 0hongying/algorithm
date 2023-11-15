package main

func numPairsDivisibleBy60(time []int) int {
	m := map[int]int{}
	ans := 0

	for _, v := range time {
		v %= 60
		ans += m[(60-v)%60]
		m[v]++

	}

	return ans
}

func main() {
	time := []int{60, 60, 60}
	numPairsDivisibleBy60(time)
}
