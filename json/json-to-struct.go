package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: json2struct <filename>")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	contents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Parsing:", string(contents))

	var obj map[string]interface{}
	json.Unmarshal(contents, &obj)

	fmt.Println(obj)
}
