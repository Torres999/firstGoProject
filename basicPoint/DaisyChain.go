package main

import "fmt"

func f(left, right chan int) {
	left <- 1 + <-right
}
func main() {
	const count = 3
	leftmost := make(chan int)
	right := leftmost
	left := leftmost

	fmt.Println("right0:", right)
	for i := 0; i < count; i++ {
		right = make(chan int)
		fmt.Println("right:", right)

		go f(left, right)

		left = right
	}
	go func(c chan int) {
		fmt.Println("right1:", right)
		c <- 1
	}(right)

	fmt.Println(<-leftmost)
}
