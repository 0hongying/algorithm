package main

// 给定一个经过编码的字符串，返回它解码后的字符串。
// 编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数

func decodeString(s string) string {
	numStack := []int{}
	strStack := []string{}
	sum := 0

	for _, v := range s {
		if v >= '0' && v <= '9' {
			sum = sum*10 + int(v-'0')
		} else if v == '[' {
			strStack = append(strStack, string(v))
			numStack = append(numStack, sum)
			sum = 0
		} else if v >= 'a' && v <= 'z' || v >= 'A' && v <= 'Z' {
			strStack = append(strStack, string(v))
		} else {
			temp := ""
			sTemp := ""
			for len(strStack) > 0 && strStack[len(strStack)-1] != "[" {
				str := strStack[len(strStack)-1]
				strStack = strStack[:len(strStack)-1]
				temp = str + temp
			}
			strStack = strStack[:len(strStack)-1]
			num := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			for i := 0; i < num; i++ {
				sTemp = temp + sTemp
			}
			strStack = append(strStack, sTemp)
		}
	}

	ans := ""
	for _, v := range strStack {
		ans += v
	}
	return ans
}

// func main() {
// 	s := "3[a2[c]]"
// 	ret := decodeString(s)
// 	fmt.Printf("%v", ret)
// }
