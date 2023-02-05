package slice

import "fmt"

// 基本语法
// 1. 可变长度的数组
// 2. 切片是引用类型
func t1() {
	var arr [6]int
	var slice []int = arr[2:5]
	fmt.Println("slice: ", slice)

	//修改slice底层数组会同步修改slice的数据
	arr[3] = 6
	fmt.Println("slice: ", slice)

	fmt.Println("len of arr: ", len(arr))
	fmt.Println("len of slice: ", len(slice))
	fmt.Println("cap of arr: ", cap(slice))

}

// 用make创建slice：当相关数组还没有定义时，可以使用 make () 函数来创建一个切片 同时创建好相关数组
func t2() {
	var slice []int = make([]int, 10)
	fmt.Println("t2(): slice: ", slice)
	fmt.Println("t2(): cap(slice): ", cap(slice))
	fmt.Println("t2(): len(slice): ", len(slice))
}

func TSlice() {
	t1()
	t2()
}
