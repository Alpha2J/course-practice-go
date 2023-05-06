package main

import (
	"container/list"
	"fmt"
)

// 知识点:
// 标准库：
// 1. unsafe: 包含了一些打破 Go 语言 “类型安全” 的命令，一般的程序中不会被使用，可用在 C/C++ 程序的调用中。
// 2. syscall - os - os/exec:
//   - os: 提供给我们一个平台无关性的操作系统功能接口，采用类 Unix 设计，隐藏了不同操作系统间差异，让不同的文件系统和操作系统对象表现一致。
//   - os/exec: 提供我们运行外部操作系统命令和程序的方式。
//   - syscall: 底层的外部包，提供了操作系统底层调用的基本接口。
//
// 3. archive/tar 和 /zip-compress: 压缩 (解压缩) 文件功能。
// 4. fmt - io - bufio-path/filepath-flag:
//   - fmt: 提供了格式化输入输出功能。
//   - io: 提供了基本输入输出功能，大多数是围绕系统功能的封装。
//   - bufio: 缓冲输入输出功能的封装。
//   - path/filepath: 用来操作在当前系统中的目标文件名路径。
//   - flag: 对命令行参数的操作。
//
// 5. strings - strconv - unicode - regexp - bytes:
//   - strings: 提供对字符串的操作。
//   - strconv: 提供将字符串转换为基础类型的功能。
//   - unicode: 为 unicode 型的字符串提供特殊的功能。
//   - regexp: 正则表达式功能。
//   - bytes: 提供对字符型分片的操作。
//   - index/suffixarray: 子字符串快速查询。
//
// 6. math - math/cmath - math/big - math/rand - sort:
//   - math: 基本的数学函数。
//   - math/cmath: 对复数的操作。
//   - math/rand: 伪随机数生成。
//   - sort: 为数组排序和自定义集合。
//   - math/big: 大数的实现和计算。
//
// 7. container-/list-ring-heap: 实现对集合的操作。
//   - list: 双链表。
//   - ring: 环形链表。
// 8. time - log:
//   - time: 日期和时间的基本操作。
//   - log: 记录程序运行时产生的日志，我们将在后面的章节使用它。
// 9. encoding/json - encoding/xml - text/template:
//   - encoding/json: 读取并解码和写入并编码 JSON 数据。
//   - encoding/xml: 简单的 XML1.0 解析器，有关 JSON 和 XML 的实例请查阅第 12.9/10 章节。
//   - text/template: 生成像 HTML 一样的数据与文本混合的数据驱动模板（参见第 15.7 节）。
// 10. net - net/http - html:（参见第 15 章）
//   - net: 网络数据的基本操作。
//   - http: 提供了一个可扩展的 HTTP 服务器和基础客户端，解析 HTTP 请求和回复。
//   - html: HTML5 解析器。
// 11. runtime: Go 程序运行时的交互操作，例如垃圾回收和协程创建。
// 12. reflect: 实现通过程序运行时反射，让程序操作任意类型的变量。
//

// 练习9.1: 使用 container/list 包实现一个双向链表，将 101、102 和 103 放入其中并打印出来。
func exercise9_1() {
	l := list.New()

	l.PushBack(101)
	l.PushBack(102)
	l.PushBack(103)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func main() {
	exercise9_1()
}
