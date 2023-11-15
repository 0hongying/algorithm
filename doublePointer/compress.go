package main

import (
	"strconv"
)

//从一个空字符串 s 开始。对于 chars 中的每组 连续重复字符
// 压缩后得到的字符串 s 不应该直接返回 ，需要转储到字符数组 chars 中。需要注意的是，如果组长度为 10 或 10 以上，则在 chars 数组中会被拆分为多个字符

// func compress(chars []byte) int {
// 	// left:子字符串最左侧
// 	// write:要写入字符的下标
// 	left, write := 0, 0
// 	for i, c := range chars {
// 		if i == len(chars)-1 || c != chars[i+1] {
// 			chars[write] = c
// 			num := i - left + 1
// 			left = i + 1
// 			write++
// 			if num > 1 {
// 				// 将数字变成字符串来记录
// 				numStr := strconv.Itoa(num)
// 				for _, val := range numStr {
// 					chars[write] = byte(val)
// 					write++
// 				}
// 			}
// 		}
// 	}
// 	return write
// }

func compress(chars []byte) int {
	// 写入的位置
	write := 0
	n := len(chars)
	for i := 0; i < n; i++ {
		count := 1
		v := chars[i]
		// 统计重复的字母
		for i < n-1 && chars[i+1] == v {
			i++
			count++
		}
		// 写入字母
		chars[write] = v
		write++
		// 写入次数
		if count > 1 {
			numStr := strconv.Itoa(count)
			for _, v := range numStr {
				chars[write] = byte(v)
				write++
			}
		}

	}
	return write
}

// func main() {
// 	char := []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
// 	ret := compress(char)
// 	fmt.Printf("%v", ret)
// }
