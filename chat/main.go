package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

var connections []net.Conn

func server(input chan string) {
	for message := range input {
		for _, c := range connections {
			io.WriteString(c, message)
		}
	}
}

func handle(conn net.Conn, output chan string) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		var message string = scanner.Text()
		output <- fmt.Sprintf("%s\n", message)
	}
}

func main() {
	ln, err := net.Listen("tcp", "localhost:9000") // Create a server on 9000
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	var c chan string = make(chan string)
	go server(c)

	for {
		conn, err := ln.Accept() // Accept incoming connections
		if err != nil {
			panic(err)
		}
		connections = append(connections, conn)

		fmt.Println("Someone connected:", conn.RemoteAddr())

		go handle(conn, c)
	}
}
