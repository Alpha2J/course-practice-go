package arr

import "fmt"

// 基本语法
func t1() {
	//声明数组，用0值初始化
	var arr [5]int
	fmt.Println("arr: ", arr)

	//获取数组长度
	arrLen := len(arr)
	fmt.Println("arr len: ", arrLen)

	arr[0] = 2
	arr[1] = 8
	//超出数组长度，无法通过编译
	//arr[6] = 8
	fmt.Println("arr[0]: ", arr[0])
	fmt.Println("arr[1]: ", arr[1])

	//遍历数组并初始化
	for i := 0; i < len(arr); i++ {
		arr[i] = i + 5
	}
	fmt.Println("initialized arr: ", arr)

	//前者是index，后者是index对应的值
	for i, i2 := range arr {
		fmt.Println("range arr i=", i, ", i2=", i2)
	}
}

// 声明与赋值
func t2() {
	//声明变量并赋值
	var arr [5]int = [5]int{11, 22, 33, 44, 55}
	//arr1 := [5]int{11, 22, 33, 44, 55}
	fmt.Println("arr: ", arr)
}

// 数组是值类型
func t3() {
	arr := [5]int{2, 5, 6, 7, 8}
	fmt.Println("arr: ", arr)
	//因为数组是值类型，所以这里将arr的值拷贝了一份给arr1
	arr1 := arr
	arr1[1] = 31
	//这里arr打印出来的值并没有发生变化
	fmt.Println("after arr: ", arr)
	fmt.Println("after arr1: ", arr1)
}

func t4() {
	//new返回的是指针
	var arr1 *[5]int = new([5]int)
	var arr2 [5]int = *arr1
	fmt.Println(arr2)
}

func TArr() {
	t1()
	t2()
	t3()
}
