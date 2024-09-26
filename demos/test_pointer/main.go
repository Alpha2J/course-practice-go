package main

import "fmt"

func main() {
	var a int = 1
	var aPointer *int = &a
	fmt.Println("before, a: ", a)
	*aPointer = 2
	fmt.Println("after, a: ", a)
}
