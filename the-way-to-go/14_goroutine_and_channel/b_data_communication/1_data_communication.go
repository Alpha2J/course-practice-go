package main

import (
	"fmt"
	"time"
)

// 通信的方式：
// 1. 变量共享：不提倡，会给所有的共享内存的多线程带来困难
// 2. channel：channel是go的一个特殊类型，可以通过它发送类型化的数据在协程之间通信。
// 数据通过通道，同一时间只有一个协程可以访问数据，不会出现数据竞争。可以读写数据的能力被传递。

func t1() {
	// 3. 声明通道的格式是: var identifier chan datatype
	// 一个通道只能传输一种类型的数据
	var ch1 chan string
	// 未初始化的通道的值是nil
	fmt.Printf("undefined ch1: %v", ch1)
	ch1 = make(chan string)
	// or:
	//ch1 := make(chan string)

	// 尽管不必要，但是通道的命名通常以ch开头或者包含chan
	// 函数通道
	//funcChan := chan func()
}

// 4. 使用通信操作符往通道传输数据或者从通道读取数据
func t2() {
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)

	time.Sleep(1e9)
}

func sendData(ch chan string) {
	// 通信操作符
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
}

func getData(ch chan string) {
	var input string

	//time.Sleep(2e9)

	for {
		input = <-ch
		fmt.Printf("%s ", input)
	}
}
