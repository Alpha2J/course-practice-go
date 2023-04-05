package main

import (
	"fmt"
	"reflect"
)

// 知识点:
// 1. 可变长度的数组: 对数组一个连续片段的引用，这个片段可以是整个数组，也可以是由起始和终止索引标识的一些项的子集。
// 2. 切片是引用类型
// 3. 相关函数:
//   - len(): 获取切片长度的方式是
//   - cap(): 测量切片最长可以达到多少
// 4. 切片长度可在运行时修改：最小为0，最大为相关数组的长度

// 1. 声明格式为: var identifier []type
func t1() {

	var arr [5]int  // 这样是数组
	var slice []int // 这样是切片

	fmt.Println("isArr: ", isArray(arr))
	fmt.Println("isArr: ", isArray(slice))
}
func isArray(v interface{}) bool {
	t := reflect.TypeOf(v)
	return t.Kind() == reflect.Array
}

// 2. 初始化方式为:
//   - var slice []type = arr[start:end]
//   - var slice []type = arr[:]
func t2() {
	var arr [5]int = [5]int{0, 1, 2, 3, 4}
	fmt.Println("arr: ", arr)
	// 表示slice是由arr的第1到第3个元素组成
	var slice []int = arr[1:4]
	fmt.Println("slice: ", slice)

	var arr1 [5]int = [5]int{0, 1, 2, 3, 4}
	fmt.Println("arr1: ", arr1)
	// 表示slice是由完整的arr数组组成，相当于: arr[0:len(arr)]
	var slice1 []int = arr[:]
	fmt.Println("slice1: ", slice1)

	// todo: 好像有问题
	//var arr2 [5]int = [5]int{0, 1, 2, 3, 4}
	//fmt.Println("arr2: ", arr1)
	//var slice2 []int = &arr2
	//fmt.Println("slice2: ", slice2)

	var slice3 []int = []int{0, 1, 2, 3, 4}[:]
	fmt.Println("slice3: ", slice3)

	var slice4 []int = []int{0, 1, 2, 3, 4}
	fmt.Println("slice4: ", slice4)
}

// 3. 切片容量
//   - 当前长度
//   - 最大容量
//   - 缩容
//   - 扩容
func t3() {
	var arr [5]int = [5]int{0, 1, 2, 3, 4}
	var slice []int = arr[1:3]
	// len()拿到的是切片长度, cap()拿到的是切片可用的最大长度, 值为相关数组起始位置到len(arr)
	fmt.Println("len of slice: ", len(slice), " cap of slice: ", cap(slice))

	var slice1 = []int{0, 1, 2, 3, 4}
	fmt.Println("len of slice1: ", len(slice1), " cap of slice1: ", cap(slice1))
}

// 4. 切片缩扩容
func t4() {
	var arr [5]int = [5]int{0, 1, 2, 3, 4}
	var slice []int = arr[1:3]
	fmt.Println("len of slice: ", len(slice), " cap of slice: ", cap(slice))
	for i := 0; i < len(slice); i++ {
		fmt.Println("printing slice, i: ", i, " value: ", slice[i])
	}

	fmt.Println()

	// 缩容: 从slice的第一个元素开始, 一直切到最后一个元素, 然后复制给新的变量slice1
	// slice 和 slice1的底层数组共享, 只是它们自己的长度和容量回变化
	var slice1 []int = slice[1:]
	// 可以看到len变了, slice的len是2, 这里又减1, 所以这里len=1
	fmt.Println("after extension, len of slice1: ", len(slice1), " cap of slice1: ", cap(slice1))
	for i := 0; i < len(slice1); i++ {
		fmt.Println("printing slice1, i: ", i, " value: ", slice1[i])
	}

	fmt.Println()

	// 扩容: 可以将切片扩容到相关数组的终止下标位置, 可以反复扩展, 知道占领整个相关数组
	var slice2 []int = slice[:len(slice)+2]
	fmt.Println("len of slice2: ", len(slice2), " cap of slice2: ", cap(slice2))
	for i := 0; i < len(slice2); i++ {
		fmt.Println("printing slice2, i: ", i, " value: ", slice2[i])
	}
}

// 5. 切片元素赋值及使用
func t5() {
	var arr [5]int = [5]int{0, 1, 2, 3, 4}
	var slice []int = arr[1:3]
	for i, v := range slice {
		fmt.Println("i : ", i, ", v: ", v)
	}

	// 修改相关数组元素的数据会同步修改slice的数据
	arr[2] = 6
	fmt.Println("new slice: ", slice)

	// 修改slice会同步修改相关数组元素的数据
	slice[0] = 10
	fmt.Println("new arr: ", arr)
}

// 6. 将切片传递给函数: 数组是值类型, 切片是引用类型, 因此一般将切片传递给函数而不是数组, 这样做能避免大数组拷贝, 提高性能
func t6() {
	var arr [5]int = [5]int{0, 1, 2, 3, 4}
	fmt.Println("sum: ", sum(arr[:]))
}
func sum(arr []int) int {
	s := 0
	for _, i2 := range arr {
		s += i2
	}

	return s
}

// 7. 用make创建slice: 当相关数组还没有定义时, 可以使用 make() 函数来创建一个切, 同时创建好相关数组
func t7() {
	var slice []int = make([]int, 10)
	fmt.Println("t7(): slice: ", slice)
	fmt.Println("t7(): cap(slice): ", cap(slice), ",len(slice): ", len(slice))

	// 如果想创建一个slice, 但是这个slice不占用整个数组, 那么可以使用: make([]type, len, cap)
	var slice1 []int = make([]int, 5, 10)
	fmt.Println("t7(): slice1: ", slice1)
	fmt.Println("t7(): cap(slice1): ", cap(slice1), ",len(slice1): ", len(slice1))

}

func main() {
	//TSlice()
	//t2()
	//t3()
	t4()
	//t5()
	//t6()
	//t7()
}
