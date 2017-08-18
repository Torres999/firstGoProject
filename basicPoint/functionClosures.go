package main

import (
	"fmt"
	"sort"
)

func intSeq(a int) func() int {
	if a == 1 {
		return func() int {
			return a * 2
		}
	} else {
		return func() int {
			return a * 10
		}
	}
}

func intSeq1() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

/**
闭包里的非传递参数外部变量值是传引用的，在闭包函数里那个i就是外部非闭包函数自己的参数，所以是相当于引用了外部的变量
*/
func main() {
	nextInt0 := intSeq1()
	fmt.Println(nextInt0())
	fmt.Println(nextInt0())
	fmt.Println(nextInt0())

	nextInt1 := intSeq(1)
	nextInt2 := intSeq(2)

	fmt.Println("nextInt1:", nextInt1())
	fmt.Println("nextInt2:", nextInt2())
	fmt.Println("----------")

	p1 := Player{name:"1", level:4}
	p2 := Player{name:"2", level:3}
	p3 := Player{name:"3", level:2}
	p4 := Player{name:"4", level:1}

	t1 := Team{p1, p2, p3, p4}

	t1.SortByName()
	fmt.Println("SortByName:", t1)
	t1.SortByLevel()
	fmt.Println("SortByLevel:", t1)

	s1 := sortLevel(t1)
	fmt.Println("len:", s1.Len())

}

type Player struct {
	name  string
	level int
}
type Team []Player

func (t *Team) Join(p Player) {
	// 实现同上
}
func (t *Team) Quit(p Player) {
	// 实现同上
}
func (t *Team) SortByLevel() {
	// 根据级别排序
	sort.Sort(sortLevel(*t))
}
func (t *Team) SortByName() {
	// 根据名字排序
	sort.Sort(sortName(*t))
}

// 级别排序
type sortLevel Team

func (t sortLevel) Len() int {
	return len(t)
}
func (t sortLevel) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
func (t sortLevel) Less(i, j int) bool {
	return t[i].level < t[j].level
}

// 名字排序
type sortName Team

func (t sortName) Len() int {
	return len(t)
}
func (t sortName) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
func (t sortName) Less(i, j int) bool {
	return t[i].name < t[j].name
}
