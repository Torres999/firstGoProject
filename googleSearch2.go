package main

import (
	"fmt"
	"time"
)

var (
	Web   = fakeSearch("Web")
	Image = fakeSearch("Image")
	Vedio = fakeSearch("Vedio")
)

//type Result string
type Search func(msg string) string

func fakeSearch(kind string) Search {
	return func(msg string) string {
		return string(fmt.Sprintf("%s result for %q\n", kind, msg))
	}
}

/*
评估有会返回一个新的切片地址，取代原来的
In a function call, the function value and arguments are evaluated（评估） in the usual order. After they are evaluated,
the parameters of the call are passed by value to the function and the called function begins execution.
The return parameters of the function are passed by value back to the calling function when the function returns.
*/
func Google(msg string) []string {
	// 方式1
	//result := make([]string, 3)
	//result[0] = string(Web(msg))
	//result[1] = string(Image(msg))
	//result[2] = string(Vedio(msg))
	// 方式2
	result := make([]string, 0)
	result = append(result, Web(msg))
	result = append(result, Image(msg))
	result = append(result, Vedio(msg))
	//如果如下操作，报错：append(result, Web(msg)) evaluated but not used，原因见方法注释
	//append(result, Web(msg))
	return result
}

func main() {
	start := time.Now()

	result := Google("golang")

	elapsed := time.Since(start)

	fmt.Println(result)
	fmt.Println(elapsed)
}
