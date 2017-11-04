package main

import (
	"fmt"
	"time"
)

//This ping function only accepts a channel for sending values.
// It would be a compile-time error to try to receive on this channel.
func ping(pings chan<- string, msg string) {
	pings <- msg
}

//The pong function accepts one channel for receives (pings) and a second for sends (pongs).
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	//pings := make(chan string, 1)
	//pongs := make(chan string, 1)
	//
	//ping(pings, "passed message")
	////fmt.Println(<-pings)//passed message
	//
	//pong(pings, pongs)
	//fmt.Println(<-pongs)

	count := 10
	quit = make(chan int) // 无缓冲
	for i := 0; i < count; i++ {
		go foo(i)
	}
	time.Sleep(time.Second)
	for i := 0; i < 9; i++ { //i<=10都可以，超过count就会报死锁的错，因为通道干涸了
		<-quit
	}
}

var quit chan int // 只开一个信道
func foo(id int) {
	fmt.Println(id)
	quit <- 0 // ok, finished
}
