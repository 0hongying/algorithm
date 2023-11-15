package main

import (
	"strconv"
)

const SEG_COUNT = 4

// 在字符串中找到所有符合规范的ip地址
func restoreIpAddresses(s string) []string {
	ret := []string{}
	segment := make([]int, SEG_COUNT)
	var dfs func(s string, segId, segIndex int)
	dfs = func(s string, segId, segIndex int) {
		if segId == SEG_COUNT {
			// 等于四段 && 字符串用完
			if segIndex == len(s) {
				ipAddr := ""
				for i := 0; i < SEG_COUNT; i++ {
					ipAddr += strconv.Itoa(segment[i])
					if i != SEG_COUNT-1 {
						ipAddr += "."
					}
				}
				ret = append(ret, ipAddr)
			}
			// 不符合跳出
			return
		}

		if segIndex == len(s) {
			return
		}

		// 前导为0的话，ip地址只能为0
		if s[segIndex] == '0' {
			segment[segId] = 0
			dfs(s, segId+1, segIndex+1)
		}

		addr := 0
		for i := segIndex; i < len(s); i++ {
			addr = addr*10 + int(s[i]-'0')
			if addr > 0 && addr <= 255 {
				segment[segId] = addr
				// 从 i 开始遍历
				dfs(s, segId+1, i+1)
			} else {
				// 不符合跳出
				return
			}
		}
	}
	dfs(s, 0, 0)
	return ret
}

// func main() {
// 	s := "2552325523"
// 	ret := restoreIpAddresses(s)
// 	fmt.Printf("%v", ret)
// }
