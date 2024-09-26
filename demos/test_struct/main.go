package main

import "fmt"
import person "cource-practice-go/demos/test_struct/diff_package"

func main() {
	t1()
	t2()
}

// 这说明了方法不只是可以加到struct上的
func t1() {
	var i aliasInt = 2
	i.modifyValue()
	fmt.Println("after func call modifyValue(), i is: ", i)
}

type aliasInt int

func (i *aliasInt) modifyValue() {
	fmt.Println("before modifying value, i is: ", *i)
	*i = 3
	fmt.Println("after modifying value, i is: ", *i)
}

func t2() {
	var person person.Person = person.Person{}
	// 下面这一行会报错，说明了结构体内的字段不是公开的，在其他包也无法访问
	// person.firstName
	fmt.Println("person: ", person)

	// 可以通过getter和setter方法来处理
	person.SetFirstName("jeb")
	fmt.Println("person: ", person.GetFirstName())
}
