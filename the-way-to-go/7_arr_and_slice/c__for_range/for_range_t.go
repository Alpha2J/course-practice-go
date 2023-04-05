package main

import "fmt"

// 知识点:
// 1. for-range结构可用于数组和切片
// 2. 第一个参数表示索引, 第二个参数表示索引对应的值
// 3. 如果只需要索引, 可以忽略第二个变量
// 4. 如果只需要索引对应的值, 不需要索引, 可以用_来表示第一个变量

func t1() {
	var slice1 []int = make([]int, 4)

	slice1[0] = 1
	slice1[1] = 2
	slice1[2] = 3
	slice1[3] = 4

	for ix, value := range slice1 {
		fmt.Printf("Slice at %d is: %d\n", ix, value)
	}
}

// 如果只需要索引, 可以忽略第二个变量
func t2() {
	var seasons []string = []string{"Spring", "Summer", "Autumn", "Winter"}

	for ix := range seasons {
		fmt.Printf("%d", ix)
	}
}

// 如果只需要索引对应的值, 不需要索引, 可以用_来表示第一个变量
func t3() {
	var seasons []string = []string{"Spring", "Summer", "Autumn", "Winter"}

	for _, value := range seasons {
		fmt.Printf("%s ", value)
	}
}

// item 是元素的拷贝
func t4() {
	// 如下代码并不会让items数组内的每个元素的值double
	// ... 表示可变长度
	items := [...]int{10, 20, 30, 40, 50}
	for _, item := range items {
		item *= 2
	}
	fmt.Printf("items: %d", items)

	fmt.Println("")

	// 如果要让items内每个元素的值double, 需要改写成这样
	for i := range items {
		items[i] *= 2
	}
	fmt.Printf("items: %d", items)
}

func main() {
	//t1()
	//t2()
	//t3()
	t4()
}
