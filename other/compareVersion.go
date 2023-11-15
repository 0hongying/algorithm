package main

import (
	"strconv"
	"strings"
)

// 比较版本是否一致
// version1 = "1.01", version2 = "1.001"  ==> 0
// version1 = "1.0", version2 = "1.0.0"   ==> 0
// version1 = "0.1", version2 = "1.1"     ==> -1

func compareVersion(version1 string, version2 string) int {
	str1 := strings.Split(version1, ".")
	str2 := strings.Split(version2, ".")
	for i, j := 0, 0; i < len(str1) || j < len(str2); i, j = i+1, j+1 {
		num1, num2 := 0, 0
		if i < len(str1) {
			num1, _ = strconv.Atoi(str1[i])
		} else {
			num1 = 0
		}
		if j < len(str2) {
			num2, _ = strconv.Atoi(str2[j])
		} else {
			num2 = 0
		}

		if num1 < num2 {
			return -1
		} else if num1 > num2 {
			return 1
		}
	}
	return 0
}

// func main() {
// 	version1 := "1"
// 	version2 := "1.0.0"
// 	ret := compareVersion(version1, version2)
// 	fmt.Printf("%v", ret)
// }
