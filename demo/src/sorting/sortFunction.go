package main

import (
	"sort"
	"fmt"
)

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])//sort by len
	//return s[i] > s[j]//sort by desc
}

func main() {
	fruits := []string{"a", "bb", "ccc"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}