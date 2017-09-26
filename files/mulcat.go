package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Usage: mulcat.go <f1> ...")
	}

	for _, filename := range os.Args[1:] {
		file, err := os.Open(filename)
		if err != nil {
			log.Fatalln(err)
		}
		defer file.Close()

		io.Copy(os.Stdout, file)
	}
}
