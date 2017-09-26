package main

import (
	"net/http"
)

// Simple file server

func main() {
	dir := http.Dir(".")
	http.Handle("/", http.FileServer(dir))

	http.ListenAndServe("localhost:9000", nil)
}
