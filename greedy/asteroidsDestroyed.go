package main

import (
	"sort"
)

func asteroidsDestroyed(mass int, asteroids []int) bool {
	sort.Ints(asteroids)
	for _, v := range asteroids {
		if mass < v {
			return false
		}
		mass += v
	}
	return true
}

// func main() {
// 	mass := 5
// 	asteroids := []int{4, 9, 23, 4}
// 	ret := asteroidsDestroyed(mass, asteroids)
// 	fmt.Printf("%v", ret)
// }
