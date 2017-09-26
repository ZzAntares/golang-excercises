package main

import (
	"fmt"
	"time"
)


func pinger(id int, tunnel chan string) {
	for {
		tunnel <- fmt.Sprint(id, "ping")
	}
}


func printer(tunnel chan string, delay time.Duration) {
	time.Sleep(delay)
	for {
		var msg string = <- tunnel
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}


func main() {
	var c chan string = make(chan string)

	go pinger(0, c)
	go pinger(1, c)
	go pinger(2, c)
	go printer(c, 0)
	go printer(c, time.Millisecond * 500)

	var input string
	fmt.Scanln(&input)
}
