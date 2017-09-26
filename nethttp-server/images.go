package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func CatHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Hello there")

	file, err := os.Open("grumpy-cat.jpg")
	if err != nil {
		http.Error(res, "Sorry no cat for you", 404)
		return
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		http.Error(res, "No image info", 404)
		return
	}

	// io.Copy(res, file) // Send image bytes to the response
	// Using ServeContent sets propper Etag to allow browsers cache content
	http.ServeContent(res, req, file.Name(), info.ModTime(), file)
}

func main() {
	http.HandleFunc("/grumpy-cat.jpg", CatHandler)
	http.HandleFunc("/dog-meme.png", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, "dog-meme.png") // Simpler way to serve files
	})

	http.HandleFunc("/cat/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, `<!DOCTYPE html>
<html>
<head></head>
<body>
<img src="/grumpy-cat.jpg" />
</body>
</html>
`)
	})

	http.HandleFunc("/dog/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, `<!DOCTYPE html>
<html>
<head></head>
<body>
<img src="/dog-meme.png" />
</body>
</html>
`)
	})

	http.ListenAndServe("localhost:9000", nil)
}
