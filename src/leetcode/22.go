package leetcode

import "fmt"

func Problem22() {
	fmt.Println(generateParenthesis(3))
}
func generateParenthesis(n int) []string {
	a, b := n, n

	res := fun22(a, b)
	return res
}

func fun22(a int, b int) (res []string) {
	//fmt.Printf("a:%d,b:%d\n",a,b)
	if a == 0 {
		byteRes := []byte{}
		for i := 0; i < b; i++ {
			byteRes = append(byteRes, ')')
		}
		res = append(res, string(byteRes))
		return
	}

	if a == b {
		data := fun22(a-1, b)
		for i := 0; i < len(data); i++ {
			str := data[i]
			str = "(" + str
			res = append(res, str)
		}
	} else {
		data := fun22(a-1, b)
		for i := 0; i < len(data); i++ {
			str := data[i]
			str = "(" + str
			res = append(res, str)
		}

		data = fun22(a, b-1)
		for i := 0; i < len(data); i++ {
			str := data[i]
			str = ")" + str
			res = append(res, str)
		}
	}

	return
}
