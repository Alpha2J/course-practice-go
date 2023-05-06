package main

import "fmt"

type tagType struct {
	// with tag
	field1 bool "tag for field1"
}

// 1. 结构体中的字段除了有名字和类型外，还可以有一个可选的标签（tag）: 它是一个附属于字段的字符串，可以是文档或其他的重要标记
// 标签的内容不可以在一般的编程中使用，只有包 reflect 能获取它
func t1() {
	tType := &tagType{field1: true}
	fmt.Printf("tType: %v\n", tType)
}

func main() {
	t1()
}
