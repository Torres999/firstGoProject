package main

import (
	"encoding/xml"
	"fmt"
)

type person struct {
	Name string
	//如果需要把Name设置成person节点的属性，如下声明：
	//Name string `xml:"name,attr"`
	Age int
}

func main() {
	p := person{Name: "Jack", Age: 22}

	var data []byte
	var err error

	//普通序列化
	if data, err = xml.Marshal(p); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("data：", data)             //data： [60 112 101 114 115 111 110 62 60 78 97 109 101 62 108 105 108 105 110 108 105 110 60 47 78 97 109 101 62 60 65 103 101 62 50 50 60 47 65 103 101 62 60 47 112 101 114 115 111 110 62]
	fmt.Println("Formater：", string(data)) //Formater： <person><Name>Jack</Name><Age>22</Age></person>

	/*
		//格式化加前缀序列化
		if data, err = xml.MarshalIndent(p, "", " "); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("data：", data)
		fmt.Println("Formater：")
		fmt.Println(string(data))
	*/
	fmt.Println("========================")
	p2 := new(person)
	//反序列化
	if err = xml.Unmarshal(data, p2); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("p2：", p2) //p2： &{Jack 22}    -- 指针
}
