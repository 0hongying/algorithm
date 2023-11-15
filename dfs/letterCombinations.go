package main

func letterCombinations(digits string) []string {
	ans := []string{}
	path := make([]byte, len(digits))

	if len(digits) == 0 {
		return ans
	}

	phoneMap := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	var dfs func(index int)
	dfs = func(index int) {
		if index == len(digits) {
			ans = append(ans, string(path))
			return
		}
		d := digits[index]
		for _, v := range phoneMap[string(d)] {
			path[index] = byte(v)
			dfs(index + 1)
		}

	}

	dfs(0)
	return ans
}

// func main() {
// 	digits := "23"
// 	letterCombinations(digits)
// }
