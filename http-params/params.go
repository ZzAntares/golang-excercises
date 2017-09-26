package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/file", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/html")

		file, header, err := req.FormFile("llave")
		fmt.Println("File: ", file)
		fmt.Println("Header: ", header)
		fmt.Println("Err: ", err)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		bs, _ := ioutil.ReadAll(file)
		fmt.Println("Contents:", string(bs))

		form := `<form method="POST" enctype="multipart/form-data">
<input type="file" name="llave" />
<input type="submit" />
</form>`
		fmt.Fprintf(res, form)
	})

	http.HandleFunc("/get", func(res http.ResponseWriter, req *http.Request) {
		parameters := req.URL.Query()
		fmt.Fprintf(res, "Data: %v\n", parameters)
	})

	http.HandleFunc("/post", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/html")

		value := req.FormValue("llave") // Also gives query string params on GET
		form := `<form method="POST">
<input type="text" name="llave" />
<input type="submit" />
</form>`
		fmt.Fprintf(res, form)
		fmt.Println("You send on post: llave=" + value)
	})

	http.ListenAndServe("localhost:9000", nil)
}
