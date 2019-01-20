package leetcode

import (
	"fmt"
	"math"
	"strings"
)

func Problem8() {
	fmt.Println(math.MinInt32)
	//var str = "12"
	fmt.Println(myAtoi("   -42"))

}

func myAtoi(str string) int {
	str = strings.Trim(str, " ")
	if len(str) < 1 {
		return 0
	}
	isPositive := true
	if str[0] == '-' {
		str = str[1:]
		isPositive = false
	} else if str[0] == '+' {
		str = str[1:]
	}
	if len(str) < 1 {
		return 0
	}
	if str[0] < '0' || str[0] > '9' {
		return 0
	}
	var num []byte
	for i := 0; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			num = append(num, str[i])
		} else {
			break
		}
	}

	res := 0
	if isPositive {
		for i := 0; i < len(num); i++ {
			n := int(num[i]) - int('0')
			if res > (math.MaxInt32-n)/10 {
				return math.MaxInt32
			}
			res = res*10 + int(n)
		}
	} else {
		//fmt.Println(string(num))
		for i := 0; i < len(num); i++ {
			n := int('0') - int(num[i])

			//fmt.Println(int(n))\

			if res < (math.MinInt32-n)/10 {
				//fmt.Println("ssss")
				return math.MinInt32
			}
			res = res*10 + int(n)
			//fmt.Printf("n:%d\n",n)
			//fmt.Printf("res:%d\n",res)

		}
	}
	return res
}
