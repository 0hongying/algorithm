package main

import (
	"strconv"
	"strings"
)

// 1.     1
// 2.     11
// 3.     21
// 4.     1211
// 5.     111221
// 第一项是数字 1
// 描述前一项，这个数是 1 即 “ 一 个 1 ”，记作 "11"
// 描述前一项，这个数是 11 即 “ 二 个 1 ” ，记作 "21"
// 描述前一项，这个数是 21 即 “ 一 个 2 + 一 个 1 ” ，记作 "1211"
// 描述前一项，这个数是 1211 即 “ 一 个 1 + 一 个 2 + 二 个 1 ” ，记作 "111221"

func countAndSay(n int) string {
	prev := "1"
	for i := 2; i <= n; i++ {
		cur := &strings.Builder{}
		// 记录个数
		count := 1
		for j := 0; j < len(prev); j++ {
			if j < len(prev)-1 && prev[j] == prev[j+1] {
				count++
				continue
			}
			cur.WriteString(strconv.Itoa(count))
			cur.WriteByte(prev[j])
			count = 1
		}
		prev = cur.String()
	}
	return prev
}

// func main() {
// 	ret := countAndSay(6)
// 	fmt.Printf("%v", ret)
// }
