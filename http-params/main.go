package main

import (
	"html/template"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/receiver",
		func(res http.ResponseWriter, req *http.Request) {
			file, fheader, err := req.FormFile("ufile")
			if err != nil {
				http.Error(res, "Can't handle file "+err.Error(), 500)
				return
			}
			defer file.Close()

			// Not recommended in prod
			outputFile, err := os.Create("/tmp/" + fheader.Filename)
			if err != nil {
				http.Error(res, "Error uploading file to destination.", 500)
				return
			}
			defer outputFile.Close()

			// Write uploaded byte stream to output file

			_, err = io.Copy(outputFile, file)
			if err != nil {
				http.Error(res, "Error writing to destination.", 500)
				return
			}

			io.WriteString(res, "Upload OK!")
		})

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		tpl, err := template.ParseFiles("form.html.tpl")
		if err != nil {
			http.Error(res, "No template found", 500)
			return
		}

		// Render form page

		tpl.Execute(res, nil)
	})

	http.ListenAndServe("localhost:9000", nil)
}
