package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func MyHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Gorilla!\n"))
}

func gorilla() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function
	r.HandleFunc("/", MyHandler)

	// Bind to a port and pass our router in
	fmt.Println("Listening on http://localhost:8000/ ...")
	log.Fatal(http.ListenAndServe(":8000", r))
}
