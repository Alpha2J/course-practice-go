package basicgrammar

import "fmt"

func basicType() {
	var b bool
	//整数
	var i int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	//无符号整数
	var ui uint
	var ui8 uint8
	var ui16 uint16
	var ui32 uint32
	var ui64 uint64

	var uiptr uintptr
	//没有float类型，只有float32和float64
	var f32 float32
	var f64 float64
	fmt.Println("b is: ", b)
	fmt.Println("i is: ", i)
	fmt.Println("i8 is: ", i8)
	fmt.Println("i16 is: ", i16)
	fmt.Println("i32 is: ", i32)
	fmt.Println("i64 is: ", i64)

	fmt.Println("ui is: ", ui)
	fmt.Println("ui8 is: ", ui8)
	fmt.Println("ui16 is: ", ui16)
	fmt.Println("ui32 is: ", ui32)
	fmt.Println("ui64 is: ", ui64)

	fmt.Println("uiptr is: ", uiptr)

	fmt.Println("f32 is: ", f32)
	fmt.Println("f64 is: ", f64)
}

func ifEquals() {
	var a bool
	var b bool = true
	fmt.Println("a equals b?: ", a == b)
}

func init() {
	basicType()
	ifEquals()
}
