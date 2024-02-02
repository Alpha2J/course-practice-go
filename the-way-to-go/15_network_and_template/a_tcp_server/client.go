package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func startClient() {
	// 1. 创建和服务器之间的连接，返回的conn可以用来发送和接收数据
	// IPv4、IPv6，TCP、UDP都可以使用这个接口
	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("First, what is your name?")
	clientName, _ := inputReader.ReadString('\n')
	trimmedClient := strings.Trim(clientName, "\r\n")

	for {
		fmt.Println("What to send to the server? Type Q to quit.")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		if trimmedInput == "Q" {
			return
		}

		_, err = conn.Write([]byte(trimmedClient + " says: " + trimmedInput))
	}
}
