package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/record", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("hokaj")
		var firstname string = req.FormValue("firstname")
		var lastname string = req.FormValue("lastname")

		io.WriteString(res, firstname+" "+lastname)
	})

	http.HandleFunc("/main", func(res http.ResponseWriter, req *http.Request) {
		form := `
<!DOCTYPE html>
<html>
<head></head>
<body>
  <form method="POST" action="record">
    <input name="firstname" type="text" />
    <input name="lastname" type="text" />
    <input type="submit" />
  </form>
</body>
</html>
`
		io.WriteString(res, form)
	})

	http.ListenAndServe("localhost:9000", nil)
}
