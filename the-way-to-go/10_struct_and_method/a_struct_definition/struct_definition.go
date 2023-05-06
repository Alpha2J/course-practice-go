package a_struct_definition

// 1. 定义
type identifier struct {
	field1 int
	field2 int
	//    如果字段在代码中从来也不会被用到，那么可以命名它为 _
	_ int
}

// 2. 声明
var id identifier = null
// 3. 赋值
id.field1 = 1
id.field2 = 2