package main

func checkDistances(s string, distance []int) bool {
	m := map[byte]int{}
	for i := 0; i < len(s); i++ {
		v := s[i]
		if m[v] == 0 {
			m[v] = i + 1
		} else {
			m[v] = i - m[v]
		}
	}

	for i, _ := range m {
		if distance[i-97] != m[i] {
			return false
		}
	}
	return true
}

// func main() {
// 	s := "aa"
// 	distance := []int{1, 3, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
// 	ret := checkDistances(s, distance)
// 	fmt.Printf("%v", ret)
// }
