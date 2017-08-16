package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	stringsFunc()
	strconvFunc()
}

/*
strings
*/
func stringsFunc() {
	s := "hello world"
	fmt.Println(s)

	fmt.Println("Contains:", strings.Contains(s, "hello"))
	fmt.Println("Count:", strings.Count(s, "o"))
	fmt.Println("Index:", strings.Index(s, "h"))

	var s1 string
	s1 = "123*134*36*678"
	fmt.Println("Split & Get Arrays:", strings.Split(s1, "*")[1]) //134
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
