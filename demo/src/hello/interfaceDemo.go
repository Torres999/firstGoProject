package main

import (
	"math"
	"fmt"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

/*
结构体实现某一个interface，实现方法的参数可以使用指针，区别与struct中定义自己的方式时可以是用指针一样
 */
func (r *rect) area() float64 {
	return r.width * r.height
}

func (r *rect) perim() float64 {
	return 2 * r.width + 2 * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

/**
	一个接口有多个实现且可能返回不同的实现类对象的时候可以返回接口对象，然后再通过"interfaceObject.(*具体实现类)"的方式转换成具体的实现类对象
	具体代码参考errorsInterface.go
 */
/*
	两种实例化接口的方式：
	var interfaceObj1 testInterface  = &testStruct1{"123asdf"}
	var interfaceObj2 testInterface  = new(testStruct2)
	interfaceObj2.(*testStruct2).str = "2341234"
	具体代码参考errorsInterface.go
 */
func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	measure(&r)
	measure(c)
}