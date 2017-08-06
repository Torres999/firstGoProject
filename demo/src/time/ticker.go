package main

import (
	"time"
	"fmt"
)

func main() {
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println(t)
		}
	}()

	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println("Ticker stopped")


	/*
	Ticker和Timer的区别：timer只会执行一次
	 */
	//ticker := time.NewTimer(time.Second )
	//go func() {
	//	for t := range ticker.C {
	//		fmt.Println(t)
	//	}
	//}()
	//
	//time.Sleep(time.Second * 12)
	//ticker.Stop()
	//fmt.Println("Ticker stopped")
}