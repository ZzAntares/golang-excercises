package main

import (
	"fmt"
	"time"
)

func rutineA(channel chan string) {
	for i := 0; i <= 10; i++ {
		var message string = fmt.Sprintf("Message #%d\n", i)
		channel <- message
		fmt.Println("Sent: ", message)
	}
}

func rutineB(channel chan string) {
	for {
		var message string = <- channel
		fmt.Println("Received: ", message)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var channel chan string = make(chan string)

	go rutineA(channel)
	go rutineB(channel)

	var input string
	fmt.Scanln(&input)
}
