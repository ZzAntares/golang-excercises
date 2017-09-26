package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func md5file(filename string) ([]byte, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	hasher := md5.New()
	io.Copy(hasher, file)

	return hasher.Sum(nil), nil
}

func main() {
	// Iterate every directory and file beneath
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		hash, err := md5file(path)
		if err != nil {
			return err
		}
		fmt.Printf("%x\n", hash)

		return nil
	})
}
