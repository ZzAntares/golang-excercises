package main

import (
	"fmt"
	"time"
)

func routineA(send, rec chan string) {
	for i := 0; i <= 10; i++ {
		send <- "ping"
		var response string = <- rec
		fmt.Println(response)
	}
}

func routineB(rec, send chan string) {
	for {
		var message string = <- rec
		fmt.Println(message)
		time.Sleep(time.Second * 1)
		send <- "pong"
	}
}

func main() {
	var rec chan string = make(chan string)
	var send chan string = make(chan string)

	go routineA(send, rec)
	go routineB(send, rec)

	var input string
	fmt.Scanln(&input)
}
