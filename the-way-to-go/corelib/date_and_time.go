package corelib

import (
	"fmt"
	"time"
)

func now() {
	fmt.Println("now is: ", time.Now())
}

func init() {
	now()
}
