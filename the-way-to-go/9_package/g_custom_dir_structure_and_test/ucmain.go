package main

import (
	"cource-practice-go/the-way-to-go/9_package/g_custom_dir_structure_and_test/src/uc"
	"fmt"
)

func main() {
	str1 := "USING package uc!"
	fmt.Println(uc.UpperCase(str1))
}

//todo: https://learnku.com/docs/the-way-to-go/98-custom-package-directory-structure-go-install-and-go-test/3633
// go install uc 报错: package uc is not in GOROOT (/usr/local/go/src/uc)
// 还不知道怎么解决, 有时间再看看
