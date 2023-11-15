package main

import (
	"fmt"
	"sort"
)

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	m1 := map[byte]int{}
	m2 := map[byte]int{}
	for i, _ := range word1 {
		v1 := word1[i]
		v2 := word2[i]
		m1[v1]++
		m2[v2]++
	}

	arr1 := []int{}
	arr2 := []int{}

	for k, _ := range m1 {
		if _, ok := m2[k]; !ok {
			return false
		}
	}

	for _, v := range m1 {
		arr1 = append(arr1, v)
	}

	for _, v := range m2 {
		arr2 = append(arr2, v)
	}

	sort.Ints(arr1)
	sort.Ints(arr2)

	return fmt.Sprint(arr1) == fmt.Sprint(arr2)
}

// func main() {
// 	word1 := "abbzccca"
// 	word2 := "babzzczc"
// 	ret := closeStrings(word1, word2)
// 	fmt.Printf("%v", ret)
// }
