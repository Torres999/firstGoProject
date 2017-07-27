package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	//stringsFunc()
	strconvFunc()
}

/*
strings
 */
func stringsFunc() {
	s := "hello world"
	fmt.Println(s)

	fmt.Println(strings.Contains(s, "hello"))
	fmt.Println(strings.Index(s, "h"))

	var s1 string
	s1 = "123*134*36*678"
	fmt.Println(strings.Split(s1, "*")[1]) //134
}

/*
strconv
 */
func strconvFunc() {

	fmt.Println(strconv.Itoa(123))

	a, errorMsg := strconv.Atoi("123h123")
	fmt.Println("value:", a)
	fmt.Println("errorMsg:", errorMsg)

	// 进制转换
	fmt.Println(strconv.FormatInt(123, 2))

}