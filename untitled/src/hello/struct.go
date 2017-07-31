package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"jack", 14})

	fmt.Println(person{name:"jack1"})//{jack1 0},int's default value is 0.

	fmt.Println(person{age:14})

	s := person{name: "Sean", age: 50}
	fmt.Println("Getter:", s.name)
	fmt.Println("Getter:", s.age)

	s.age = 22
	fmt.Println("Setter:", s.age)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
