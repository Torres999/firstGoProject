package main

import "fmt"

func main() {
	nums := []int{123, 12, 54, 56, 674}

	for i, num := range nums {
		fmt.Println("index:", i)
		fmt.Println("value:", num)
	}

	fmt.Println("=====================")

	for i, num := range nums {
		if i == 3 {
			fmt.Println("index:", num)
		}
	}

	fmt.Println("=====================")

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v) // string替换符
	}

	fmt.Println("=====================")

	for k := range kvs {
		fmt.Println("key:", k)
	}

	fmt.Println("=====================")

	/*
		range on strings iterates over Unicode code points. The first value is the starting byte index of the rune and the second the rune itself.
	*/
	for i, c := range "go1" {
		fmt.Println(i, c)
	}

	fmt.Println("=====================")

	/*
		Multiple Return Values
	*/
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}

// Multiple Return Values
func vals() (int, int) {
	return 3, 7
}
