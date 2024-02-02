package main

import (
	"fmt"
	"time"
)

func t1() {
	fmt.Println("In t1()")
	go longWait()
	go shortWait()
	fmt.Println("About to sleep in t1()")
	// 1. 这里暂停10秒，会在longWait结束之后才结束
	// 如果暂停4秒，小于longWait()的暂停时间，那么longWait()提前结束，
	// 因为当 main() 函数返回的时候，程序退出：它不会等待任何其他非 main 协程的结束
	time.Sleep(10 * 1e9)
	fmt.Println("At the end of t1()")
}

func longWait() {
	fmt.Println("Beginning longWait()")
	time.Sleep(5 * 1e9)
	fmt.Println("End of longWait()")
}

func shortWait() {
	fmt.Println("Beginning shortWait()")
	time.Sleep(2 * 1e9)
	fmt.Println("End of shortWait()")
}

func main() {
	t1()
}
