package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(time.Second * 2)
	<-timer1.C
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)

	go func() {
		fmt.Println("-------")

		for t := range timer2.C {
			//<-timer2.C
			fmt.Println("Timer 2 expired", t)
		}
	}()
	time.Sleep(time.Second * 10)
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
