package leetcode

import (
	"fmt"
	"strconv"
)

func Problem9() {
	fmt.Println(isPalindrome(1001))
}
func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	length := len(str)
	i := 0
	j := length - 1
	for {
		if i > j {
			break
		}
		s1 := str[i : i+1]
		s2 := str[j : j+1]

		fmt.Println("..")
		if s1 != s2 {
			return false
		}
		i++
		j--
	}
	return true
}
