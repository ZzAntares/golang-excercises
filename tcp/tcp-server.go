package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081")

	// accept connection on port
	connection, _ := ln.Accept()

	// run loop forever (or until ctrl-c)
	for {
		// listen for message to process ending in newline (\n)
		reader := bufio.NewReader(connection)
		message, _ := reader.ReadString('\n')

		// output message received
		fmt.Print("Message Received: ", string(message))

		// sample process for string received
		new_message := strings.ToUpper(message)

		// send the new string back to client
		connection.Write([]byte(new_message + "\n"))
	}
}
