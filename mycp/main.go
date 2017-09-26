package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func cp(fromFile, toFile string) error {
	f, err := os.Open(fromFile)
	defer f.Close()

	if err != nil {
		return fmt.Errorf("Couldn't open source file: %v", err)
	}

	content, err := ioutil.ReadAll(f)
	if err != nil {
		return fmt.Errorf("Couldn't read source file: %v", err)
	}

	f, err = os.Create(toFile)
	if err != nil {
		return fmt.Errorf("Couldn't open destination file: %v", err)
	}
	defer f.Close()

	_, err = f.Write(content)
	if err != nil {
		return fmt.Errorf("Couldn't write to destination file: %v", err)
	}

	return nil
}

func main() {
	var fromFile, toFile string = os.Args[1], os.Args[2]

	err := cp(fromFile, toFile)

	if err != nil {
		log.Fatalln(err)
	}
}
