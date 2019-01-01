package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "abcdefg higklin"
	s1 := strings.LastIndex(str, " ")
	fmt.Println(str[s1+1:])

}
