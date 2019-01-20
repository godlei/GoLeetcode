package leetcode

import "fmt"

func Problem5() {
	fmt.Println(longestPalindrome("ab"))
}
func longestPalindrome(s string) string {

	if len(s) <= 1 {
		return s
	}
	if len(s) == 2 && s[0:1] == s[1:2] {
		return s
	}
	var from, to, max int

	for i := 0; i < len(s); i++ {
		//fmt.Println("aaa")
		start := i - 1
		end := i + 1
		//flag := false
		for start >= 0 && end < len(s) {
			if s[start:start+1] == s[end:end+1] {
				start--
				end++
			} else {
				break
			}
		}
		length := end - start - 1
		if length > max {
			from = start + 1
			to = end - 1
			max = length
		}

		if i+1 < len(s) {

			if s[i:i+1] == s[i+1:i+2] {
				start = i - 1
				end = i + 2
				for start >= 0 && end < len(s) {
					if s[start:start+1] == s[end:end+1] {
						start--
						end++
					} else {
						break
					}
				}
				length := end - start - 1
				if length > max {
					from = start + 1
					to = end - 1
					max = length
				}
			} else {
				continue
			}
		}

	}
	return s[from : to+1]
}
