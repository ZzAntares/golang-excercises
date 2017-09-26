package main

// This program acts as a simple redis server with support of GET, SET and DEL
// However if the server stops, information will be gone!
// Maybe can add file storage support

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func db(input, output chan string, database map[string]string) {
	for command := range input {
		var tokens []string = strings.Fields(command)

		if len(tokens) == 0 {
			output <- "No commands were given!"
			return
		}

		tokens[0] = strings.ToUpper(tokens[0])

		switch tokens[0] {
		case "SET":
			if len(tokens) < 3 {
				output <- fmt.Sprintf(
					"SET requires 2 parameters %d given.", len(tokens))
				break
			}
			var key, value string = tokens[1], tokens[2]
			database[key] = value
			var response string = fmt.Sprintf("%s=%s", tokens[1], tokens[2])
			log.Printf("Key was SET: %s\n", response)
			output <- response
		case "GET":
			if len(tokens) < 2 {
				output <- fmt.Sprintf(
					"GET requires 1 parameter %d given.",
					len(tokens))
				break
			}

			if response, created := database[tokens[1]]; created {
				log.Printf("Record was retrieved: %s:%s\n", tokens[1], response)
				output <- response
				break
			}

			output <- fmt.Sprintf("Key %s does not exist!", tokens[1])
		case "DEL":
			if len(tokens) < 2 {
				output <- fmt.Sprintf(
					"DEL requires 1 parameter %d given.",
					len(tokens))
				break
			}

			if val, present := database[tokens[1]]; !present {
				output <- fmt.Sprintf("%s does not exist, nothing to do.", val)
				break
			}

			delete(database, tokens[1])
			log.Println("Record was deleted:", tokens[1])
			output <- fmt.Sprintf("ok!")
		default:
			output <- fmt.Sprintf("%s command not supported.", tokens[0])
		}
	}
}

func process(conn net.Conn, tunnel, responseTunnel chan string) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		tunnel <- scanner.Text()
		var response string = fmt.Sprintf("%s\n", <-responseTunnel)
		io.WriteString(conn, response)
	}
}

func main() {
	ln, err := net.Listen("tcp", "localhost:9000") // Create a server on 9000
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	var c chan string = make(chan string)
	var out chan string = make(chan string)
	var dbmap map[string]string = make(map[string]string)
	go db(c, out, dbmap)

	for {
		conn, err := ln.Accept() // Accept incoming connections
		if err != nil {
			panic(err)
		}

		fmt.Println("Someone connected:", conn.RemoteAddr())

		go process(conn, c, out)
	}
}
