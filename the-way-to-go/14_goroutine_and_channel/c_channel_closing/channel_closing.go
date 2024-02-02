package main

import "fmt"

// todo
// 1. 通道可以被显示关闭，尽管和文件不同（不必每次都关闭），
// 只有当需要告诉接收者不会再提供新的值的时候，才需要关闭通道
// 只有发送者需要关闭通道，接收者不需要
func t1() {
	ch := make(chan string)
	go sendData(ch)
	getData(ch)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
	// 关闭通道
	close(ch)
}

func getData(ch chan string) {
	//for {
	//	input, open := <-ch
	//	if !open {
	//		break
	//	}
	//	fmt.Printf("%s ", input)
	//}

	// 这种方式和上面的一样，for...range...语句会自动检测通道是否关闭
	for input := range ch {
		fmt.Printf("%s ", input)
	}
}

func main() {
	t1()
}
