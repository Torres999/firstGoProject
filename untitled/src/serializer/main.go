package main

import (
	"encoding/xml"
	"fmt"
)

type person struct {
	Name string
	Age  int
}

func main() {
	p := person{Name:"lilinlin", Age:22}

	//普通反序列化
	if data, err := xml.Marshal(p); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("data：", data)
		fmt.Println("Formater：", string(data))
	}

	//格式化加前缀反序列化
	if data, err := xml.MarshalIndent(p, "", " "); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("data：", data)
		fmt.Println("Formater：")
		fmt.Println(string(data))
	}
}