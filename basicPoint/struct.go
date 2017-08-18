package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	// 省略字段名
	fmt.Println(person{"jack", 14})

	// 省略部分字段
	fmt.Println(person{name: "jack1"}) //{jack1 0},int's default value is 0.
	fmt.Println(person{age: 14})

	// 全部字段名
	s := person{name: "Sean", age: 50}
	//s = &person{name: "Sean", age: 50}
	fmt.Println("Getter:", s.name)
	fmt.Println("Getter:", s.age)

	s.age = 22
	fmt.Println("Setter:", s.age)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

	a1 := new(person)
	var a2 *person = new(person)
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a2 == a1)
}
