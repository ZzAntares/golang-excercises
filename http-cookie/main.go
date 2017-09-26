package main

import (
	"github.com/gorilla/sessions"
	"io"
	"net/http"
)

var cookieStore = sessions.NewCookieStore([]byte("auth-secret-passwd"))

func IsLoggedIn(req *http.Request) bool {
	session, err := cookieStore.Get(req, "my-session")
	if err != nil {
		// Session has been tampered
		return false
	}

	_, ok := session.Values["email"]
	return ok
}

func main() {
	http.HandleFunc("/logout", func(w http.ResponseWriter, req *http.Request) {
		session, err := cookieStore.Get(req, "my-session")
		if err != nil {
			// Session has been tampered
			http.Error(w, "O. "+err.Error(), http.StatusInternalServerError)
			return
		}

		delete(session.Values, "email")
		session.Save(req, w)

		http.Redirect(w, req, "/", 302)
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, req *http.Request) {
		if IsLoggedIn(req) {
			http.Error(w,
				"2. You're already logged in.",
				http.StatusInternalServerError)
			return
		}

		if req.Method == "GET" {
			// draw form and return
			form := `<!DOCTYPE html>
<html>
<head></head>
<body>
  <form method="POST">
    <input name="email" type="text" placeholder="User" />
    <input name="passwd" type="password" placeholder="Password" />
    <input type="submit" />
  </form>
</body>
</html>
`
			io.WriteString(w, form)
			return
		}

		session, err := cookieStore.Get(req, "my-session")
		if err != nil {
			// Session has been tampered
			http.Error(w, "1. "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Email is not present in session, so we set it.
		session.Values["email"] = req.FormValue("email")
		session.Save(req, w)

		// Redirect to home
		http.Redirect(w, req, "/", 302)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if IsLoggedIn(req) {
			io.WriteString(w, "Proceed to /logout")
			return
		}

		// is not logged in
		io.WriteString(w, "Proceed to /login")
	})

	http.ListenAndServe("localhost:9000", nil)
}
