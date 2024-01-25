package main

import (
	//不能在不同的目录内声明同一个package，这样无法import
	//"cource-practice-go/the-way-to-go/9_package/e_custom_package/duplicatepack1"
	"cource-practice-go/the-way-to-go/9_package/e_custom_package/pack1"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

import . "cource-practice-go/the-way-to-go/9_package/e_custom_package/pack2"

// 4. 一般情况下包引入后就必须被使用, 如果不想使用这个包, 只执行它的 init 函数并初始化其中的全局变量, 可以使用这种方式
import _ "cource-practice-go/the-way-to-go/9_package/e_custom_package/pack3"
import _ "cource-practice-go/the-way-to-go/9_package/e_custom_package/pack4"

// 1. 包的引入方式一般为: import "包的路径或 URL 地址"
//   - 如: import "github.com/org1/pack1”
func t1() {
	fmt.Println(pack1.Pack1Int)
	fmt.Println(pack1.ReturnStr())
}

// 2. 在同一个包内的变量可以直接引用
func t2() {
	fmt.Println(Pack2Int)
}

// 3. 使用 import . "./pack1" 的方式引入包后, 可以不通过包名来使用其中的项目
func t3() {
	fmt.Println(Pack2Float)
}

// 5. 导入外部安装包:
//   - go get codesite.ext/author/goExample/goex
//   - import "github.com/gin-gonic/gin"
//
// 可以在这里搜索常用的外部包 https://pkg.go.dev/
func t4() {
	// 创建 Gin 引擎
	r := gin.Default()

	// 定义 HTTP 路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, world!",
		})
	})

	// 启动 Web 服务器
	r.Run(":19923")
}

// 6. 包名和目录名不同的时候，import的时候还是使用目录名，但是使用包内关键字的时候是需要用包名来调用
// ps: 在Go语言中，包的导入路径通常与包的名称相同，这是Go语言的一种约定.
func t5() {
	//代码能正常跑，但是ide报红，不清楚原因，先注释了
	//fmt.Println(pack5_diff.Pack5Int)
}

// 7. 程序的执行:
//   - 执行导入包的init()函数
//   - 执行main 包的init()函数
//   - 执行main()代码
func init() {
	fmt.Println("i am initialing main")
}

func main() {
	//t1()
	//t2()
	t3()
	//t4()
	t5()
}
