package main

//
//import "fmt"
//
//type Stringer interface {
//	String() string
//}
//
//type StringerImplement struct {
//}
//
//func (si *StringerImplement) String() string {
//	return "hello world"
//}
//
//func t1() {
//	var i *StringerImplement = new(StringerImplement)
//
//	// 不行啊，这一章讲解得有点问题哦？
//	if si, ok := i.(Stringer); ok {
//		fmt.Printf("i implements String(): %T, %v", si, si)
//	} else {
//		fmt.Printf("i implements String(): %T, %v", si, si)
//	}
//}
//
//func main() {
//	t1()
//}
