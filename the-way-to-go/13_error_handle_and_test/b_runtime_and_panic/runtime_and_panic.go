package main

import (
	"fmt"
	"os"
)

// 1. panic是什么
// 2. 如何产生panic

func t1() {
	fmt.Println("Starting the program")
	panic("A severe error occurred: stopping the program!")
	fmt.Println("Ending the program")
}

var user = os.Getenv("USER")

func t2() {
	if user == "" {
		panic("Unknown user: no value for $USER")
	}
}

func t3() {
	//if err != nil {
	//	panic("ERROR occurred:" + err.Error())
	//}
}

func main() {
	//t1()
	t2()
}
