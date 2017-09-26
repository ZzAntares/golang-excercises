package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl := template.New("addition.gohtml")

	tpl.Funcs(template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
	})

	tpl, err := tpl.ParseFiles("addition.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
