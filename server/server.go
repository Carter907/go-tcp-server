package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {

	listen, err := net.Listen("tcp", "localhost:8081")
	if err != nil {

		fmt.Println("Error: ", err.Error())
	}
	defer listen.Close()
	for {
		connection, err := listen.Accept()
		if err != nil {
			fmt.Println("Error: ", err.Error())
			return
		}
		go handleRequest(connection)
	}
}

func handleRequest(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}

	}
	c.Close()
}
