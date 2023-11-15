package main

// func avoidFlood(rains []int) []int {
// 	m := map[int]int{}
// 	ans := make([]int, len(rains))
// 	zoreIndex := []int{}

// 	for i, v := range zoreIndex {
// 		if v == 0 {
// 			zoreIndex = append(zoreIndex, i)
// 			ans[i] = 1
// 		} else {
// 			ans[i] = -1
// 		}
// 	}

// 	for i, v := range rains {
// 		m[v]++
// 		ans[i] = -1
// 		if v > 0 && m[v] > 1 {
// 			if len(zoreIndex) > 0 {
// 				m[0]--
// 				m[v]--
// 				ans[i] = -1
// 				index := zoreIndex[len(zoreIndex)-1]
// 				ans[index] = v
// 				zoreIndex = zoreIndex[:len(zoreIndex)-1]
// 			} else {
// 				return []int{}
// 			}
// 		}
// 		if v == 0 {
// 			zoreIndex = append(zoreIndex, i)
// 		}
// 	}
// 	return ans
// }

// func main() {
// 	rains := []int{1, 2, 0, 0, 2, 1}
// 	avoidFlood(rains)
// }
