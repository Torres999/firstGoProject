package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("!")

	os.Exit(3)
}

/*
两种运行方式
 */
//$ go run exit.go
//exit status 3

//$ go build exit.go
//$ ./exit
//$  echo  $？
//3
