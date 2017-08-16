package main

import "fmt"

//func f(left, right chan int) {
//	i := <-right
//	fmt.Println("    i:", i)
//	left <- 1 + i
//}
func main() {
	const count = 3
	leftmost := make(chan int)
	right := leftmost
	left := leftmost
	fmt.Println("leftmost:", leftmost)
	fmt.Println("=====----")

	for i := 0; i < count; i++ {
		right = make(chan int)
		go func(left1, right1 chan int) {
			i := <-right1
			fmt.Println("    i:", i)
			left1 <- 1 + i
		}(left, right)
		left = right

		fmt.Println("leftmost:", leftmost)
		fmt.Println("   right:", right)
		fmt.Println("    left:", left)
		fmt.Println()
	}
	go func(c chan int) {
		fmt.Println("==============")
		c <- 1
		fmt.Println("--------------")
	}(right)

	fmt.Println("--------------")
	fmt.Println("leftmost:", leftmost)
	fmt.Println("   right:", right)
	fmt.Println("    left:", left)
	fmt.Println()
	fmt.Println(<-leftmost)
}
