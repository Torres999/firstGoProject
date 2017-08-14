package main

import "fmt"
import "math"

const s1 string = "constant"
const (
	s = "constantstring"
)

func main() {
	fmt.Println(s)

	const n = 50000
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

	//basicTypeFunc()
}

func basicTypeFunc() {
	var a string = "dalin\n"
	fmt.Printf("hello " + a)

	var c, b int = 1, 2
	fmt.Println("c + b =", c+b)
	fmt.Println("true && false =", true && false)
	fmt.Println("true || false =", true || false)
	fmt.Println("!true =", !true)

	f := "shorts"
	fmt.Println(f)

	d := 1
	fmt.Println(c + b + d)
}
