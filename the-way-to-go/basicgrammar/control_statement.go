package basicgrammar

import (
	"fmt"
	"math/rand"
)

func ifElseStatement() {
	if rand.Int()%2 == 1 {
		fmt.Println("executing in if statement")
	} else if rand.Int()%2 == 0 {
		fmt.Println("executing in else if statement")
	} else {
		fmt.Println("executing else statement")
	}
}

func switchStatement() {
	var var1 = 100
	switch var1 {
	case 1:
		fmt.Println("switch statement result is: ", 1)
	case 100:
		fmt.Println("switch statement result is: ", 100)
		//不需要使用break
		//break
	default:
		fmt.Println("switch statement result is: default")
	}
}

func selectStatement() {

}

func multiReturnValue() {

}

func forLoopStatement() {
	for i := 0; i < 2; i++ {
		fmt.Println("in for loop statement, i is: ", i)
	}
}

func forRangeStatement() {
	//	for pos, char := range str {
	//...
	//}
}

func init() {
	ifElseStatement()
	switchStatement()
	selectStatement()
	multiReturnValue()
	forLoopStatement()
}
