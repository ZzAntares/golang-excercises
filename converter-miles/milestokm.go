package main

import "fmt"

const MILE float64 = 1.60934

func GetDividerFor(str string) string {
	var divider string = "+"
	for i := 0; i < len(str)-2; i++ {
		divider += "-"
	}

	return divider + "+"
}

func GetSpan(a, b string) string {
	var la, lb int = len(a), len(b)
	var diff int

	if la < lb {
		diff = lb - la
	} else {
		diff = la - lb
	}

	var span string
	for i := 0; i <= diff; i++ {
		span += " "
	}

	return span
}

func main() {
	var miles float64

	fmt.Println("Miles: ")
	fmt.Scanf("%f", &miles)

	var query string = fmt.Sprintf("Miles: %.2f", miles)
	var result string = fmt.Sprintf("Kilometers: %.2f", miles*MILE)
	var span string = GetSpan(query, result)
	var divider string = GetDividerFor(fmt.Sprintf("| %s |", result))

	fmt.Println(divider)
	fmt.Printf("| %s%s|\n", query, span)
	fmt.Println(divider)
	fmt.Printf("| %s |\n", result)
	fmt.Println(divider)
}
