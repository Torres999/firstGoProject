package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)
	//time.Sleep(time.Second*2)
	<-done //If you removed the <- done line from this program, the program would exit before the worker even started.

	// 默认的，信道的存消息和取消息都是阻塞的,也就是说, 无缓冲的信道在取消息和存消息的时候都会挂起当前的goroutine，除非另一端已经准备好。
	// 直到线程跑完, 取到消息. main在此阻塞住, "<-done"的目的就是取消息，以达到阻塞的目的
	// 如果不用信道来阻塞主线的话，主线就会过早跑完，go worker(done)线都没有机会执行
	/*
		其实，无缓冲的信道永远不会存储数据，只负责数据的流通，为什么这么讲呢？
		  从无缓冲信道取数据，必须要有数据流进来才可以，否则当前线阻塞
		  数据流入无缓冲信道, 如果没有其他goroutine来拿走这个数据，那么当前线阻塞
	*/
}

/*
==============================================================================
	死锁：fatal error: all goroutines are asleep - deadlock!
		非缓冲信道上如果发生了流入无流出，或者流出无流入，也就导致了死锁。
	解决办法：
		1.所有进出成对出现
		2.添加一个缓冲区，不达容量不阻塞。达容量后再放一个会报错，缓冲信道会在满容量的时候加锁
==============================================================================
// =============1=============
// 只进不出
func main() {
	ch := make(chan int)
	ch <- 1 // 1流入信道，堵塞当前线, 没人取走数据信道不会打开
	fmt.Println("This line code wont run") //在此行执行之前Go就会报死锁
}

// =============2=============
// 只出不进：ch2
var ch1 chan int = make(chan int)
var ch2 chan int = make(chan int)
func say(s string) {
	fmt.Println(s)
	ch1 <- <- ch2 // ch1 等待 ch2流出的数据
}
func main() {
	go say("hello")
	<- ch1  // 堵塞主线
}

// =============3=============
// 不会死锁，在go执行前main就跑完了
func main() {
    c := make(chan int)

    go func() {
       c <- 1
    }()
}

*/
