package main

import "fmt"

func main() {
	// first step in Golang
	fmt.Printf("hello, world\n")

	// 2
	key := "a"
	value := "apple"
	fmt.Printf("%s -> %s\n", key, value)

	// 3
	res := plus(1, 2)
	fmt.Println("1+2 =", res)
	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	//4
	sum(1, 2, 3)

	//5
	a, b := multipleRtn()
	fmt.Println(a)
	fmt.Println(b)
	_, c := multipleRtn()
	fmt.Println(c)
}

/*
方法／返回值
结构体或者接口的实现方法格式：参数／方法／返回值
 */
func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

/*
Multiple Return Values
 */
func multipleRtn() (int, int) {
	return 3, 7
}

/*
Variadic Functions:多参数方法
 */
func sum(nums ...int) {
	fmt.Println("nums:", nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("total:", total)
}