package basicgrammar

import "fmt"

func accessByPointer() {
	var i = 5
	fmt.Println("address of i is: ", &i)
	fmt.Println("value of i is: ", i)
}

func modifyByPointer() {
	var str = "original string"
	fmt.Println("original str is: ", str)

	var strPointer *string = &str
	// 指针的类型必须和指针指向的变量的类型一样
	//var strPointer *int = &str
	//var strPointer = &str
	fmt.Println("address or str is: ", strPointer)

	//string类型也是值类型，所以修改后
	str = "new original string"
	//这里打印出的值是修改后的值
	fmt.Println("modified original str is: ", str)
	//地址还是没变
	fmt.Println("modified address or str is: ", strPointer)
	//通过在指针类型前面加上*来访问指针的值
	fmt.Println("access str by pointer: ", *strPointer)
	//通过指针修改变量的值
	*strPointer = "modified string by pointer"
	//可以看到原始变量的值被修改了
	fmt.Println("modified string by pointer, access by original: ", str)
	fmt.Println("modified string by pointer, access by pointer: ", *strPointer)

}

func init() {
	accessByPointer()
	modifyByPointer()
}
