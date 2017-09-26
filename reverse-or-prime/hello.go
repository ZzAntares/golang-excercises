package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Reverse(message string) string {
	runes := []rune(message) // Convierte a un arreglo de runas la cadena

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func PrimeOrNot(number int) bool {
	var upto int = number / 2

	for i := 2; i <= upto; i++ {
		if number%i == 0 {
			return true
		}
	}

	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Reverse a string")
	fmt.Println("----------------")

	fmt.Println("Write something: ")
	response, _ := reader.ReadString('\n')

	// Clean string
	response = strings.Trim(response, "\n")
	response = strings.Trim(response, " ")

	fmt.Printf("You wrote is: %s\n", response)

	fmt.Println("Reversing string ...")
	fmt.Println(Reverse(response))
}
