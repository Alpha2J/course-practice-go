package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

type Shaper interface {
	Area() float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}

func t1() {
	// 1. 判断实际类型的语法叫做“类型断言”，但是只能用于接口，下面用结构体的写法会报错
	//var circle Circle = Circle{}
	//isC := circle.(Shaper)
}

func t2() {
	// 2. 类型断言可能是无效的，虽然编译器会尽力检查转换是否有效，但是不可能预见所有的可能性
	// 如果转换失败会有错误产生，如下
	var s Shaper
	// 像这样程序就直接 panic 了
	//isOk := s.(*Circle)
	//fmt.Println("isOk: ", isOk)

	// 所以建议使用这种方式来做检查，
	// 如果转换成功，v 是 s 转换到类型 Circle 的实际值，isOk 会是 true
	// 否则，s 是类型 Circle 的零值，isOk 是 false
	if v, isOk := s.(*Circle); isOk {
		fmt.Println("yes, it is Ok: ", isOk, v)
	} else {
		fmt.Println("no, it is not Ok: ", isOk, v)
	}
}

// 3. 完整的程序
func t3() {
	var areaIntf Shaper
	sq1 := new(Square)
	sq1.side = 5

	areaIntf = sq1
	// Is Square the type of areaIntf?
	if t, ok := areaIntf.(*Square); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}
	if u, ok := areaIntf.(*Circle); ok {
		fmt.Printf("The type of areaIntf is: %T\n", u)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}
}

func main() {
	//t1()
	//t2()
	t3()
}
