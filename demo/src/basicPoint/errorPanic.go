package main

import "os"

func main() {
	panic("a problem")

	_, err := os.Create("/t1mp/file")
	if err != nil {
		panic(err)
	}
}
