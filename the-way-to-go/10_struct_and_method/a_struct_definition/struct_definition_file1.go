package main

// 4. 一个结构体在一个包中只能定义一次，所以这里无法编译通过
//type struct1 struct {
//	i1   int
//	f1   float32
//	str1 string
//	//    如果字段在代码中从来也不会被用到，那么可以命名它为 _
//	_ int
//}
