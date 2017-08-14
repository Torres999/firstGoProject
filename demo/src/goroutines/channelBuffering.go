package main

import "fmt"

/*
 By default channels are unbuffered, meaning that they will only accept sends (chan <-)
 if there is a corresponding receive (<- chan) ready to receive the sent value.
 Buffered channels accept a limited number of values without a corresponding receiver for those values.
*/
func main() {
	messages := make(chan string, 2)

	var str, str1 string
	fmt.Scanln(&str)
	fmt.Scanln(&str1)

	messages <- str
	fmt.Println(<-messages)
	messages <- str1
	fmt.Println(<-messages)

	/*
		缓冲信道会在满容量的时候加锁。如果不把容量中的取走，不管有没有取数据的代码都会报错 - deadlock!
	*/
	messages <- "123"
	messages <- "456"
	//messages <- "456"//fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
