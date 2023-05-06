package main

import "fmt"

type File struct {
	fd   int
	name string
}

// 1. 按惯例，工厂的名字以 new 或 New 开头
func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}

	return &File{fd: fd, name: name}
}

func t1() {
	myFile := NewFile(1, "my file")
	fmt.Printf("myFile: %v\n", myFile)
}

// 2. 强制使用工厂方法 todo
func t2() {
	//https://learnku.com/docs/the-way-to-go/102-creates-an-instance-of-a-structure-using-a-factory-method/3639
}

// 3. map和struct vs new()和make()
// https://learnku.com/docs/the-way-to-go/72-slice/3613#9b53a0
//试图 make() 一个结构体变量，会引发一个编译错误，这还不是太糟糕，
//但是 new() 一个映射并试图使用数据填充它，将会引发运行时错误！
//因为 new(Foo) 返回的是一个指向 nil 的指针，它尚未被分配内存。所以在使用 map 时要特别谨慎。
func t3() {
	//package main
	//
	//type Foo map[string]string
	//type Bar struct {
	//    thingOne string
	//    thingTwo int
	//}
	//
	//func main() {
	//    // OK
	//    y := new(Bar)
	//    (*y).thingOne = "hello"
	//    (*y).thingTwo = 1
	//
	//    // NOT OK
	//    z := make(Bar) // 编译错误：cannot make type Bar
	//    (*z).thingOne = "hello"
	//    (*z).thingTwo = 1
	//
	//    // OK
	//    x := make(Foo)
	//    x["x"] = "goodbye"
	//    x["y"] = "world"
	//
	//    // NOT OK
	//    u := new(Foo)
	//    (*u)["x"] = "goodbye" // 运行时错误!! panic: assignment to entry in nil map
	//    (*u)["y"] = "world"
	//}
}

func main() {

}
