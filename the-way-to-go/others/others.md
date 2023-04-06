1. goroot gopath gomodule 分别是什么
  - OROOT：GOROOT 是 Go 语言的安装根目录，它包含了 Go 语言的标准库和一些工具。GOROOT 环境变量指向安装根目录，用于告诉 Go 编译器和工具在哪里找到标准库和工具。 
  - GOPATH：GOPATH 是 Go 语言的工作目录，用于存放你的 Go 项目和第三方依赖库。它是一个包含多个目录的列表，通常包括一个 bin 目录、一个 pkg 目录和一个 src 目录。你的 Go 项目应该位于 GOPATH/src 目录下的子目录中，编译后的二进制文件将被放置在 GOPATH/bin 目录中，而依赖库则会被放置在 GOPATH/pkg 目录中。 
  - Go Modules：Go Modules 是 Go 语言的依赖管理工具，它可以帮助你管理你的项目依赖库。与传统的 GOPATH 目录结构不同，Go Modules 可以将依赖库下载到项目根目录下的 vendor 目录中，而不是 GOPATH/pkg 目录。它使用 go.mod 文件来管理依赖库的版本和依赖关系。
2. go get 下载的包被安装到哪里了
  - 通过 go get 命令下载的包会被安装在 GOPATH 环境变量指定的路径下。通常情况下，GOPATH 的默认值为 $HOME/go（在 Windows 中为 %USERPROFILE%\go）。在这个路径下，每个包都会被安装到一个特定的目录中，例如 github.com/gin-gonic/gin 这个包会被安装到 $GOPATH/src/github.com/gin-gonic/gin 目录下。 
  - 需要注意的是，如果你使用了 Go 1.16 及以上的版本，那么 GOPATH 的默认值会被设置为一个空目录，也就是说，通过 go get 命令下载的包会被安装在 $HOME/go/pkg/mod 目录下，而不再是 $GOPATH/src 目录。

