package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", "localhost:9000") // Create a server in 9000
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept() // Accept incoming connection
		if err != nil {
			panic(err)
		}

		fmt.Println("Someone connected:", conn.RemoteAddr())
		io.WriteString(conn, fmt.Sprintf("Welcome %s!", conn.RemoteAddr()))

		// Support multiple connections
		go func() {
			// io.Copy(conn, conn)
			scanner := bufio.NewScanner(conn)
			for scanner.Scan() {
				io.WriteString(conn, scanner.Text())
			}

			conn.Close()
		}()
	}
}
