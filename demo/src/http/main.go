package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type Data1 struct {
	CooperationNum  int `json:"cooperationNum"`
	ShelfProductNum int `json:"shelfProductNum"`
	FactorySellNum  int `json:"factorySellNum"`
}

type Body1 struct {
	Code int    `json:"code"`
	Info string    `json:"info"`
	Data Data1      `json:"data"`
}

func printResponse(resp *http.Response, err error) {
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("Status:", resp.Status)

	// 两种解析response的方式
	if resp.ContentLength < 0 {
		fmt.Println("resp.ContentLength:", resp.ContentLength)
		fmt.Println("resp.Body:", resp.Body)

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("err:", err)
		}

		fmt.Println("string(body):", string(body))//{"code":0,"info":"OK","data":{"cooperationNum":6,"shelfProductNum":894,"factorySellNum":2371}}

		num := Body1{}
		json.Unmarshal(body, &num)
		fmt.Println("shelfProductNum:", num.Data.ShelfProductNum)

	} else {
		fmt.Println("resp.ContentLength:", resp.ContentLength)
		fmt.Println("resp.Body:", resp.Body)

		byt := make([]byte, resp.ContentLength)
		resp.Body.Read(byt)
		//内置的toString方法
		fmt.Println("string(body):", string(byt))//{"code":5000,"info":"服务端异常","data":{}}

		var dat map[string]interface{}
		if err := json.Unmarshal(byt, &dat); err != nil {
			panic(err)
		}
		fmt.Println(dat)//map[code:5000 info:服务端异常 data:map[]]
	}
}

func main() {
	resp, err := http.Get("http://localhost:9006/product/getProductStatistic?factoryId=3")
	printResponse(resp, err)

	defer resp.Body.Close()

	fmt.Println("========")
	sendGetRequest()
}

func sendGetRequest() {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://localhost:9006/product/getProductStatistic?factoryId=3", nil)
	if err != nil {
		fmt.Println("err:", err)
	}
	req.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiJ9.eyJqdGkiOiJqd3QiLCJpYXQiOjE1MDE0ODk3MTksInN1YiI6IntcInVzZXJUeXBlXCI6MSxcImVudGVycHJpc2VJZFwiOjIsXCJ1c2VyTmFtZVwiOlwiMTg2NjI3MTA5OTZ1NGxJUE9tVWcyRDRNQTZOVzN4UFwifSIsImV4cCI6MTUwNDA4MTcxOX0.66t06IZ6u_EBqIlI_gq7p3jvvHjSuW_Kpcf-dDCzCGo")

	resp, err := client.Do(req)
	printResponse(resp, err)

	defer resp.Body.Close()
}