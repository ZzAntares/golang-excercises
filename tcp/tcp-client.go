package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// connect to this socket (the server)
	connection, _ := net.Dial("tcp", "127.0.0.1:8081")

	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')

		// send to socket
		fmt.Fprintf(connection, text+"\n")

		// listen for reply
		reader = bufio.NewReader(connection)
		message, _ := reader.ReadString('\n')

		fmt.Print("Message from server: " + message)
	}
}
