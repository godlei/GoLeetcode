package leetcode

import (
	"fmt"
	"math"
)

func Problem7() {
	//res := 9646324351
	res := reverse(-2147483648)
	fmt.Println(res)
}

func reverse(x int) int {
	res := 0
	for {
		if x == 0 {
			break
		}
		//fmt.Printf("x:%d\n",x)
		n1 := x % 10
		//fmt.Printf("n1:%d\n",n1)
		if res > (math.MaxInt32-n1)/10 {
			return 0
		}
		if res < (math.MinInt32-n1)/10 {
			return 0
		}
		res = res*10 + n1
		//fmt.Printf("res:%d\n",res)

		//fmt.Println(res)
		x /= 10
	}

	return res

}
