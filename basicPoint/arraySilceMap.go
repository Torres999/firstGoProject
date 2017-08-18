package main

import "fmt"

func main() {

	//arrayFunc()

	slicesFunc()

	//mapFunc()
}

/*
Array是在声明的时候都确定了长度，之后不可更改。Slice和数组类似，也是表示一个有序元素，但这个序列的长度可变。
*/
func arrayFunc() {
	var a [5]int
	fmt.Println("emp:", a)

	var a1 []int
	fmt.Println("emp:", a1)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

/*
区别：1.数组是值类型，切片是引用类型，在使用的时候有区别。
       如果用一个数组A创建新数组B的时候，需要将A中的所有元素全部复制一遍过去。如果是切片的话，就会指向原来的数据。
	 2.Array是在声明的时候都确定了长度，之后不可更改。Slice和数组类似，也是表示一个有序元素，但这个序列的长度可变。
*/
func slicesFunc() {
	//s := make([]string, 3)
	//fmt.Println("emp:", s)
	//fmt.Println("emp'len:", len(s))
	//
	//s[0] = "a"
	//s[1] = "b"
	//s[2] = "c"
	//fmt.Println("set:", s)
	//fmt.Println("get:", s[2])
	//
	//fmt.Println("len:", len(s))
	//
	//s = append(s, "d")
	//s = append(s, "e", "f")
	//fmt.Println("append:", s)
	//
	//c := make([]string, len(s))
	//copy(c, s)
	//fmt.Println("cpy:", c)
	//
	////这里的[low:high]，用数学集合的方式来讲，就是[low, high)，即左闭右开。
	////这种创建slice的方式，不需要make()函数。
	//l := s[2:5]
	//fmt.Println("s[2:5]:", l)
	//fmt.Println("地址比较，切片前value=c的元素地址:", &s[2]) //0xc420068080
	//fmt.Println("地址比较，切片后value=c的元素地址:", &l[0]) //0xc420068080
	//
	//l = s[:5]
	//fmt.Println("s[:5]:", l)
	//
	//l = s[2:]
	//fmt.Println("s[2:]:", l)
	//
	//t := []string{"g", "h", "i"}
	//fmt.Println("dcl:", t)
	//
	//twoD := make([][]int, 3)
	//for i := 0; i < 3; i++ {
	//	innerLen := i + 1
	//	twoD[i] = make([]int, innerLen)
	//	for j := 0; j < innerLen; j++ {
	//		twoD[i][j] = i + j
	//	}
	//}
	//fmt.Println("2d: ", twoD)

	// 切片地址可变
	twoC := make([]int, 3, 5)//make([]T, length, capacity)，初始值为3个0，如果是string的，同样会有三个空的初始化值
	twoC = append(twoC, 1, 2, 3)//超出容量后一倍扩容，切片扩容后会形成一个新的切片，具体参考：GoogleSearch2.go
	for _, v := range twoC {
		fmt.Printf("\n%d,", v)
	}
	fmt.Printf("\n新分配内存后切片容量%d", cap(twoC))
	fmt.Printf("\n新分配内存后切片长度%d", len(twoC))
}

func mapFunc() {
	m := make(map[string]int, 2)

	m["k1"] = 12
	m["k2"] = 34
	m["k3"] = 45
	fmt.Println("m[\"k3\"]:", m["k3"])
	fmt.Println("map:", m)
	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	//map 每次返回的数据有两部分组成：值 + 是否存在，如果用一个变量接收就是value
	//是否存在可以用来区分改key对应的value为“”或者0和改key不存在的情况
	_, prs := m["k2"]
	fmt.Println("prs:", prs) //是否存在
	qq := m["k2"]
	fmt.Println("qq:", qq) //值
	aa, _ := m["k2"]
	fmt.Println("aa:", aa) //值
	bb, cc := m["k2"]
	fmt.Println("value:", bb) //值
	fmt.Println("exist:", cc) //是否存在

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}
