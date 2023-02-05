package basicgrammar

import "fmt"

func alias() {
	//type关键字定义类型别名
	type TZ int
	var varByTz TZ = 100
	fmt.Println("var_by_TZ is: ", varByTz)
}

func init() {
	alias()
}
