package main

import (
	"fmt"
	"github.com/gorilla/sessions"
	"io"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("very-secret-stuff"))

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// Get a session not a specific cookie
		session, err := store.Get(req, "my-session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// retrieve counter
		rawCounter := session.Values["counter"]
		var counter int
		counter, _ = rawCounter.(int)
		counter += 1

		// Set the new counter
		session.Values["counter"] = counter

		session.Save(req, w)
		io.WriteString(w, fmt.Sprintf("Número de visitas: %d\n", counter))
	})

	http.ListenAndServe("localhost:9000", nil)
}
