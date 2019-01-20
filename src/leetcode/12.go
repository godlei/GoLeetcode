package leetcode

import (
	"fmt"
	"strings"
)

func Problem12() {
	fmt.Println(intToRoman(30))
}

func intToRoman(num int) string {
	s1 := []string{"", "M", "MM", "MMM"}
	s2 := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	s3 := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	s4 := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

	n1 := num / 1000
	num = num % 1000
	n2 := num / 100
	num = num % 100
	n3 := num / 10
	n4 := num % 10
	res := s1[n1] + s2[n2] + s3[n3] + s4[n4]
	res = strings.Trim(res, " ")

	return res
}
