package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

var rot map[rune]rune = map[rune]rune{
	'A': 'N',
	'B': 'O',
	'C': 'P',
	'D': 'Q',
	'E': 'R',
	'F': 'S',
	'G': 'T',
	'H': 'U',
	'I': 'V',
	'J': 'W',
	'K': 'X',
	'L': 'Y',
	'M': 'Z',
	'N': 'A',
	'O': 'B',
	'P': 'C',
	'Q': 'D',
	'R': 'E',
	'S': 'F',
	'T': 'G',
	'U': 'H',
	'V': 'I',
	'W': 'J',
	'X': 'K',
	'Y': 'L',
	'Z': 'M',
	'a': 'n',
	'b': 'o',
	'c': 'p',
	'd': 'q',
	'e': 'r',
	'f': 's',
	'g': 't',
	'h': 'u',
	'i': 'v',
	'j': 'w',
	'k': 'x',
	'l': 'y',
	'm': 'z',
	'n': 'a',
	'o': 'b',
	'p': 'c',
	'q': 'd',
	'r': 'e',
	's': 'f',
	't': 'g',
	'u': 'h',
	'v': 'i',
	'w': 'j',
	'x': 'k',
	'y': 'l',
	'z': 'm',
	' ': ' ',
}

func rot13(message string) string {
	var input []rune = []rune(message)
	var encoded []rune = make([]rune, len(input))

	for i, char := range input {
		encoded[i] = rot[char]
	}

	return string(encoded)
}

func handle(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		var input string = scanner.Text()
		io.WriteString(conn, rot13(input)+"\n\n")
	}
}

func main() {
	ln, err := net.Listen("tcp", "localhost:9000") // Create a server

	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()

		if err != nil {
			panic(err)
		}

		fmt.Println("Someone connected:", conn.RemoteAddr())

		go handle(conn)
	}

}
