package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: by-lines <SRC FILE>")
	}

	src, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	defer src.Close()

	scanner := bufio.NewScanner(src)

	for scanner.Scan() {
		var line string = scanner.Text()
		fmt.Println(">>>", line)
	}
}
