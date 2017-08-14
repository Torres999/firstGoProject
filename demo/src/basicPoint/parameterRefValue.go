package main

import "fmt"

/**
类似java中的值传递、引用传递
*/
func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i) //The &i syntax gives the memory address of i
	fmt.Println("zeroptr:", i)
	fmt.Println("pointer:", &i)
}

/**
zeroptr in contrast has an *int parameter, meaning that it takes an int pointer.
The *iptr code in the function body then dereferences the pointer from its memory
address to the current value at that address. Assigning a value to a dereferenced
pointer changes the value at the referenced address.
*/
func zeroptr(iptr *int) {
	fmt.Println("========")
	*iptr = 0
}

func zeroval(ival int) {
	fmt.Println("--------")
	ival = 0
}
