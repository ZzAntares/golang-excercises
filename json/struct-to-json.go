package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name   string
	Age    int
	Salary float64
}

func main() {
	var messi Person = Person{Name: "Lionel Messi", Age: 33, Salary: 123234.45}
	contents, err := json.Marshal(messi)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(contents))
}
