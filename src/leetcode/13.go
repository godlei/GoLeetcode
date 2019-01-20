package leetcode

import "fmt"

func Problem13() {
	str := "MCMXCIV"
	fmt.Println(romanToInt(str))
}

func romanToInt(s string) int {
	if len(s) <= 0 {
		return 0
	}

	maps := make(map[string]int)
	maps["I"] = 1
	maps["V"] = 5
	maps["X"] = 10
	maps["L"] = 50
	maps["C"] = 100
	maps["D"] = 500
	maps["M"] = 1000

	if len(s) == 1 {
		return maps[s]
	}
	var res int

	i := 0
	for i < len(s) {
		j := i + 1
		if j == len(s) {
			res += maps[s[i:i+1]]
			break
		}

		if maps[s[i:i+1]] >= maps[s[j:j+1]] {
			res += maps[s[i:i+1]]
			i++
		} else {
			res += maps[s[j:j+1]] - maps[s[i:i+1]]
			i = j + 1
		}

	}

	return res
}
