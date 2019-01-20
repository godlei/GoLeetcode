package leetcode

import "fmt"

func Problem14() {
	strs := []string{"flower", "flow", "flight"}

	str := longestCommonPrefix(strs)
	fmt.Println(str)

}
func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	var minStr string
	for i := 0; i < len(strs); i++ {
		if i == 0 {
			minStr = strs[0]
			continue
		}
		s := strs[i]
		if len(s) < len(minStr) {
			minStr = s
		}
	}

	var byteRes []byte
	for i := 0; i < len(minStr); i++ {

		finished := false
		for j := 0; j < len(strs) && finished == false; j++ {
			if minStr[i] != strs[j][i] {
				finished = true
				break
			}
		}
		if finished == false {
			byteRes = append(byteRes, minStr[i])
		} else {
			break
		}
	}
	return string(byteRes)
}
