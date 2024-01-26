package main

import (
	"fmt"
)

func t1() {
	innerPanic()
	fmt.Println("after innerPanic")
}

func innerPanic() {
	defer catchPanic()

	fmt.Println("start")
	panic("custom panic")
}

func catchPanic() {
	// 1. recover 函数被用于从panic或错误场景中回复，让程序可以从panicking重新获得控制权，停止终止进程而恢复正常执行
	// 只能用于defer修饰的函数
	// 2. 原则是，即使在包内使用了panic，在它的对外接口中必须用recover处理成返回显示的错误
	if err := recover(); err != nil {
		fmt.Printf("runtime panic: %v\n", err)
		return
	}

	fmt.Println("no panic")
}

func main() {
	t1()
}
