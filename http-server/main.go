package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// Server replies with the requested url... literally

func handle(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	var i int
	for scanner.Scan() {
		line := scanner.Text()

		if i == 0 {
			// GET /url HTTP/1.1
			var tokens []string = strings.Fields(line)
			var method, url, version string = tokens[0], tokens[1], tokens[2]
			fmt.Println("Method:", method)
			fmt.Println("Url:", url)
			fmt.Println("Version:", version)

			var response string
			response += "HTTP/1.1 200 OK\r\n"
			response += fmt.Sprintf("Content-Length: %d\r\n", len(url))
			response += "\r\n"
			response += url

			fmt.Fprintf(conn, response)
		} else {
			// Headers section
			fmt.Println(line)
			if line == "" {
				// EOF headers
				break
			}
		}

		i++
	}
}

func main() {
	ln, err := net.Listen("tcp", "localhost:9000") // Create a server on 9000
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept() // Accept incoming connections
		if err != nil {
			panic(err)
		}

		go handle(conn)
	}
}
