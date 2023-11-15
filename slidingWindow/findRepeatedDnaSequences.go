package main

// 给定一个表示DNA序列的字符串s，返回所有在 DNA 分子中出现不止一次的长度为10的序列(子字符串)
func findRepeatedDnaSequences(s string) []string {
	m := map[string]int{}
	ret := []string{}
	for i := 0; i <= len(s)-10; i++ {
		key := s[i : i+10]
		m[key]++
		if m[key] == 2 {
			ret = append(ret, key)
		}
	}
	return ret
}

// func main() {
// 	s := "AAAAAAAAAAAAA"
// 	ret := findRepeatedDnaSequences(s)
// 	fmt.Printf("%v", ret)
// }
