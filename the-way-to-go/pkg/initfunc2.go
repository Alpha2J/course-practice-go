package pkg

import "fmt"

// 每个源文件都可以有init函数，在main函数之前执行
func init() {
	fmt.Println("init func in initfunc2.go")
}
