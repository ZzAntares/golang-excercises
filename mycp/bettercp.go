package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func cp(source, destination string) error {
	fsrc, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("Couldn't open source: %v", err)
	}
	defer fsrc.Close()

	fdst, err := os.Create(destination)
	if err != nil {
		return fmt.Errorf("Couldn't create destination: %v", err)
	}
	defer fdst.Close()

	_, err = io.Copy(fdst, fsrc) // Copy on the fly
	if err != nil {
		return fmt.Errorf("Couldn't copy source to destination: %v", err)
	}

	return err
}

func main() {
	if len(os.Args) < 3 {
		log.Fatalln("Usage: bettercp <SRC> <DST>")
	}

	var source, destination string = os.Args[1], os.Args[2]

	err := cp(source, destination)

	if err != nil {
		log.Fatalln(err)
	}
}
