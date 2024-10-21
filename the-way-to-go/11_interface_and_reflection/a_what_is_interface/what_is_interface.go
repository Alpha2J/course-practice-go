package main

import (
	"fmt"
	"io"
	"math"
)

// （按照约定，只包含一个方法的）接口的名字由方法名加 [e]r 后缀组成，例如 Printer、Reader、Writer、Logger、Converter 等等。
// 还有一些不常用的方式（当后缀 er 不合适时），比如 Recoverable，此时接口名以 able 结尾，或者以 I 开头（像 .NET 或 Java 中那样）。

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Circle struct {
	radius float32
}

func (c *Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}

// 1. 实现了接口方法的类型自动实现了接口
// 2. 接口变量接收的是类型的指针
func t1() {
	var square Square = Square{side: 22.2}
	fmt.Println("square: ", square)

	var shaper Shaper = &square
	fmt.Println("shaper.Area(): ", shaper.Area())

	var circle *Circle = &Circle{radius: 4}
	fmt.Println("circle: ", circle)

	shaper = circle
	fmt.Println("shaper.Area(): ", shaper.Area())
}

type Interface1 interface {
	Method1()
	Method2()
}

type Implement1 struct {
}

func (i1 *Implement1) Method1() {
	fmt.Println("this is Method1")
}

// 3. 类型要实现接口中的所有方法，才能是这个接口的实现类型
func t2() {
	//var i1 *Implement1 = &Implement1{}
	// 这里只实现了Method1, 所以会报错
	//var it1 Interface1 = i1
	//fmt.Println("Interface1.Method1(): ", it1.Method1)
}

type MyReader struct {
}

func (mr *MyReader) Read(p []byte) (n int, err error) {
	fmt.Println("this is method from MyReader")
	return 0, err
}

// 4. 可以实现不同包下的接口
func t3() {
	var myReader io.Reader = &MyReader{}
	var ok, _ = myReader.Read(make([]byte, 10))
	fmt.Println("read ok? : ", ok)
}

// 5. 接口类型变量的初始值是 nil
func t4() {
	var i1 Interface1
	fmt.Println("is i1 equals nil: ", i1 == nil)
}

func main() {
	//t1()
	//t2()
	//t3()
	t4()
}
