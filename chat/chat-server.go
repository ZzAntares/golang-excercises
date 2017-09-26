package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type Message struct {
	Username string
	Text     string
}

type User struct {
	Name string
	Feed chan Message
}

type ChatServer struct {
	RecMessage chan Message
	Join       chan User
	Leave      chan User
	Users      map[string]User
}

func (server *ChatServer) Start() {
	for {
		select {
		case user := <-server.Join:
			server.Users[user.Name] = user // Join the user

			go func() {
				var msg Message = Message{
					Text:     user.Name + " has joined.",
					Username: "SYSTEM",
				}

				server.RecMessage <- msg
			}()
		case user := <-server.Leave:
			delete(server.Users, user.Name) // Remove user

			go func() {
				var msg Message = Message{
					Text:     user.Name + " has left.",
					Username: "SYSTEM",
				}

				server.RecMessage <- msg
			}()
		case message := <-server.RecMessage:
			// Send messages to every user
			for username, user := range server.Users {
				// Don't send the message to the user whose sent the message
				if message.Username == username {
					continue
				}

				user.Feed <- message
			}
		}
	}
}

func handle(server *ChatServer, conn net.Conn) {
	defer conn.Close()

	// Ask for user name
	fmt.Fprint(conn, "Enter your username: ")
	scanner := bufio.NewScanner(conn)
	scanner.Scan()
	var username string = scanner.Text()

	// Create the User struct and the Feed channel
	var inputChannel chan Message = make(chan Message)
	var user User = User{
		Feed: inputChannel,
		Name: username,
	}

	// Join and broadcast
	server.Join <- user
	defer func() {
		server.Leave <- user
	}()

	// Read input and send to server async
	go func() {
		for scanner.Scan() {
			server.RecMessage <- Message{
				Username: username,
				Text:     scanner.Text(),
			}
		}
	}()

	// Get messages in a for loop could also be async?
	for message := range user.Feed {
		_, err := fmt.Fprintf(conn, "%s: %s\n", message.Username, message.Text)
		if err != nil {
			break // User left
		}
	}
}

func main() {
	ln, err := net.Listen("tcp", "localhost:9000") // Create a server on 9000
	if err != nil {
		log.Fatalln(err)
	}
	defer ln.Close()

	// Start a chat server
	var server ChatServer = ChatServer{
		RecMessage: make(chan Message),
		Join:       make(chan User),
		Leave:      make(chan User),
		Users:      make(map[string]User),
	}
	go server.Start()

	for {
		conn, err := ln.Accept() // Accept incoming connections

		if err != nil {
			log.Fatalln(err)
		}

		log.Println("Someone connected:", conn.RemoteAddr())

		go handle(&server, conn)
	}
}
