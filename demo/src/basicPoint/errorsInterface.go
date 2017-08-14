package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

/*
It’s possible to use custom types as errors by implementing the Error()
method on them. Here’s a variant on the example above that uses a custom
type to explicitly represent an argument error.
*/
//@Overwrite
func (e *argError) Error() string {
	return fmt.Sprintf("%d ==== %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	//for _, i := range []int{7, 42} {//_:index  i:value
	//	if r, e := f1(i); e != nil {
	//		fmt.Println("f1 failed:", e)
	//	} else {
	//		fmt.Println("f1 worked:", r)
	//	}
	//}

	//for _, i := range []int{7, 42} {
	//	if r, e := f2(i); e != nil {
	//		fmt.Println("f2 failed:", e)
	//	} else {
	//		fmt.Println("f2 worked:", r)
	//	}
	//}

	fmt.Println("1=================")
	_, e := f2(42)
	if ae, ok := e.(*argError); ok { // e.(*argError)：Error有多个实现的时候，指定将Error转换成ergError对象
		fmt.Println(&argError{42, "can't work with it"})
		fmt.Println(e)
		fmt.Println("ae:", ae)
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
		fmt.Println(ok)
	}

	fmt.Println("2=================")

	////方法一：
	////采用errors包的New方法 返回一个err的类型
	//var err error = errors.New("this is a new error")
	////由于已经实现了error接口的方法 因此可以直接调用对应的方法
	//fmt.Println(err.Error())
	//
	////方法二：
	////采用fmt.Errof 将string信息转化为error信息 并返回
	//err = fmt.Errorf("%s", "the error test for fmt.Errorf")
	//fmt.Println(err.Error())
	//
	////方法三：
	////采用自定义的方式实现一个error的 一个duck 类型
	//err = &Customerror{
	//	infoa: "err info a",
	//	infob: "err info b",
	//	Err:   errors.New("test custom err"),
	//}
	//fmt.Println(err.Error())

	//========================================
	// 一个接口有多个实现，用b.(*testStruct2)的方式将接口转换成testStruct2对象;如果不转换的话只能调用接口中的Dalin()方法
	//========================================
	fmt.Println("3=========")
	a, b := ff(42)
	fmt.Println("a:", a)
	c, d := b.(*testStruct2)
	fmt.Println("c:", c.str)
	fmt.Println("c:", c.Dalin())
	fmt.Println("d:", d)

	fmt.Println("4=========")
	var interfaceObj1 testInterface = &testStruct1{"123asdf"}
	var interfaceObj2 testInterface = new(testStruct2)
	interfaceObj2.(*testStruct2).str = "2341234"
	fmt.Println("interfaceObj1:", interfaceObj1)
	fmt.Println("interfaceObj2:", interfaceObj2)
}

//========================================
// 一个接口有多个实现，用b.(*testStruct2)的方式将接口转换成testStruct2对象;如果不转换的话只能调用接口中的Dalin()方法
//========================================
func ff(arg int) (int, testInterface) { //返回的是接口对象，但是具体的实现是实现类对象
	if arg == 42 {
		return -1, &testStruct2{"asfdasgf"}
	}
	return arg + 3, nil
}

type testStruct1 struct {
	str string
}

func (t testStruct1) Dalin() string {
	return "Hello Dalin111"
}

type testStruct2 struct {
	str string
}

func (t testStruct2) Dalin() string {
	return "Hello Dalin8888"
}

type testInterface interface {
	//Dalin() string
	Dalin() (str string)
}

//========================================

type Customerror struct {
	infoa string
	infob string
	Err   error
}

func (cerr Customerror) Error() string {
	errorinfo := fmt.Sprintf("infoa : %s , infob : %s , original err info : %s ", cerr.infoa, cerr.infob, cerr.Err.Error())
	return errorinfo
}
