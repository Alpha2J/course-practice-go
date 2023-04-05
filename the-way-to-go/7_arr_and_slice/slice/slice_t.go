package main

import "fmt"

// 1. 可变长度的数组: 对数组一个连续片段的引用，这个片段可以是整个数组，也可以是由起始和终止索引标识的一些项的子集。
// 2. 切片是引用类型
// 3. 相关函数：
//   - len(): 获取切片长度的方式是
//   - cap(): 测量切片最长可以达到多少
// 4. 切片长度可在运行时修改：最小为0，最大为相关数组的长度

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

func main() {
	TSlice()
}
