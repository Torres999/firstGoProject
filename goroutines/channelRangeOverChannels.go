package main

import "fmt"

func main() { // 其他遍历方式参考channelsRangeClosingChannels.go、channelRange.go
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)
	for elem := range queue {
		fmt.Println(elem)
	}
}
