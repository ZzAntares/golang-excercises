package main

import "fmt"

func filter(seq []int, callback func(int) bool) []int {
	var result []int

	for _, el := range seq {
		if callback(el) {
			result = append(result, el)
		}
	}

	return result
}

func main() {
	var numbers []int = []int{1, 2, 3, 4, 5, 6, 7, 8}

	var oddNumbers []int = filter(numbers, func(n int) bool {
		return n%2 != 0
	})

	var bigNumbers []int = filter(numbers, func(n int) bool {
		return n > 5
	})

	fmt.Println("Odd numbers:", oddNumbers)
	fmt.Println("Big numbers:", bigNumbers)
}
