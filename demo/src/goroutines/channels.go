package main

import (
	"fmt"
)

func main() {
	// Create a new channel with make(chan val-type). Channels are typed by the values they convey.
	messages := make(chan string)

	// Send a value into a channel using the channel <- syntax.
	// Here we send "ping" to the messages channel we made above, from a new goroutine.
	var str string
	fmt.Scanln(&str)
	go func() {
		fmt.Println("======")
		messages <- str
		messages <- "ping"//不会执行
	}()

	// The <-channel syntax receives a value from the channel.
	// Here we’ll receive the "ping" message we sent above and print it out.
	msg := <-messages
	fmt.Println(msg)
	// By default sends and receives block until both the sender and receiver are ready.
	// This property allowed us to wait at the end of our program for the "ping" message
	// without having to use any other synchronization.

	fmt.Println("开启5个routine,先到先出")//channelDirections.go
	// 开启5个routine
	for i := 0; i < 5; i++ {
		go foo(i)
	}

	// 取出信道中的数据
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
}

var ch chan int = make(chan int)

func foo(id int) {
	//id: 这个routine的标号
	fmt.Println("id:", id)
	ch <- id
}