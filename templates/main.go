package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("template.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	var fruits []string = []string{
		"limon",
		"orange",
		"watermelon",
	}

	err = tpl.Execute(os.Stdout, fruits)
}
