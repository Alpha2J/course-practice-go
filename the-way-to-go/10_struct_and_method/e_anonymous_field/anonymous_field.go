package main

import "fmt"

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b int
	c float32

	int
	innerS
}

// 字段没有显式的名字，只有字段的类型是必须的，此时类型就是字段的名字
func t1() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 5
	outer.in2 = 10

	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)

	// 使用结构体字面量
	outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
	fmt.Println("outer2 is:", outer2)
}

type A struct {
	ax, ay int
}

type B struct {
	A
	bx, by float32
}

// go 无继承，但可通过内嵌结构体模拟类似其他语言的继承行为
func t2() {
	b := B{A{1, 2}, 3.0, 4.0}
	fmt.Println(b.ax, b.ay, b.bx, b.by)
	fmt.Println(b.A)
}

func main() {
	//t1()
	t2()
}
