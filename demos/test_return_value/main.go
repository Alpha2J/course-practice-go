package main

import "fmt"

type mystruct struct {
	name string
}

func returnStruct() mystruct {
	var ms mystruct = mystruct{name: "jeb"}

	fmt.Printf("before return, addr is: %p\n", &ms)
	return ms
}

func returnPointer() *mystruct {
	var ms *mystruct = &mystruct{name: "jeb1"}
	fmt.Printf("before return, addr is: %p\n", ms)
	return ms
}

func main() {
	// 可以看到，当返回结构体，而不是结构体指针时，地址的值发生了变化，说明返回之前执行了复制操作
	var ms mystruct = returnStruct()
	fmt.Printf("after return, addr is: %p\n", &ms)

	var ms1 *mystruct = returnPointer()
	fmt.Printf("after return, addr1 is: %p\n", ms1)
}
