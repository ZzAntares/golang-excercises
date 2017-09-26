package main

import (
	"fmt"
	"time"
	"math/rand"
)

func f(n int) {
	for i := 0; i <= 10; i++ {
		timeDuration := time.Duration(rand.Intn(250))
		time.Sleep(timeDuration)
		fmt.Println(n, "=>", i)
	}
}

func main() {
	for i := 0; i <= 10; i++ {
		go f(i)
	}

	var input string
	fmt.Scanln(&input)
}
