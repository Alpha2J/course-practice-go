package main

import (
	"cource-practice-go/the-way-to-go/10_struct_and_method/a_struct_definition/struct_sub_pkg"
	"fmt"
)

// 结构体在面向对象的编程语言中，跟一个无方法的轻量级类一样。
// 不过因为 Go 语言中没有类的概念，因此在 Go 中结构体有着更为重要的地位。

// 1. 定义结构体
type struct1 struct {
	i1   int
	f1   float32
	str1 string
	//    如果字段在代码中从来也不会被用到，那么可以命名它为 _
	_ int
}

// 2. 实例化对象、赋值并读取字段值
func t1() {
	// 初始化结构体实例（对象）：使用 new 函数给一个新的结构体变量分配内存，它返回指向已分配内存的指针
	var sct *struct1 = new(struct1)
	// 这样也会分配内存，并零值化内存
	var sct1 struct1
	// 这样写无法通过编译，会报错，因为new返回的是新分配的内存的指针，应该用指针变量来接收返回值
	//var sct2 struct1 = new(struct1)

	// 赋值：直接用点号可以为字段赋值
	fmt.Printf("sct: before set value: %v\n", sct)
	sct.i1 = 222
	sct.f1 = 3.33
	sct.str1 = "hello sct"
	fmt.Printf("sct: after set value: %v\n", sct)

	fmt.Printf("sct1: before set value: %v\n", sct1)
	sct1.i1 = 103
	sct1.f1 = 1.22
	sct1.str1 = "hello sct1"
	fmt.Printf("sct1: after set value: %v\n", sct1)

	// 读取字段值：直接用点号可以读取字段值
	// 无论变量是结构体类型还是结构体类型指针，都可以使用点号来处理字段的赋值与读取
	fmt.Printf("sct: value of field i1: %d\n", sct.i1)
	fmt.Printf("sct1: value of field i1: %d\n", sct1.i1)
}

type struct2 struct {
	i2   int
	f2   float32
	str2 string
}

// 3. 使用结构体字面量实例化对象
func t2() {
	// 底层仍然是调用new()
	// 全部字段：可以省略字段名
	var literalSct *struct2 = &struct2{10, 12, "literalSct"}
	// 部分字段：字面量实例化时需要加上字段名
	var literalSct1 *struct2 = &struct2{i2: 10, str2: "literalSct1"}
	fmt.Printf("literalSct: after instantiate: %v\n", literalSct)
	fmt.Printf("literalSct1: after instantiate: %v\n", literalSct1)

	// 使用结构体字面量实例化对象时，也可以省略 &
	var literalSct2 struct2 = struct2{10, 12, "literalSct"}
	fmt.Printf("literalSct2: after instantiate: %v\n", literalSct2)
}

// 5. 结构体在不同包中的调用方式
func t3() {
	structInSub := struct_sub_pkg.StructInSub{F1: 123}
	// 字段如果是私有的，那么在包外不能访问到
	//structInSub := struct_sub_pkg.StructInSub{F1: 123, f2: 33}
	//structInSub.f2 = 3

	fmt.Printf("structInSub: after instantiate: %v\n", structInSub)
}

// 5. 结构体是值类型
func t4() {
	var literalSct struct2 = struct2{10, 12, "literalSct"}
	var literalSct1 struct2 = literalSct

	fmt.Printf("literalSct: before field modification, literalSct: %v literalSct1: %v \n", literalSct, literalSct1)

	literalSct1.i2 = 200
	fmt.Printf("literalSct: after field modification, literalSct: %v literalSct1: %v \n", literalSct, literalSct1)
}

type Node struct {
	data float64
	// 这样写不行，要使用指针
	//su Node
	// 这里可以不使用指针，就是说如果引用自身类型，必须使用指针
	nodeExtra NodeExtra
	su        *Node
}
type NodeExtra struct {
	extra string
}

// 6. 递归结构体
func t5() {
	head := Node{data: 1.0}
	second := Node{data: 2.0}
	third := Node{data: 3.0}

	head.su = &second
	second.su = &third

	fmt.Printf("head: after instantiate: %v\n", head)
}

// 7. 类型转换
type number struct {
	f float32
}

// 这样是为number类型的结构体声明别名，该类型和number有相同的底层类型
type numberAlias number // alias type
func t6() {
	num := number{5.0}
	numAlias := numberAlias{6.0}

	fmt.Printf("num: after instantiate: %v\n", num)
	fmt.Printf("numAlias: after instantiate: %v\n", numAlias)

	// 别名类型和原始类型不能直接赋值
	//var convertedAliasNum1 number = numAlias
	// 需要转换后才能赋值
	var convertedAliasNum number = number(numAlias)
	fmt.Printf("aliasNum: after instantiate, num: %v, numAlias: %v, convertedAliasNum: %v\n", num, numAlias, convertedAliasNum)
}

func main() {
	//t1()
	//t2()
	//t3()
	t4()
	//t5()
	//t6()
}
