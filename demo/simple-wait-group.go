package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printer(id int) {
	fmt.Println("Hello", id)
	wg.Done()  // Mark just one work as done (decrement)
}

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)  // Add one work to the wait group (increment)
		go printer(i)
	}

	wg.Wait()  // Wait until all works are done (increment - decrement == 0)
}
