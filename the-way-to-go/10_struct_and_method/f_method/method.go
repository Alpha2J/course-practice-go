package main

import "fmt"

// 1. go方法是作用在接收者上的一个函数，接收者是某种类型的变量
// 2. 接收者类型可以是任何类型，不仅仅是结构体类型
// 3. 类型加上它的方法等价于面向对象中的类，区别是，go中，类型的代码
// 和绑定在它上面的方法的代码可以不放置在一起，它们可以存在于不同的源文件，
// 唯一的要求是：它们必须是同一个包的
// 4. 方法不允许重载，即对于一个类型，只能有一个给定名称的方法
// 5. go中没有类似self或者this的关键字

type TwoInts struct {
	a int
	b int
}

func (tn *TwoInts) AddThem() int {
	return tn.a + tn.b
}
func (tn *TwoInts) AddToParam(param int) int {
	return tn.a + tn.b + param
}

// 方法于结构体上
func t1() {
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 22

	fmt.Printf("The sum is: %d\n", two1.AddThem())
	fmt.Printf("Add them to the param: %d\n", two1.AddToParam(20))

	two2 := TwoInts{3, 4}
	fmt.Printf("The sum is: %d\n", two2.AddThem())
}

type IntVector []int

func (v IntVector) Sum() (s int) {
	for _, x := range v {
		s += x
	}
	return
}

func t2() {
	fmt.Println(IntVector{1, 2, 3}.Sum())
}

type B struct {
	thing int
}

func (b *B) change() {
	b.thing = 1
}

func (b B) write() {
	b.thing = 1
}

// 指针作为接收者：可改变接收者的数据
// 值作为接收者：不可改变接收者的数据
func t3() {
	var b B
	b.change()
	fmt.Printf("b: %v", b)

	var b1 B
	b1.write()
	fmt.Printf("b1: %v", b1)
}

func main() {
	//t1()
	//t2()
	t3()
}
