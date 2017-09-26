package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"io"
	"net/http"
)

func getTamperingCode(data string) string {
	h := hmac.New(sha256.New, []byte("not-so-secret"))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		cookie, err := req.Cookie("my-cookie")
		if err != nil {
			// The cookie is not present so set it
			uuid, _ := uuid.NewV4() // Create a "session" id

			code := getTamperingCode(cookie.Value)
			http.SetCookie(w, &http.Cookie{
				Name:  "my-cookie-session",
				Value: code + "|" + uuid.String(),
			})

		}

		io.WriteString(w, "Value of cookie: "+cookie.Value)
		io.WriteString(w, "\nCookie was set!")
	})

	http.ListenAndServe("localhost:9000", nil)
}
