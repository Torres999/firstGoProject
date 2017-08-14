package main

import "fmt"

func intSeq(a int) func() int {
	if a == 1 {
		return func() int {
			return a * 2
		}
	} else {
		return func() int {
			return a * 10
		}
	}
}

func intSeq1() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

/**
闭包里的非传递参数外部变量值是传引用的，在闭包函数里那个i就是外部非闭包函数自己的参数，所以是相当于引用了外部的变量
*/
func main() {
	nextInt0 := intSeq1()
	fmt.Println(nextInt0())
	fmt.Println(nextInt0())
	fmt.Println(nextInt0())

	nextInt1 := intSeq(1)
	nextInt2 := intSeq(2)

	fmt.Println("nextInt1:", nextInt1())
	fmt.Println("nextInt2:", nextInt2())
}
