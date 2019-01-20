package example

import (
	"fmt"
	"time"
)

func GororoutineExample() {
	fmt.Println("main start")

	ch := make(chan int)

	go fun2(ch, 100)

	num := 0
	num = <-ch
	fmt.Println(num)
	fmt.Println("main end")
}

func fun1() {
	fmt.Println(time.Now())

	time.Sleep(time.Second * 5)

	fmt.Println(time.Now())
}

func fun2(ch chan int, num int) {
	fmt.Println("fun2 start")
	time.Sleep(time.Second * 3)
	res := 0
	for i := 1; i <= num; i++ {
		res += i
	}

	ch <- res
	fmt.Println("func2 end")
}
