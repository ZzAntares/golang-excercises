package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	var filename string = "demo.go"
	f, err := os.Open(filename)

	if err != nil {
		log.Fatalln("Can't open file!:", err.Error())
	}

	defer f.Close()

	content, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln("Can't open file!:", err.Error())
	}

	var lines []string = strings.Split(string(content), "\n")
	var tokens []string
	for _, v := range lines {
		var stokens []string = strings.Fields(v)
		for _, t := range stokens {
			tokens = append(tokens, strings.TrimSpace(t))
		}
	}

	content = []byte(strings.Join(tokens, ";"))
	f, err = os.Create(filename)
	defer f.Close()

	if err != nil {
		log.Fatalln("Can't write!:", err)
	}

	st, err := f.Write(content)

	fmt.Println("Status:", st)
	fmt.Println("Error?:", err)
}
