package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	defer closeFile1(f) //先于上一句执行
	writeFile(f)

	fmt.Println(fDefer())
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}
func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}
func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}
func closeFile1(f *os.File) {
	fmt.Println("closing1")
}

func fDefer() (result int) {
	fmt.Println("111")
	defer func() {
		result++
		fmt.Println("333")
	}()
	fmt.Println("222")
	return result
}
