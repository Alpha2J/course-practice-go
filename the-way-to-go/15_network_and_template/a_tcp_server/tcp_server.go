package main

import "time"

func main() {
	go startServer()
	time.Sleep(1e9)
	startClient()
}
