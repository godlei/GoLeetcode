package leetcode

func letterCombinations(digits string) []string {

	if len(digits) < 1 {
		return []string{}
	}

	maps := make(map[int]string)
	maps[2] = "abc"
	maps[3] = "def"
	maps[4] = "ghi"
	maps[5] = "jkl"
	maps[6] = "mno"
	maps[7] = "pqrs"
	maps[8] = "tuv"
	maps[9] = "wxyz"

	res := fun(digits, maps)

	return res
}

func fun(digits string, maps map[int]string) (result []string) {
	if len(digits) < 1 {
		return result
	}

	if len(digits) == 1 {
		ch := digits[0]
		n := int(ch - '0')

		str := maps[n]
		for i := 0; i < len(str); i++ {
			result = append(result, str[i:i+1])
		}
		return result
	}

	data := fun(digits[1:], maps)
	ch := digits[0]
	n := int(ch - '0')

	str := maps[n]
	for i := 0; i < len(str); i++ {
		for j := 0; j < len(data); j++ {
			result = append(result, str[i:i+1]+data[j])
		}
	}
	return result
}
