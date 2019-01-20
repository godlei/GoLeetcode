package leetcode

import "fmt"

func Problem3() {
	l := lengthOfLongestSubstring("pwwkew")
	fmt.Println(l)
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	length := len(s)

	array := make([]int, length)
	max := 0
	for i := 0; i < length; i++ {
		if i == 0 {
			array[0] = 1
			continue
		}

		val := 1

		for j := i - 1; j >= i-array[i-1]; j-- {

			if s[i:i+1] == s[j:j+1] {
				break
			}
			val++

		}

		array[i] = val
		if val > max {
			max = val
		}

		//fmt.Println(array)
	}

	return max
}
