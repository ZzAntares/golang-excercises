package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		fmt.Println(scanner.Text())

		// Read client text
		fmt.Print("> ")
		var input string
		fmt.Scanln(&input)

		if strings.TrimSpace(input) == "" {
			break
		}

		io.WriteString(conn, input)
		fmt.Println("Sent:", input)
	}
}
