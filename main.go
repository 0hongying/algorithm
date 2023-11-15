package main

// var wg sync.WaitGroup

// func cat(fishChan, catChan chan bool) {
// 	defer wg.Done()
// 	defer close(catChan)

// 	for i := 0; i < 4; i++ {
// 		<-fishChan
// 		fmt.Println("cat")
// 		catChan <- true
// 	}
// }

// func dog(catChan, dogChane chan bool) {
// 	defer wg.Done()
// 	defer close(dogChane)

// 	for i := 0; i < 4; i++ {
// 		<-catChan
// 		fmt.Println("dog")
// 		dogChane <- true
// 	}
// }

// func fish(dogChane, fishChan chan bool) {
// 	defer wg.Done()
// 	defer close(fishChan)

// 	for i := 0; i < 4; i++ {
// 		<-dogChane
// 		fmt.Println("fish")
// 		fishChan <- true
// 	}
// }

// func fibonacci(n int, c chan int) {
// 	defer close(c)
// 	x, y := 0, 1
// 	for i := 0; i < n; i++ {
// 		c <- y
// 		x, y = y, x+y
// 	}
// }

// // 服务费 fee = (单数 - 对应的 速算扣除单数[fc]) * 对应的 服务费 + 速算扣除数[fd]

// func senven(n int) []int {
// 	ret := []int{}
// 	for i := 1; i <= n; i++ {
// 		if i%7 == 0 {
// 			ret = append(ret, i)
// 		} else {
// 			k := i
// 			for k > 0 {
// 				j := k % 10
// 				if j == 7 {
// 					ret = append(ret, i)
// 				}
// 				k = k / 10
// 			}
// 		}
// 	}
// 	return ret
// }

// func main() {
// catChan := make(chan bool, 1)
// dogChan := make(chan bool, 1)
// fishChan := make(chan bool, 1)
// fishChan <- true

// go cat(fishChan, catChan)
// go dog(catChan, dogChan)
// go fish(dogChan, fishChan)

// wg.Add(3)
// wg.Wait()
// }

// 	// c := make(chan int, 10)
// 	// go fibonacci(cap(c), c)
// 	// for i := range c {
// 	// 	fmt.Println(i)
// 	// }

// 	word := "abcc"
// 	ret := equalFrequency(word)
// 	fmt.Printf("%v", ret)
// }

// func equalFrequency(word string) bool {
// 	m := make(map[rune]int, 26)
// 	maxNum := 0
// 	minNum := math.MaxInt64
// 	for _, w := range word {
// 		m[w]++
// 	}

// 	count := 0
// 	for _, v := range m {
// 		maxNum = max(maxNum, v)
// 		minNum = min(minNum, v)
// 	}

// 	for _, v := range m {
// 		if v != minNum && v != maxNum {
// 			count++
// 		} else if count > 0 && v == minNum {
// 			return false
// 		}
// 	}

// 	return count == 0
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func min(a, b int) int {
// 	if a > b {
// 		return b
// 	}
// 	return a
// }

// func canPartitionKSubsets(nums []int, k int) bool {
// 	sum := 0
// 	for _, v := range nums {
// 		sum += v
// 	}
// 	if sum%k != 0 {
// 		return false
// 	}

// 	vis := make([]bool, len(nums))
// 	target := sum / k

// 	var dfs func(start, cur, count int) bool
// 	dfs = func(start, cur, count int) bool {
// 		if count == k {
// 			return true
// 		}

// 		if cur > target {
// 			return false
// 		}

// 		for i := start; i < len(nums); i++ {
// 			if !vis[i] {
// 				vis[i] = true
// 				if cur+nums[i] == target {
// 					dfs(i+1, 0, count+1)
// 				} else if cur+nums[i] < target {
// 					dfs(i+1, cur+nums[i], count)
// 				}
// 				vis[i] = false
// 			}
// 		}
// 		return false
// 	}

// 	ret := dfs(0, 0, 0)
// 	return ret
// }

// func main() {
// 	nums := []int{1, 1, 1, 1, 2, 2, 2, 2}
// 	k := 4
// 	ret := canPartitionKSubsets(nums, k)
// 	fmt.Printf("%v", ret)
// }

// func minNumber(nums1 []int, nums2 []int) int {
// 	sort.Ints(nums1)
// 	sort.Ints(nums2)

// 	val := union(nums1, nums2)
// 	if val != 0 {
// 		return val
// 	}

// 	if nums1[0] < nums2[0] {
// 		return nums1[0]*10 + nums2[0]
// 	} else {
// 		return nums2[0]*10 + nums1[0]
// 	}
// }

// func union(nums1 []int, nums2 []int) int {
// 	m := map[int]bool{}
// 	for i := 0; i < len(nums1) || i < len(nums2); i++ {
// 		if i < len(nums1) {
// 			if m[nums1[i]] {
// 				return nums1[i]
// 			}
// 			m[nums1[i]] = true
// 		}
// 		if i < len(nums2) {
// 			if m[nums2[i]] {
// 				return nums2[i]
// 			}
// 			m[nums2[i]] = true
// 		}
// 	}
// 	return 0
// }

// func main() {
// 	nums1 := []int{3, 5, 2, 6}
// 	nums2 := []int{3, 1, 7}
// 	minNumber(nums1, nums2)
// }

// func main() {
// 	var i int8 = 127
// 	fmt.Println(i, i+1, i*2)
// }

// 100 4
// 1000 16
// 10000 32
// 100000 64
// 1000000 128
// func main() {
// 	slice := make([]int, 5)
// 	println(len(slice))
// 	println(cap(slice))
// }

// func test1() bool {
// 	a := false
// 	defer func() {
// 		a = true
// 	}()
// 	return a
// }
// func test2() (a bool) {
// 	a = false
// 	defer func() {
// 		a = true
// 	}()
// 	return a
// }
