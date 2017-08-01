package main

import "fmt"
import "time"

func main() {
	forFunc()
	fmt.Println("=============")
	ifElseFunc()
	switchFunc()
}

func forFunc() {
	i := 1

	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n % 2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

func ifElseFunc() {
	if 7 % 2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8 % 4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if xcd := 9111; xcd > 0 {
		fmt.Println(xcd, "is negative")
	} else if xcd < 10 {
		fmt.Println(xcd, "has 1 digit")
	} else {
		fmt.Println(xcd, "has multiple digits")
	}
}

func switchFunc() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")

	}

}
