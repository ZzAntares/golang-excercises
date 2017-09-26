package main

import (
	"fmt"
	"io/ioutil"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9000")
	defer conn.Close()

	if err != nil {
		panic(err)
	}

	bs, err := ioutil.ReadAll(conn)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bs))
}
