package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func md5file(filename string) [16]byte {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	hasher := md5.New()
	io.Copy(hasher, file)

	return md5.Sum(nil)
}

func walkStep(dir string, output chan<- string) {
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		output <- path
		return nil
	})

	close(output)
}

func md5Step(input <-chan string, output chan<- string) {
	for filename := range input {
		bs := md5file(filename)
		output <- fmt.Sprintf("%s %x\n", filename, bs)
	}

	close(output)
}

func printStep(input <-chan string) {
	for sumInfo := range input {
		fmt.Println(sumInfo)
	}
}

func main() {
	// Finds the MD5 hash of all files in current dir concurrently
	// Using the pipeline approach with unidirectional channels
	filenameChannel, sumChannel := make(chan string), make(chan string)

	go walkStep(".", filenameChannel)
	go md5Step(filenameChannel, sumChannel)
	printStep(sumChannel)
}
