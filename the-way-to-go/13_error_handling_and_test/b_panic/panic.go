package main

import (
	"fmt"
	"os"
)

func t1() {
	fmt.Println("Starting the program")
	// 1. 产生panic后，程序会终止运行，不能随意地panic终止程序，必须尽力补救错误让程序能够继续执行
	panic("A severe error occurred: stopping the program!")
	fmt.Println("Ending the program")
}

var user = os.Getenv("USER")

func t2() {
	// 2. 所以想让程序立马停止的时候，可以制造一个panic
	if user == "" {
		panic("Unknown user: no value for $USER")
	}
}

// 3. 可以同时打印错误信息
func t3() {
	//if err != nil {
	//	panic("ERROR occurred:" + err.Error())
	//}
}

func main() {
	t1()
	//t2()
}
