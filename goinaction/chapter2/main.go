package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	fmt.Print("hello go!!!")
}
