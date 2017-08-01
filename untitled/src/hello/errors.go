package main

import (
	"fmt"
	"errors"
)

//func f1(arg int) (int, error) {
//	if arg == 42 {
//		return -1, errors.New("can't work with 42")
//	}
//	return arg + 3, nil
//}




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
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
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

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	fmt.Println(e)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
		fmt.Println(ok)
	}

	fmt.Println("=================")

	//方法一：
	//采用errors包的New方法 返回一个err的类型
	var err error = errors.New("this is a new error")
	//由于已经实现了error接口的方法 因此可以直接调用对应的方法
	fmt.Println(err.Error())

	//方法二：
	//采用fmt.Errof 将string信息转化为error信息 并返回
	err = fmt.Errorf("%s", "the error test for fmt.Errorf")
	fmt.Println(err.Error())

	//方法三：
	//采用自定义的方式实现一个error的 一个duck 类型
	err = &Customerror{
		infoa: "err info a",
		infob: "err info b",
		Err:   errors.New("test custom err"),
	}
	fmt.Println(err.Error())
}


type Customerror struct {
	infoa string
	infob string
	Err   error
}

func (cerr Customerror) Error() string {
	errorinfo := fmt.Sprintf("infoa : %s , infob : %s , original err info : %s ", cerr.infoa, cerr.infob, cerr.Err.Error())
	return errorinfo
}