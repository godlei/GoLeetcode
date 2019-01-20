package leetcode

import "fmt"

func Problem6() {
	fmt.Println(convert("ab", 1))
}

func convert(s string, numRows int) string {
	length := len(s)
	if length <= 1 {
		return s
	}
	if numRows == 1 {
		return s
	}
	var result []byte
	for i := 0; i < numRows; i++ {
		if i == 0 {
			n1 := 0
			n2 := 2 * (numRows - 1)
			result = append(result, []byte(s[n1 : n1+1])[0])
			for n2 < length {
				//fmt.Printf("n2:%d\n",n2)
				result = append(result, []byte(s[n2 : n2+1])[0])
				n2 = n2 + 2*(numRows-1)
				//fmt.Println(string(result))
			}
		} else if i == numRows-1 {
			n1 := numRows - 1
			if n1 >= length {
				continue
			}
			result = append(result, []byte(s[n1 : n1+1])[0])
			n2 := 2*(numRows-1) + n1
			for n2 < length {
				//fmt.Printf("n2:%d\n",n2)
				result = append(result, []byte(s[n2 : n2+1])[0])
				n2 = n2 + 2*(numRows-1)
				//fmt.Println(string(result))
			}
		} else {
			n1 := i
			n2 := n1 + 2*(numRows-1-i)

			for n1 < length && n2 < length {
				result = append(result, []byte(s[n1 : n1+1])[0])
				result = append(result, []byte(s[n2 : n2+1])[0])
				n1 = n1 + 2*(numRows-1)
				n2 = n2 + 2*(numRows-1)
			}
			if n1 < length {
				result = append(result, []byte(s[n1 : n1+1])[0])
			}
		}

	}
	res := string(result)
	return res
}

//LDR EOEII ECIHN TSG
//LDR EOEII ECIHN TSG
