package main

import "fmt"

type List []int

func (l List) Len() int {
	return len(l)
}

func (l *List) Append(val int) {
	*l = append(*l, val)
}

type Appender interface {
	Append(int)
}

func CountInto(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

type Lener interface {
	Len() int
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}

// 1. 书本上的案例
func t1() {
	// A bare value
	var lst List

	// compiler error:
	// cannot use lst (type List) as type Appender in argument to CountInto:
	//       List does not implement Appender (Append method has pointer receiver)
	//CountInto(lst, 1, 10)
	if LongEnough(lst) { // VALID:Identical receiver type
		fmt.Printf("- lst is long enough\n")
	}

	// A pointer value
	plst := new(List)
	CountInto(plst, 1, 10) //VALID:Identical receiver type
	if LongEnough(plst) {
		// VALID: a *List can be dereferenced for the receiver
		fmt.Printf("- plst is long enough\n")
	}
}

// 2.
func t2() {
	// 2. 在接口上调用方法时，必须有和方法定义时相同的接收者类型或者是可以从具体类型

	// 1) 测试指针，结论: 方法接受者是指针的方法只可以通过指针调用，所以只能将实现者类型的指针赋值给接口变量
	//var l List
	//var lp *List
	//var appender Appender
	//appender = l
	//appender = lp

	// 2) 测试非指针，结论：方法接收者是值的方法可以通过指针调用，也可以通过值调用，无论是实现者类型的值还是指针都可以赋值给接口变量
	//var l List
	//var lp *List
	//var lener Lener
	//lener = l
	//lener = lp
	//fmt.Println(lener)
}

func main() {
	//t1()
	//t2()
}
