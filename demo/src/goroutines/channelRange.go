package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3

	// 1和2至少存在一个即可不会报死锁的错，用close更优雅
	// 1.显式地关闭信道
	close(ch)//被关闭的信道会禁止数据流入, 是只读的。我们仍然可以从关闭的信道中取出数据，但是不能再写入数据了。

	for v := range ch {// 其他遍历方式参考channelsRangeOverChannels.go、channelRangeOverChannels.go
		fmt.Println(v)
		// 2.如果不加这个判断，会报死锁错误的，原因是range不等到信道关闭是不会结束读取的。也就是缓冲信道干涸了range还在读
		if len(ch) <= 0 {
			// 如果现有数据量为0，跳出循环
			break
		}
	}
}
