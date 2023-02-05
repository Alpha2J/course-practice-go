package basicgrammar

import "fmt"

func constVariable() {
	const Pi = 3.14159
	fmt.Println("Pi is: ", Pi)

	//const 不能重新赋值
	//Pi = 3.1415
}

func omitTypeDeclaring() {
	const str string = "abc"
	//类型可以省略，编译器自动推断
	const str1 = "efg"

	fmt.Println("str is: ", str)
	fmt.Println("str1 is: ", str1)
}

func omitKeywordDeclaring() {
	//这样写，可以省略关键字 'var'
	str := "omit keyword declaring"
	fmt.Println("omitted keyword declaring str is: ", str)
}

func skipVarDeclaring() {
	//_表示是一个只写变量，不能读，用于skip它的值
	_, b := 5, 7
	//fmt.Println("skip var declaring _ is: ", _)
	fmt.Println("skip var declaring b is: ", b)
}

func constAsEnum() {
	const (
		Unknown = 0
		Female  = 1
		Male    = 2
	)
}

func varTypeDeclaring() {
	var identifier int
	//声明变量的时候，没有初始化，使用默认值
	fmt.Println("identifier is: ", identifier)
}

func init() {
	fmt.Println("variable.go initialing...")
	constVariable()
	omitTypeDeclaring()
	varTypeDeclaring()
	omitKeywordDeclaring()
	fmt.Println("variable.go initialized end...")
}
