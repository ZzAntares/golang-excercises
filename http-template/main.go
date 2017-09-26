package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		tpl, err := template.ParseFiles("hello.html.tmpl")
		if err != nil {
			panic(err)
		}

		tpl.Execute(res, map[string]string{
			"Message": req.RequestURI,
		})
	})

	http.ListenAndServe("localhost:9000", nil)
}
