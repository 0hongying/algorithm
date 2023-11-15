package main

import (
	"strconv"
)

// 除法
func fractionToDecimal(numerator, denominator int) string {
	// 刚好整除
	if numerator%denominator == 0 {
		return strconv.Itoa(numerator / denominator)
	}

	s := []byte{}
	// 先判断结果是否为负数
	if numerator*denominator < 0 {
		s = append(s, '-')
	}

	numerator = abs(numerator)
	denominator = abs(denominator)

	// 整数部分
	integerPart := numerator / denominator
	s = append(s, strconv.Itoa(integerPart)...)
	s = append(s, '.')

	// 小数部分
	indexMap := map[int]int{}
	remainder := numerator % denominator

	for remainder != 0 {
		// 当余数重复时，则证明是循环小数
		if _, ok := indexMap[remainder]; ok {
			break
		}
		indexMap[remainder] = len(s)
		remainder *= 10
		s = append(s, '0'+byte(remainder/denominator))
		remainder %= denominator
	}
	if remainder > 0 { // 有循环节
		insertIndex := indexMap[remainder]
		s = append(s[:insertIndex], append([]byte{'('}, s[insertIndex:]...)...)
		s = append(s, ')')
	}

	return string(s)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// func main() {
// 	numerator := -50
// 	denominator := 8
// 	ret := fractionToDecimal(numerator, denominator)
// 	fmt.Printf("%v", ret)
// }
