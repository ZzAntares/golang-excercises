package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func md5file(filename string, c chan string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	hasher := md5.New()
	io.Copy(hasher, file)

	c <- fmt.Sprintf("%x\t%s\n", hasher.Sum(nil), filename)
}

func main() {
	var responseChannel chan string = make(chan string)
	var counter int

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		go md5file(path, responseChannel)
		counter++

		return nil
	})

	for i := 0; i < counter; i++ {
		var hashfile string = <-responseChannel
		fmt.Println(hashfile)
	}
}
