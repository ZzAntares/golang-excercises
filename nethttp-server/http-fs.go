package main

import (
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		http.HandleFunc("/"+path, func(res http.ResponseWriter, req *http.Request) {
			http.ServeFile(res, req, path)
		})

		return nil
	})

	http.ListenAndServe("localhost:9000", nil)
}
