package main

import "fmt"

func f(str string) {
	for i := 0; i < 3; i++ {
		fmt.Println(str, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutine")
	var input string

	fmt.Scanln(&input)

	//使用了go，所以没打印出来going，如果去掉go就会直接打印going
	go func(msg string) {
		fmt.Scanln(&input)
		fmt.Println(msg)
	}("going") //():匿名函数会马上执行

	//var input string
	//fmt.Scanln(&input)
	//fmt.Println("input from console:", input)
	//fmt.Println("done")
}
