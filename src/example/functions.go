package example

/*
	函数相关
*/
import "fmt"

func FunctionTest() {
	/* 函数变量*/
	var f func(int, int) int
	x := 4
	y := 3

	f = add
	fmt.Println(f(x, y))
	f = sub
	fmt.Println(f(x, y))
	f = mul
	fmt.Println(f(x, y))
	f = div
	fmt.Println(f(x, y))

	/*  匿名函数	*/
	fmt.Println(
		func(x int, y int) (res int) {
			res = 2*x + y
			return
		}(3, 4))
}

func add(x int, y int) (res int) {
	res = x + y
	return
}
func sub(x int, y int) (res int) {
	res = x - y
	return
}
func mul(x int, y int) (res int) {
	res = x * y
	return
}
func div(x int, y int) (res int) {
	res = x / y
	return
}
