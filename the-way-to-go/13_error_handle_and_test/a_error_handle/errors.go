package main

import (
	"errors"
	"fmt"
)

var errNotFound = errors.New("not found error")

func t1() {
	fmt.Printf("error: %v", errNotFound)
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math - square root of negative number")
	}

	return f * f, nil
}

func t2() {
	if _, err := Sqrt(-1); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

func main() {
	//t1()
	t2()
}
