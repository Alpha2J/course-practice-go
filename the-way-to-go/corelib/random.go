package corelib

import (
	"fmt"
	"math/rand"
)

// 随机数生成
func randomInt() int {
	return rand.Int()
}

func init() {
	fmt.Println("random int: ", randomInt())
}
