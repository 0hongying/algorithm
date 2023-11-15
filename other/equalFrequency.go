package main

import (
	"sort"
)

// 给你一个下标从 0 开始的字符串 word ，字符串只包含小写英文字母。你需要选择 一个 下标并 删除 下标处的字符，使得 word 中剩余每个字母出现 频率 相同

func equalFrequency(word string) bool {
	mCnt := map[rune]int{}
	for _, c := range word {
		mCnt[c]++
	}
	cnt := make([]int, 0, len(mCnt))
	for _, c := range mCnt {
		cnt = append(cnt, c)
	}
	sort.Ints(cnt) // 出现次数从小到大排序
	m := len(cnt)
	// 只有一种字符 or 去掉次数最少的（最少次数为1） or 去掉次数最多的（其他次数都为最多-1）
	return m == 1 ||
		cnt[0] == 1 && isSame(cnt[1:]) ||
		cnt[m-1] == cnt[m-2]+1 && isSame(cnt[:m-1])
}

func isSame(a []int) bool {
	for _, x := range a[1:] {
		if x != a[0] {
			return false
		}
	}
	return true
}

// func main() {
// 	word := "abcc"
// 	ret := equalFrequency(word)
// 	fmt.Printf("%v", ret)
// }
