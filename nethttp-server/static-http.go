package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	var root string = "."

	if len(os.Args) == 2 {
		root = os.Args[1]
	}

	fmt.Println("Listening on http://localhost:9000/ ...")
	http.ListenAndServe("localhost:9000", http.FileServer(http.Dir(root)))
}
