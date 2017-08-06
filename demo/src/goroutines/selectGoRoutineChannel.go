package main

import (
	"time"
	"fmt"
)

/*
	select:信道
 */
func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		fmt.Println("====", 1)
		time.Sleep(time.Second * 2)
		c1 <- "one"
	}()

	go func() {
		fmt.Println("====", 2)
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	//注意，如果没有default，select 会一直等待等到某个 case 语句完成， 也就是等到成功从 ch 或者 timeout 中读到数据，否则一直阻塞。
	for i := 0; i < 2; i++ {//大于2会死锁，因为两个chan都已经关闭了
		//加for只是为了结束第一次阻塞后再执行一次
		fmt.Println("====", 3)
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		// 放开注解后1、2不会在执行
		//default:
		//	fmt.Println("default case is running")
		}
	}

	//判断是否满
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	select {
	case <-ch1:
		fmt.Println("ch1 pop one element")
	case <-ch2:
		fmt.Println("ch2 pop one element")
	default:
		fmt.Println("default")
	}
}