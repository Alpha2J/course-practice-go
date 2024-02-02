package main

import (
	"fmt"
	"time"
)

// 1. 可以在扩展的make命令中设置容量，来创建带缓冲的通道
func t6() {
	// buf指的是可以同时容纳的元素数量
	// 在没有超过容量的情况下，给通道发送数据是不会阻塞的，从通道接收数据也不会阻塞，直到缓冲为空
	buf := 10
	ch1 := make(chan int, buf)

	go pumpN(ch1, 5)
	// 如果把这行注释掉，可以看到pump阻塞在第11次发送数据的时候，因为缓冲已满
	//go pumpN(ch1, 11)
	time.Sleep(3e9)

	fmt.Println("----begin to suck")
	// 这里一行可以展示，当缓冲为空时，从通道接收数据会阻塞，阻塞在了 begin to suck: i= 5
	go suckN(ch1, 6)

	time.Sleep(3e9)
	fmt.Println("----after to suck")
}

func pumpN(ch chan int, n int) {
	for i := 0; i < n; i++ {
		fmt.Println("begin to pump: i=", i)
		ch <- i
		fmt.Println("after pump: i=", i)
	}
}

func suckN(ch chan int, n int) {
	for i := 0; i < n; i++ {
		fmt.Println("begin to suck: i=", i)
		fmt.Println(<-ch)
		fmt.Println("after suck: i=", i)
	}
}
