package main

import "fmt"

type Search interface {
	Search() string
}

type ViedoSearch struct {
	msg string
}

func (v ViedoSearch) Search() string {
	return "Search " + v.msg + " by ViedoSearch"
}

type WebSearch struct {
	msg string
}

func (v WebSearch) Search() string {
	return "Search " + v.msg + " by WebSearch"
}

type NewsSearch struct {
	msg string
}

func (v NewsSearch) Search() string {
	return "Search " + v.msg + " by NewsSearch"
}

type SearchTools []Search

func (v SearchTools) Search() []string {
	msg := make([]string, len(v))
	for j, i := range v {
		msg[j] = i.Search()
	}
	return msg
}

func main() {
	var msg = "PLA055"
	result := &SearchTools{
		ViedoSearch{msg},
		WebSearch{msg},
		NewsSearch{msg},
	}
	fmt.Printf("%#v\n", result.Search())
}
