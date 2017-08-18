package main

import (
	"fmt"
)

type Pet interface {
	Name() string
	Age() uint8
}

type Dog struct {
	name string
	age  uint8
}

func (dog Dog) Name() string {
	return dog.name
}

func (dog Dog) Age() uint8 {
	return dog.age
}

func main() {
	myDog := Dog{"Little D", 3}
	_, ok1 := interface{}(&myDog).(Pet)
	_, ok2 := interface{}(myDog).(Pet)
	fmt.Printf("%v, %v\n", ok1, ok2)

	quit := make(chan string)

	go func(chan string) {
		quit <- "Bye!"
	}(quit)

	fmt.Print("Joe says:", <-quit)

}
