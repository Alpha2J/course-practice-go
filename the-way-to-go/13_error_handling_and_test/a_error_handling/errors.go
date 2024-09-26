package main

import (
	"errors"
	"fmt"
)

// 1. 创建新的错误类型
var errNotFound = errors.New("not found error")

func t1() {
	fmt.Printf("error: %v", errNotFound)
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		//return 0, errors.New("math - square root of negative number")
		return 0, fmt.Errorf("math: square root of negative number %g", f)
	}

	return f * f, nil
}

// 2. 调用函数，接收err并判断是否有err
func t2() {
	if _, err := Sqrt(-1); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

// 3. 使用fmt创建错误对象
func t3() {
	if _, err := Sqrt(-1); err != nil {
		fmt.Printf("Error: %s\n", err)
	}

}

// todo
//type CustomError struct {
//	msg  string
//	code int
//}
//
//func t3() {
//
//}
//
//func withCustomError() error {
//	return
//}

func main() {
	//t1()
	//t2()
	t3()
}
