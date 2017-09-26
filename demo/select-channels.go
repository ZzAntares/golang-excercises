package main

import (
	"fmt"
	"time"
	"math/rand"
)

func writer(msg string, c chan string) {
	for {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
		c <- msg
	}
}

func main() {
	var c chan string = make(chan string)

	go writer("Writer #1", c)
	go writer("Writer #2", c)

	for {
		// Gets the first channel with messages and discards all others
		select {
		case msg := <- c:
			fmt.Println("Got a message:", msg)
		case msg := <- c:
			fmt.Println("Got the message: ", msg)
		}
	}

	var input string
	fmt.Scanln(&input)
}
