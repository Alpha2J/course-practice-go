package basicgrammar

import (
	"fmt"
	"strings"
)

func utilStrings() {
	var str string = "this is a new string"
	// go中用strings包来完成对字符串的操作
	fmt.Println("has prefix 'this'", strings.HasPrefix(str, "this"))
}

func init() {
	utilStrings()
}
