package leetcode

func isValid(s string) bool {
	if len(s) <= 0 {
		return true
	}

	if len(s) == 1 {
		return false
	}

	var stack []byte

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else if s[i] == ')' {
			if len(stack) != 0 || stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else if s[i] == ']' {
			if len(stack) != 0 || stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else if s[i] == '}' {
			if len(stack) != 0 || stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			continue
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true

}
