package main

import (
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

// 6. 程序的执行:
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
	t4()
}
