package main

import (
	"regexp"
	"fmt"
	"bytes"
)

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r.MatchString("peach"))
	fmt.Println(r.FindString("peach punch"))
	fmt.Println(r.FindStringIndex("peach punch"))//[0 5]
	fmt.Println(r.FindStringSubmatch("peach punch"))////?????
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))
	fmt.Println(r.FindAllString("peach punch pinch", -1))
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	fmt.Println("1-----------")

	fmt.Println(r.FindAllString("peach punch pinch", 2))
	fmt.Println(r.Match([]byte("peach")))

	fmt.Println("2-----------")

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))


	fmt.Println("3-----------")
	a, _ := regexp.Compile("[1,9]{1,3}(\\.[1,9]{1,3}){3}")
	fmt.Println(a.MatchString("192.168.12.1"))

	fmt.Println("4-----------")
	b, _ := regexp.Compile("\\d{1,3}(\\.\\d{1,3}){3}")
	fmt.Println(b.MatchString("123.12.12.12"))
}