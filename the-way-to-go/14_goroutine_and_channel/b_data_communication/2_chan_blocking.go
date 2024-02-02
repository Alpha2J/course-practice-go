package main

import (
	"fmt"
	"time"
)

// 1. 默认情况下，通信是同步且无缓冲的，在有接收者接收数据之前，发送不会结束。
func t3() {
	ch1 := make(chan int)
	go pump(ch1)
	time.Sleep(3e9)
	fmt.Println("begin to receive and print")
	fmt.Println(<-ch1)
	time.Sleep(1e9)
}

func pump(ch chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println("begin to pump: i=", i)
		ch <- i
		// 所以这里会等t3 sleep 3s之后，开始接收数据了，才会被打印出来
		fmt.Println("after pump: i=", i)
	}
}

// 2. 新的输入无法在通道为非空的情况下传入
func t4() {
	ch1 := make(chan int)
	go pumpA(ch1)
	go pumpB(ch1)
	time.Sleep(3e9)
	fmt.Println("begin to receive and print")
	fmt.Println(<-ch1)
	time.Sleep(1e9)
}

func pumpA(ch chan int) {
	fmt.Println("begin to pump in a")
	ch <- 0
	fmt.Println("after pump in a")
}

func pumpB(ch chan int) {
	fmt.Println("begin to pump in b")
	ch <- 1
	fmt.Println("after pump in b")
}

// 3. 如果通道是空的，那么接收者就会阻塞
func t5() {
	ch1 := make(chan int)
	go sleepAndPump(ch1)
	fmt.Println("begin to receive in t5")
	// 所以这个地方能看到会阻塞到sleepAndPump中sleep完并且pump完之后，才会继续往下走
	fmt.Println("received in t5: ", <-ch1)
}

func sleepAndPump(ch1 chan int) {
	fmt.Println("begin to sleep in sleepAndPump")
	time.Sleep(2e9)
	fmt.Println("after sleep in sleepAndPump and begin to pump")
	ch1 <- 0
	fmt.Println("after pump in sleepAndPump and after pump")
}
