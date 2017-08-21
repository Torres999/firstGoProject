package main

import (
	"fmt"
)

func main() {
	//panic("a problem")
	//
	//_, err := os.Create("/t1mp/file")
	//if err != nil {
	//	panic(err)
	//}


	// 必须要先声明defer，否则不能捕获到panic异常
	defer func() {
		fmt.Println("c")
		//panic一般会导致程序挂掉（除非recover）
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("d")
		f123(1)
	}()
	f123(2)
}

func f123(num int) {
	fmt.Println("a")
	if num > 1 {
		panic(nil)
	}
	fmt.Println("b")
	fmt.Println("f")
}
