package main

import "net/http"
import "io"

type DogHandler struct{}

func (h DogHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var body string = `
<html>
<head></head>
<body>
<img src="http://i3.kym-cdn.com/photos/images/facebook/000/581/722/7bc.jpg" />
</body>
</html>
`
	io.WriteString(res, body)
}

type CatHandler struct{}

func (h CatHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var body string = `
<html>
<head></head>
<body>
<img src="https://media.mnn.com/assets/images/2013/05/grumpyCatComplain.jpg.838x0_q80.jpg" />
</body>
</html>
`
	io.WriteString(res, body)
}

func main() {
	var ch, dh http.Handler = CatHandler{}, DogHandler{}

	mux := http.NewServeMux()

	mux.Handle("/dog/", dh)
	mux.Handle("/cat/", ch)

	http.ListenAndServe("localhost:9000", mux)
}
