package main

import (
	"net"
	"fmt"
	"bufio"
	"strings"
)

func main() {
	fmt.Println("launching server")
	// listen on all interfaces
	ln, _ := net.Listen("tcp",":8081")
	// accept connection on port
	conn, _ := ln.Accept()


	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message recieved:", string(message))

		newmessage := strings.ToUpper(message)
		conn.Write([]byte(newmessage+ "\n"))
	}
}