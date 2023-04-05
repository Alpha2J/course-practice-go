package main

import "fmt"

// 知识点：如果想增加切片的容量, 我们必须创建一个更大的切片并把原切片的内容都拷贝过来

//  1. copy 方法将类型为 T 的切片从源地址 src 拷贝到目标地址 dst，覆盖 dst 的相关元素，并且返回拷贝的元素个数
//     拷贝个数是 src 和 dst 的长度(len)最小值: 意思是, 只会将min(len(src, dst))个元素进行拷贝, 可以打开tagA处的注释查看效果
func t1() {
	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 2)
	//sl_to := make([]int, 10)

	n := copy(sl_to, sl_from)
	fmt.Println(sl_to)
	fmt.Printf("Copied %d elements\n", n) // n == 3
}

// 2. append 将0个或多个具有相同类型的元素添加到切片后面并返回新的切片
func t2() {
	slice := []int{1, 2, 3}
	fmt.Println("before append, len(slice): ", len(slice), " cap(slice): ", cap(slice))

	slice = append(slice, 4, 5, 6)
	fmt.Println("after append, len(slice): ", len(slice), " cap(slice): ", cap(slice), " slice: ", slice)
}

//  3. 如果相关数组的容量不足以存储新的元素, append会分配新的切片来保证已有切片元素和新增元素的存储
//     因此, 返回的切片可能已经指向一个不同的相关数组了
func t3() {
	arr := []int{1, 2, 3}
	slice := arr[0:2]
	fmt.Println("before append, len(slice): ", len(slice), " cap(slice): ", cap(slice))

	slice = append(slice, 4, 5, 6)
	fmt.Println("after append, len(slice): ", len(slice), " cap(slice): ", cap(slice), " slice: ", slice)

	// 相关数组已经不是之前的那个了
	arr[0] = 111
	fmt.Println("after change arr, slice: ", slice)
}

func main() {
	t1()
	//t2()
	//t3()
}
